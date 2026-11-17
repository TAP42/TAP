package tcp

import "github.com/TAP42/TAP/game"

// handleLook describes the room the player is in.
func handleLook(sess *session, state *game.State) error {
	room := state.World.Locations[sess.player.Room]
	return writeLine(sess.conn, "OK "+room.Name)
}
