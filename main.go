package main

import (
	"fmt"
	"os"
	"yass/internal/generator"
	"yass/internal/server"
)

const (
	siteDirPath = "./web"
	outDirPath  = "./_site"
)

func main() {
	arg := ""

	if len(os.Args) >= 2 {
		arg = os.Args[1]
	}

	if arg == "server" {
		server.Start(outDirPath)
	} else if arg == "queen" {
		generator.Run(siteDirPath, outDirPath)
	} else {
		printHelp()
	}
}

func printHelp() {
	fmt.Print(`Usage: yass [command]

Commands:

    queen     Generate the site in _site/
    server    Run a local web server for _site/
`)
}
