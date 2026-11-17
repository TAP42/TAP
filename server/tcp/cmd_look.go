package tcp

import (
	"encoding/json"
	"fmt"

	"github.com/TAP42/TAP/game"
)

// lookReply is the JSON body of a LOOK response (RFC 42TAP 5.1.2).
type lookReply struct {
	Room    lookRoom `json:"room"`
	Players []string `json:"players"`
	Items   []string `json:"items"`
	NPCs    []string `json:"npcs"`
}

// lookRoom is the "room" object of a LOOK response.
type lookRoom struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Exits       map[string]string `json:"exits"`
}

// handleLook sends the room the player is in, as JSON.
func handleLook(sess *session, state *game.State) error {
	data, err := json.Marshal(newLookReply(state.Look(sess.player)))
	if err != nil {
		return fmt.Errorf("encode LOOK reply: %w", err)
	}
	return writeLine(sess.conn, "OK "+string(data))
}

// newLookReply copies a game view into the shape the RFC expects.
func newLookReply(view game.View) lookReply {
	exits := view.Room.Exits
	if exits == nil {
		exits = map[string]string{}
	}
	return lookReply{
		Room: lookRoom{
			ID:          view.RoomID,
			Name:        view.Room.Name,
			Description: view.Room.Description,
			Exits:       exits,
		},
		Players: view.Players,
		Items:   view.Items,
		NPCs:    view.NPCs,
	}
}
