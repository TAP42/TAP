package game

// Player is someone in the game: their name and the room they are in.
// The room is guarded by the registry lock: read and write it only
// through Registry methods.
type Player struct {
	Name string
	room string
}
