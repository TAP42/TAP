// package tcp.
package tcp

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/TAP42/TAP/world"
)

// listenAddr binds every interface so players on other machines can join.
const listenAddr = ":4242"

// TCP listens on listenAddr and serves each client in its own goroutine.
// It only returns when the listener cannot be created.
func TCP(gameWorld *world.World) error {
	//nolint:gosec // public on purpose, see listenAddr.
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return fmt.Errorf("listen on %s: %w", listenAddr, err)
	}
	defer closeOrLog(listener, "listener")

	state := &gameState{players: newRegistry(), world: gameWorld}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept connection: %v", err)
			continue
		}
		go handleConnection(conn, state)
	}
}

// closeOrLog closes c and logs the error, if any.
func closeOrLog(c io.Closer, name string) {
	if err := c.Close(); err != nil {
		log.Printf("close %s: %v", name, err)
	}
}
