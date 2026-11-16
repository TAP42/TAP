package tcp

import (
	"bufio"
	"log"
	"net"
)

// handleConnection greets one client, then answers its commands until it
// quits or disconnects.
func handleConnection(conn net.Conn, state *game) {
	defer closeOrLog(conn, "connection")
	if err := writeLine(conn, greeting); err != nil {
		log.Printf("write greeting: %v", err)
		return
	}
	sess := &session{conn: conn}
	defer sess.leave(state.players)
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("read: %v", err)
			return
		}
		command, args := parseCommand(message)
		quit, err := handleCommand(sess, command, args, state)
		if err != nil {
			log.Printf("write: %v", err)
			return
		}
		if quit {
			return
		}
	}
}
