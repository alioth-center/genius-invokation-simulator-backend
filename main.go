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
	s := <-sig
	persistence.Quit()

	switch *args.mode {
	case backendMode:
		backend.Quit()
	case cliMode:
		cli.Quit()
	case aiMode:

	}

	fmt.Println("quit called:", s.String())
	os.Exit(114)
}

func init() {
	signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)
	initArgs()
	flag.Parse()
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
