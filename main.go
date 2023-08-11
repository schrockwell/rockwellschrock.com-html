package main

import (
	"fmt"
	"os"
	"yass/internal/generator"
	"yass/internal/resizer"
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

	if arg == "queen" {
		generator.Run(siteDirPath, outDirPath)
	} else if arg == "convert" {
		resizer.Run(siteDirPath)
	} else if arg == "server" {
		server.Start(outDirPath)
	} else {
		printHelp()
	}
}

func printHelp() {
	fmt.Print(`Usage: yass [command]

Commands:

    convert     Reformat and rename JPEG images in-place in web/
    queen       Generate the site in _site/
    server      Run a local web server for _site/
`)
}
