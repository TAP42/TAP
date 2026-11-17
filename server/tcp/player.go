package tcp

import (
	"sync"

	"github.com/TAP42/TAP/game"
	"github.com/TAP42/TAP/world"
)

// Player is a connected client that has chosen a name with CONNECT.
type Player struct {
	name string
	room string
}

// game is the state shared by every connection.
type gameState struct {
	players *registry
	world   *world.World
}

// registry holds every connected player, indexed by name.
type registry struct {
	mu      sync.Mutex
	players map[string]*Player
}

// newRegistry returns an empty registry ready to use.
func newRegistry() *registry {
	return &registry{players: make(map[string]*Player)}
}

// startRoom is where every new player appears.
const startRoom = "start"

// add registers a new player under name, or fails if the name is taken.
func (r *registry) add(name string) (*Player, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, present := r.players[name]; present {
		return nil, game.ErrNameInUse
	}
	player := &Player{name: name, room: startRoom}
	r.players[name] = player
	return player, nil
}

// remove frees the name of a player who left.
func (r *registry) remove(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.players, name)
}
