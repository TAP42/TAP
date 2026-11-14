package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/TAP42/TAP/client/commands"
)

func main() {
	conn, err := dial("localhost:4242")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	defer conn.Close()

	if err := commands.Cli(conn); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// dial connects to the game server and prints every line it sends.
func dial(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("connect to server: %w", err)
	}
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	return conn, nil
}
