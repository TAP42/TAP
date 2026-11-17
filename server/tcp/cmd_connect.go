package tcp

import (
	"strings"

	"github.com/TAP42/TAP/game"
)

// handleConnect gives the session a player name, once, if the name is free.
func handleConnect(sess *session, name string, state *game.State) error {
	if sess.player != nil {
		return writeLine(sess.conn, alreadyConnected)
	}
	if name == "" || strings.ContainsAny(name, " \t") {
		return writeLine(sess.conn, invalidName)
	}
	player, err := state.Players.Add(name)
	if err != nil {
		return writeLine(sess.conn, nameInUse)
	}
	sess.player = player
	return writeLine(sess.conn, connectedReply)
}
