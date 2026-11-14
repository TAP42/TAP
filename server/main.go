package main

import (
	"fmt"
	"os"

	"github.com/TAP42/TAP/tcp"
	"github.com/TAP42/TAP/world"
)

func main() {
	w, err := world.Load("data/world.yaml")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", w.Locations["start"])

	if err := tcp.TCP(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
