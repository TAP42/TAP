// package tcp.
package tcp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

// listenAddr binds every interface so players on other machines can join.
const listenAddr = ":4242"

// TCP listens on listenAddr and serves each client in its own goroutine.
// It only returns when the listener cannot be created.
func TCP() error {
	//nolint:gosec // public on purpose, see listenAddr.
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return fmt.Errorf("listen on %s: %w", listenAddr, err)
	}
	defer closeOrLog(listener, "listener")

	players := newRegistry()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept connection: %v", err)
			continue
		}
		go handleConnection(conn, players)
	}
}

// closeOrLog closes c and logs the error, if any.
func closeOrLog(c io.Closer, name string) {
	if err := c.Close(); err != nil {
		log.Printf("close %s: %v", name, err)
	}
}

// handleConnection greets one client, then answers its commands until it
// quits or disconnects.
func handleConnection(conn net.Conn, players *registry) {
	defer closeOrLog(conn, "connection")
	if err := writeLine(conn, greeting); err != nil {
		log.Printf("write greeting: %v", err)
		return
	}
	sess := &session{conn: conn}
	defer sess.leave(players)
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("read: %v", err)
			return
		}
		command, args := parseCommand(message)
		quit, err := handleCommand(sess, command, args, players)
		if err != nil {
			log.Printf("write: %v", err)
			return
		}
		if quit {
			return
		}
	}
}
