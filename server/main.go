package main

import (
	"fmt"
	"os"

	"github.com/TAP42/TAP/commands"
)

func main() {
	if err := commands.Cli(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
