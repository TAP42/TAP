package main

import (
	"fmt"
	"os"

	"github.com/TAP42/TAP/tcp"
)

func main() {
	if err := tcp.TCP(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
