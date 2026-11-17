package tcp

import "strings"

// handleConnect gives the session a player name, once, if the name is free.
func handleConnect(sess *session, name string, state *gameState) error {
	if sess.player != nil {
		return writeLine(sess.conn, alreadyConnected)
	}
	if name == "" || strings.ContainsAny(name, " \t") {
		return writeLine(sess.conn, invalidName)
	}
	player, err := state.players.add(name, sess.conn)
	if err != nil {
		return writeLine(sess.conn, nameInUse)
	}
	sess.player = player
	return writeLine(sess.conn, connectedReply)
}
