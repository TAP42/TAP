package game

import (
	"github.com/TAP42/TAP/world"
)

// State holds everything the running game knows: its players and its world.
type State struct {
	Players *Registry
	World   *world.World
}

// NewState returns an empty game state for gameWorld, ready to use.
func NewState(gameWorld *world.World) *State {
	state := &State{Players: newRegistry(), World: gameWorld}
	return state
}
