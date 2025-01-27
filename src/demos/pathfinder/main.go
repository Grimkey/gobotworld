package main

import (
	"gobotworld/src/world"
	"image"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "LOG: ", log.LstdFlags)
	wld := world.DefaultWorld(logger) // Create the world with default settings

	pathFinder := world.PathFinder{World: wld, Logger: logger}

	start := image.Point{X: 1, Y: 1}
	dest := image.Point{X: 10, Y: 10}

	path := pathFinder.Find(start, dest)
	if len(path) == 0 {
		logger.Println("No path found.")
	} else {
		logger.Println("Path found:")
		for point := range path {
			logger.Printf("  -> %v", point)
		}
	}
}
