package game

import "github.com/TAP42/TAP/world"

// View is what a player sees when they look around.
type View struct {
	RoomID  string
	Room    world.Location
	Players []string
	Items   []string
	NPCs    []string
}

// Look returns what player sees in the room they are in.
func (s *State) Look(player *Player) View {
	roomID, players := s.Players.surroundings(player)
	room := s.World.Locations[roomID]

	npcs := []string{}
	for _, spawn := range room.Spawns {
		npcs = append(npcs, spawn.NPCType)
	}

	return View{
		RoomID:  roomID,
		Room:    room,
		Players: players,
		Items:   append([]string{}, room.Items...),
		NPCs:    npcs,
	}
}

// surroundings returns the room of player and the names of everyone in
// it, read under one lock so the two always match.
func (r *Registry) surroundings(player *Player) (string, []string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return player.room, r.namesInLocked(player.room)
}
