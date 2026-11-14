// Package world loads the game world from a YAML file.
package world

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// File is the root of world.yaml.
type File struct {
	World World `yaml:"world"`
}

// World holds every location, item and NPC of the game.
type World struct {
	Locations map[string]Location `yaml:"locations"`
	Items     map[string]Item     `yaml:"items"`
	NPCs      map[string]NPC      `yaml:"npcs"`
}

// Location is one room of the world.
type Location struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Exits       map[string]string `yaml:"exits"`
	Items       []string          `yaml:"items"`
	Spawns      []Spawn           `yaml:"spawns"`
}

// Spawn tells how many NPCs of one type appear in a location.
type Spawn struct {
	NPCType string `yaml:"npc_type"`
	Count   int    `yaml:"count"`
}

// Item is an object players can find in the world.
type Item struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Obtainable  bool   `yaml:"obtainable"`
}

// NPC is a non-player character template.
type NPC struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Dialogue    []string `yaml:"dialogue"`
	Stats       Stats    `yaml:"stats"`
}

// Stats holds the combat values of an NPC.
type Stats struct {
	HP int `yaml:"hp"`
}

func Load(path string) (*World, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read wolrd %w", err)
	}
	var file File
	if err := yaml.Unmarshal(data, &file); err != nil {
		return nil, fmt.Errorf("decode world: %w", err)
	}
	return &file.World, nil
}
