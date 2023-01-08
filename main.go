package main

import (
	"flag"
	"fmt"
	"github.com/sunist-c/genius-invokation-simulator-backend/exec/backend"
	"github.com/sunist-c/genius-invokation-simulator-backend/exec/cli"
	"github.com/sunist-c/genius-invokation-simulator-backend/persistence"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	_ "github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/service"
)

const (
	aiMode      = "ai"
	cliMode     = "cli"
	backendMode = "backend"
)

var (
	args = &argument{
		mode:    new(string),
		port:    new(uint),
		conf:    new(string),
		save:    new(bool),
		flush:   new(uint),
		storage: new(string),
	}
	sig = make(chan os.Signal, 4)
)

type argument struct {
	mode    *string
	conf    *string
	port    *uint
	save    *bool
	flush   *uint
	storage *string
}

func initArgs() {
	flag.StringVar(args.conf, "conf", "", "setup the backend configuration file, highest priority")
	flag.StringVar(args.mode, "mode", "backend", "setup the startup mode, available [backend, cli, ai]")
	flag.UintVar(args.port, "port", 8086, "setup the http protocol port")
	flag.BoolVar(args.save, "save", true, "setup if to enable the persistence module")
	flag.StringVar(args.storage, "storage", path.Join(os.Args[0], "../data/persistence"), "setup the persistence storage filepath")
	flag.UintVar(args.flush, "flush", 3600, "setup the flush frequency(second) of persistence module")
}

func callQuit() {
	fmt.Printf("[main.log] main.callQuit(): daemon running\n")
	s := <-sig
	fmt.Printf("[main.log] main.callQuit(): quit signal received %v\n", s.String())
	fmt.Printf("[main.log] main.callQuit(): quiting persistence\n")
	persistence.Quit()
	fmt.Printf("[main.log] main.callQuit(): quited persistence\n")

	switch *args.mode {
	case backendMode:
		backend.Quit()
	case cliMode:
		cli.Quit()
	case aiMode:
		panic("not implemented yet")
	}

	fmt.Printf("[main.log] main.callQuit(): wait 10 seconds for quit task\n")
	time.Sleep(time.Second * 10)
	os.Exit(114)
}

func init() {
	errChan := make(chan error, 10)
	fmt.Printf("[main.log] main.init(): initializing main package\n")
	signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)
	fmt.Printf("[main.log] main.init(): setuped signal channel\n")
	initArgs()
	flag.Parse()
	fmt.Printf("[main.log] main.init(): parsed command line arguments\n")
	if err := persistence.SetStoragePath(*args.storage); err != nil {
		errChan <- err
		panic(err)
	}
	fmt.Printf("[main.log] main.init(): initialized storage path\n")
	persistence.Load(errChan)
	fmt.Printf("[main.log] main.init(): initializing persistent storage\n")
	go func() {
		fmt.Printf("[main.log] main.errorHandler(): error handler running\n")
		for err := range errChan {
			fmt.Println(err)
		}
	}()
	persistence.Serve(time.Second*time.Duration(*args.flush), errChan)
	fmt.Printf("[main.log] main.init(): initialize completed\n")
}

func main() {
	go callQuit()
	switch *args.mode {
	case backendMode:
		backend.Run(*args.port)
	case cliMode:
		cli.Run()
	case aiMode:
		panic("not implemented yet")
	default:
		os.Exit(0)
	}
}
