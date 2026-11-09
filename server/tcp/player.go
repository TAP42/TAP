package tcp

import (
	"errors"
	"net"
	"sync"
)

// Player is a connected client that has chosen a name with CONNECT.
type Player struct {
	name string
	room string
	conn net.Conn
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

func (r *registry) add(name string, conn net.Conn) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	errNameInUse := errors.New("Name already in use")
	if _, present := r.players[name]; present {
		return errNameInUse
	}
	r.players[name] = &Player{name: name, conn: conn}
	return nil
}
