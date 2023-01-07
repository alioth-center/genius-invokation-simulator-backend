package main

import (
	"flag"
	"fmt"
	"github.com/sunist-c/genius-invokation-simulator-backend/exec/backend"
	"github.com/sunist-c/genius-invokation-simulator-backend/exec/cli"
	"github.com/sunist-c/genius-invokation-simulator-backend/persistence"
	"os"
	"os/signal"
	"syscall"
)

const (
	aiMode      = "ai"
	cliMode     = "cli"
	backendMode = "backend"
)

var (
	args = &argument{
		mode: new(string),
		port: new(uint),
		conf: new(string),
		save: new(bool),
	}
	sig = make(chan os.Signal, 4)
)

type argument struct {
	mode *string
	conf *string
	port *uint
	save *bool
}

func initArgs() {
	flag.StringVar(args.conf, "conf", "", "setup the backend configuration file, highest priority")
	flag.StringVar(args.mode, "mode", "backend", "setup the startup mode, available [backend, cli, ai]")
	flag.UintVar(args.port, "port", 8086, "setup the http service port")
	flag.BoolVar(args.save, "save", true, "setup if to enable the persistence module")
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

	}

	fmt.Printf("[main.log] main.callQuit(): quited with code 114\n")
	os.Exit(114)
}

func init() {
	fmt.Printf("[main.log] main.init(): initializing main package\n")
	signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)
	fmt.Printf("[main.log] main.init(): setuped signal channel\n")
	initArgs()
	flag.Parse()
	fmt.Printf("[main.log] main.init(): parsed command line arguments\n")
	_ = persistence.SetStoragePath("/Users/sunist/Projects/GitHub/gisb/data/persistence")
	ch := make(chan error, 10)
	persistence.Load(ch)
	fmt.Printf("[main.log] main.init(): initializing persistent storage\n")
	go func() {
		for err := range ch {
			fmt.Println(err)
		}
	}()
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
