package object_test

import (
	"gobotworld/src/world/object"
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectionString(t *testing.T) {
	tests := []struct {
		direction object.Direction
		expected  string
	}{
		{object.North, "North"},
		{object.West, "West"},
		{object.South, "South"},
		{object.East, "East"},
		{object.Direction(99), "Unknown"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.direction.String(), "Direction string should match expected value")
	}
}

func TestNewCharacter(t *testing.T) {
	start := image.Point{X: 5, Y: 5}
	char := object.NewPlayer(start)

	assert.NotNil(t, char, "NewPlayer should not return nil")
	assert.Equal(t, object.PlayerType, char.Ident().Type, "NewPlayer should create a character with PlayerType")
	assert.Equal(t, 99, char.Ident().Index, "NewPlayer should create a character with index 99")
	assert.Equal(t, &start, char.Location, "Character location should match the provided start point")
	assert.Equal(t, object.North, char.Direction, "New characters should default to North direction")
}

func TestNewPlayer(t *testing.T) {
	start := image.Point{X: 10, Y: 10}
	player := object.NewPlayer(start)

	assert.NotNil(t, player, "NewPlayer should return a valid Character")
	assert.Equal(t, object.PlayerType, player.Ident().Type, "Player type should be PlayerType")
	assert.Equal(t, 99, player.Ident().Index, "Player index should be 99")
	assert.Equal(t, &start, player.Location, "Player location should match the provided start point")
}

func TestNewNPC(t *testing.T) {
	start := image.Point{X: 15, Y: 15}
	npc := object.NewNPC(start)

	assert.NotNil(t, npc, "NewNPC should return a valid Character")
	assert.Equal(t, object.EnemyType, npc.Ident().Type, "NPC type should be EnemyType")
	assert.Equal(t, 20, npc.Ident().Index, "NPC index should be 20")
	assert.Equal(t, &start, npc.Location, "NPC location should match the provided start point")
}

func TestCharacterIdent(t *testing.T) {
	start := image.Point{X: 0, Y: 0}
	char := object.NewPlayer(start)

	ident := char.Ident()
	assert.Equal(t, object.PlayerType, ident.Type, "Ident type should be PlayerType")
	assert.Equal(t, 99, ident.Index, "Ident index should be 99")
}

func TestCharacterPassable(t *testing.T) {
	start := image.Point{X: 0, Y: 0}
	char := object.NewPlayer(start)

	passable := char.Passable(nil)
	assert.False(t, passable, "Characters should not be passable")
}
