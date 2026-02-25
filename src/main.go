package main

import (
	"os"

	"github.com/keypad/trace/src/core"
)

func main() {
	os.Exit(core.Run(os.Args[1:], os.Stdout, os.Stderr))
}
