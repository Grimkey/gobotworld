package world

import (
	"gobotworld/src/world/object"
	"image"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomMap(t *testing.T) {
	cfg := NewConfig(
		DefaultTerrain{object.NewObject(0, object.Dirt1Type, true), 100},
		DefaultTerrain{object.NewObject(0, object.TorchType, false), 5},
	)
	cfg.rndObjextIndex = func(_ int) int { return 1 }

	geography, lights := RandomMap(10, 10, cfg)

	assert.Equal(t, 10, geography.Height(), "RandomMap should create a map with the correct height")
	assert.Equal(t, 10, geography.Width(), "RandomMap should create a map with the correct width")
	assert.LessOrEqual(t, len(lights), 6, "Number of torch lights should not exceed the expected count")
}

func TestWorldInitialization(t *testing.T) {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	worldInstance := DefaultWorld(logger)

	assert.NotNil(t, worldInstance.Player, "Player should not be nil")
	assert.NotNil(t, worldInstance.Geography, "Geography should not be nil")
	assert.Greater(t, len(worldInstance.Lights), 0, "Lights should be initialized")
	assert.Contains(t, worldInstance.Beings, worldInstance.Player, "Player should be in the list of beings")
}

func TestMoveCharacter(t *testing.T) {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	worldInstance := DefaultWorld(logger)

	player := worldInstance.Player
	startLocation := *player.Location
	moveDirection := East

	moved := worldInstance.Move(player, moveDirection)
	assert.True(t, moved, "Player should be able to move East")
	assert.Equal(t, image.Point{X: startLocation.X + 1, Y: startLocation.Y}, *player.Location, "Player location should be updated correctly")

	// Test movement into an obstacle
	eastObstacle := image.Point{X: player.Location.X + 1, Y: player.Location.Y}
	worldInstance.Geography.SetLoc(eastObstacle, object.ThingList{object.NewObject(0, object.ObstacleType, false)})
	moved = worldInstance.Move(player, moveDirection)
	assert.False(t, moved, "Player should not move into an obstacle")
}

func TestNpcMove(t *testing.T) {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	worldInstance := DefaultWorld(logger)

	enemy := &Character{
		Location:  &image.Point{X: 10, Y: 10},
		Direction: North,
	}
	worldInstance.Beings[enemy] = false

	worldInstance.NpcMove()
	assert.NotEqual(t, image.Point{X: 10, Y: 10}, *enemy.Location, "NPC should move from its starting position")
}

func TestNeighbours(t *testing.T) {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	worldInstance := EmptyWorld(logger)

	player := worldInstance.Player
	startLocation := *player.Location

	neighbours := []image.Point{}
	worldInstance.Neighbours(startLocation)(func(p image.Point) bool {
		neighbours = append(neighbours, p)
		return true
	})

	expectedNeighbours := []image.Point{
		{X: startLocation.X, Y: startLocation.Y - 1}, // North
		{X: startLocation.X + 1, Y: startLocation.Y}, // East
		{X: startLocation.X, Y: startLocation.Y + 1}, // South
		{X: startLocation.X - 1, Y: startLocation.Y}, // West
	}

	assert.ElementsMatch(t, expectedNeighbours, neighbours, "Neighbours should match the expected surrounding points")
}

func TestTimeCycle(t *testing.T) {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	worldInstance := DefaultWorld(logger)

	cycle, ticks := worldInstance.Time()
	assert.Equal(t, object.DayTime, cycle, "Initial time cycle should be DayTime")
	assert.Equal(t, 0, ticks, "Initial time ticks should be 0")

	worldInstance.Tick()

	cycle, cnt := worldInstance.Time()
	assert.Equal(t, 1, *worldInstance.time, "Cycle should be after the first tick")
	assert.Equal(t, 0, cnt, "Cnt only goes up for full days")
	assert.Equal(t, object.DayTime, cycle, "Time cycle should stay DayTime")

	// Advance time to switch to NightTime
	moveHalfDay := 44
	for i := 0; i < moveHalfDay; i++ {
		worldInstance.Tick()
	}

	cycle, cnt = worldInstance.Time()
	assert.Equal(t, 45, *worldInstance.time, "Correct calculation for ticks")
	assert.Equal(t, 1, cnt, "Cnt is how far into the day we are")
	assert.Equal(t, object.NightTime, cycle, "Time cycle should switch to NightTime after enough ticks")
}
