package tcp

import (
	"bufio"
	"log"
	"net"

	"github.com/TAP42/TAP/game"
)

// session is the state of one TCP connection.
type session struct {
	conn   net.Conn
	player *game.Player // nil until CONNECT succeeds
}

// handleConnection greets one client, then answers its commands until it
// quits or disconnects.
func handleConnection(conn net.Conn, state *gameState) {
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

// leave frees the session's player name, if it has one.
func (sess *session) leave(players *registry) {
	if sess.player != nil {
		players.remove(sess.player.Name)
	}
}
