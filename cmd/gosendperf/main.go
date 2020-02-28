package main

import (
	"fmt"
	"strings"

	"github.com/jessevdk/go-flags"
)

type params struct {
	Client  string `short:"c" long:"client" description:"client mode"`
	Server  bool   `short:"s" long:"server" description:"server mode"`
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
}

func main() {
	var p params

	fmt.Println("gosendperf started")

	parser := flags.NewParser(&p, flags.Default)
	//parser.Usage = "custom usage"
	args, err := parser.Parse()
	if err != nil {
		flagsErr, ok := err.(*flags.Error)
		if !(ok && flagsErr.Type == flags.ErrHelp) {
			panic(err)
		}
	}

	fmt.Println("Remaing args: ", strings.Join(args, ", "))

	if p.Server {
		server()
	} else {
		if len(p.Client) == 0 {
			panic("client address required or use -s")
		}

		client()
	}
}
