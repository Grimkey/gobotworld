package terrain

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestTerrain(t *testing.T) {
	myTerrain := Dirt
	assert.Assert(t, myTerrain == Dirt)
	assert.Assert(t, myTerrain != Obstacle)
}
