package world

import (
	"image"
	"log"

	"github.com/fzipp/astar"
)

type PathFinder struct {
	World  World
	Logger *log.Logger
}

func (pf PathFinder) Find(start, dest image.Point) map[image.Point]bool {
	path := make(map[image.Point]bool)
	points := astar.FindPath[image.Point](pf.World, start, dest, manhattanDistance, manhattanDistance)

	for _, point := range points {
		path[point] = true
	}

	return path
}

func manhattanDistance(p, q image.Point) float64 {
	return float64(abs(p.X-q.X) + abs(p.Y-q.Y))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
