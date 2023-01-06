package main

import (
	"flag"
	"github.com/sunist-c/genius-invokation-simulator-backend/exec/backend"
	"github.com/sunist-c/genius-invokation-simulator-backend/exec/cli"
	"os"
)

var (
	args *argument = &argument{
		mode: new(string),
		port: new(uint),
		conf: new(string),
		save: new(bool),
	}
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

func init() {
	initArgs()
}

func main() {
	flag.Parse()
	switch *args.mode {
	case "backend":
		backend.Run(*args.port)
	case "cli":
		cli.Run()
	case "ai":
		panic("not implemented yet")
	default:
		os.Exit(0)
	}
}
