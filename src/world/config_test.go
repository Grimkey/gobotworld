package world

import (
	"gobotworld/src/world/object"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	obj1 := object.NewObject(1, object.Dirt1Type, true)
	obj2 := object.NewObject(2, object.RockType, true)

	config := NewConfig(
		DefaultTerrain{Type: obj1, Units: 3},
		DefaultTerrain{Type: obj2, Units: 7},
	)

	assert.Len(t, config.terrainTypes, 2, "Config should contain 2 terrain types")
	assert.Equal(t, 10, config.terrainSum, "Config should calculate the correct terrain sum")

	assert.Equal(t, 3, config.terrainTypes[0].Units, "First terrain type should have correct cumulative units")
	assert.Equal(t, 10, config.terrainTypes[1].Units, "Second terrain type should have correct cumulative units")
}

func TestRandomObject(t *testing.T) {
	obj1 := object.NewObject(1, object.Dirt1Type, true)
	obj2 := object.NewObject(2, object.RockType, true)

	config := NewConfig(
		DefaultTerrain{Type: obj1, Units: 3},
		DefaultTerrain{Type: obj2, Units: 7},
	)

	rand.New(rand.NewSource(42)) // Fixed seed for deterministic behavior
	randValue := rand.Intn(config.terrainSum)
	config.rndObjextIndex = func(_ int) int { return randValue }
	randomObj := config.RandomObject()

	if randValue < 3 {
		assert.Equal(t, obj1.Ident(), randomObj.Ident(), "RandomObject should return obj1 for values less than 3")
	} else {
		assert.Equal(t, obj2.Ident(), randomObj.Ident(), "RandomObject should return obj2 for values greater than or equal to 3")
	}
}

func TestRandomObjectPanicsOnInvalidConfig(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Contains(t, r, "no terrain types", "Should panic with correct error message when terrainSum is invalid")
		}
	}()

	// Empty config to trigger panic
	config := NewConfig()
	_ = config.RandomObject()
}
