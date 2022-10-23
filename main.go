package main

import (
	"os"

	"gitlab.com/tokene/keyserver-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
