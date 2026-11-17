package tcp

// handleLook describes the room the player is in.
func handleLook(sess *session, state *gameState) error {
	room := state.world.Locations[sess.player.Room]
	return writeLine(sess.conn, "OK "+room.Name)
}
