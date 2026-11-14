package main

import (
	"fmt"
	"os"

	"github.com/TAP42/TAP/tcp"
	"github.com/TAP42/TAP/world"
)

func main() {
	gameWorld, err := world.Load("data/world.yaml")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	if err := tcp.TCP(gameWorld); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
