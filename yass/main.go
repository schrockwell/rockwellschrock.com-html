package main

import (
	"fmt"
	"os"
	"schrockwell/yass/generator"
	"schrockwell/yass/server"
)

func main() {
	arg := ""

	if len(os.Args) >= 2 {
		arg = os.Args[1]
	}

	if arg == "server" {
		server.Start()
	} else if arg == "build" {
		generator.Run()
	} else {
		printHelp()
	}
}

func printHelp() {
	fmt.Print(`Usage: yass [command]

Commands:

    build     Generate the site in _out/
    server    Run a local web server for _out/
`)
}
