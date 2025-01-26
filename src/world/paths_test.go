package world

import (
	"image"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathFinderFindPath(t *testing.T) {
	// Initialize a default world
	logger := log.New(os.Stdout, "", log.LstdFlags)
	worldInstance := EmptyWorld(logger)
	pathFinder := PathFinder{World: worldInstance}

	// Define start and destination points
	start := image.Point{X: 1, Y: 1}
	dest := image.Point{X: 3, Y: 1}

	// Clear obstacles between start and dest
	worldInstance.Geography.SetLoc(start, nil)
	worldInstance.Geography.SetLoc(image.Point{X: 2, Y: 2}, nil)
	worldInstance.Geography.SetLoc(dest, nil)

	// Find path
	path := pathFinder.Find(start, dest)

	// Expected path
	expectedPath := map[image.Point]bool{
		{1, 1}: true,
		{2, 1}: true,
		{3, 1}: true,
	}

	assert.Equal(t, expectedPath, path, "PathFinder should find the correct path from start to destination")
}

func TestPathFinderStartEqualsDest(t *testing.T) {
	// Initialize a default world
	logger := log.New(os.Stdout, "", log.LstdFlags)
	worldInstance := DefaultWorld(logger)
	pathFinder := PathFinder{World: worldInstance}

	// Define start and destination as the same point
	start := image.Point{X: 2, Y: 2}

	// Ensure the start point is passable
	worldInstance.Geography.SetLoc(start, nil)

	// Find path
	path := pathFinder.Find(start, start)

	// Expected path contains only the start point
	expectedPath := map[image.Point]bool{
		{2, 2}: true,
	}

	assert.Equal(t, expectedPath, path, "PathFinder should return a path with only the start point when start equals destination")
}

func TestManhattanDistance(t *testing.T) {
	// Test Manhattan distance calculation
	p := image.Point{X: 0, Y: 0}
	q := image.Point{X: 3, Y: 4}

	distance := manhattanDistance(p, q)

	assert.Equal(t, float64(7), distance, "Manhattan distance between (0,0) and (3,4) should be 7")
}
