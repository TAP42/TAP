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

// errNameInUse is returned by add when another player holds the name.
var errNameInUse = errors.New("name already in use")

// session is the state of one TCP connection.
type session struct {
	conn   net.Conn
	player *Player // nil until CONNECT succeeds
}

// add registers a new player under name, or fails if the name is taken.
func (r *registry) add(name string, conn net.Conn) (*Player, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, present := r.players[name]; present {
		return nil, errNameInUse
	}
	player := &Player{name: name, conn: conn}
	r.players[name] = player
	return player, nil
}

// remove frees the name of a player who left.
func (r *registry) remove(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.players, name)
}

// leave frees the session's player name, if it has one.
func (sess *session) leave(players *registry) {
	if sess.player != nil {
		players.remove(sess.player.name)
	}
}
