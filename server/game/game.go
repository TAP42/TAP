package game

import (
	"github.com/TAP42/TAP/world"
)

// State is the state shared by every connection.
type State struct {
	Players *Registry
	World   *world.World
}

func NewState(gameWorld *world.World) *State {
	state := &State{Players: newRegistry(), World: gameWorld}
	return state
}
