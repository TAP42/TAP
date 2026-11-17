package tcp

import "github.com/TAP42/TAP/game"

// handleCommand answers one command and reports whether the client leaves.
func handleCommand(
	sess *session,
	command string,
	args string,
	state *game.State,
) (bool, error) {
	switch command {
	case "QUIT":
		return true, writeLine(sess.conn, byeReply)
	case "CONNECT":
		return false, handleConnect(sess, args, state)
	}
	if sess.player == nil {
		return false, writeLine(sess.conn, noConnected)
	}
	switch command {
	case "LOOK":
		return false, handleLook(sess, state)
	default:
		return false, writeLine(sess.conn, unknownCommand)
	}
}
