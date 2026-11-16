package tcp

func handleLook(sess *session, state *game) error {
	room := state.world.Locations[sess.player.room]
	return writeLine(sess.conn, "OK "+room.Name)
}
