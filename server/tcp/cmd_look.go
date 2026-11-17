package tcp

import "github.com/TAP42/TAP/game"

// handleLook describes the room the player is in.
func handleLook(sess *session, state *game.State) error {
	view := state.Look(sess.player)
	return writeLine(sess.conn, "OK "+view.Room.Name)
}
