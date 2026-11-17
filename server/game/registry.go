package game

import (
	"sync"
)

// Registry holds every connected player, indexed by name.
type Registry struct {
	mu      sync.Mutex
	players map[string]*Player
}

// startRoom is where every new player appears.
const startRoom = "start"

// newRegistry returns an empty registry ready to use.
func newRegistry() *Registry {
	return &Registry{players: make(map[string]*Player)}
}

// Add registers a new player under name, or fails if the name is taken.
func (r *Registry) Add(name string) (*Player, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, present := r.players[name]; present {
		return nil, ErrNameInUse
	}
	player := &Player{Name: name, room: startRoom}
	r.players[name] = player
	return player, nil
}

// Remove frees the name of a player who left.
func (r *Registry) Remove(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.players, name)
}

// NamesIn returns the names of the players currently in room.
func (r *Registry) NamesIn(room string) []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.namesInLocked(room)
}

// namesInLocked is NamesIn for callers that already hold r.mu.
func (r *Registry) namesInLocked(room string) []string {
	names := []string{}
	for _, user := range r.players {
		if user.room == room {
			names = append(names, user.Name)
		}
	}
	return names
}
