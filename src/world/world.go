package world

import (
	"gobotworld/src/world/character"
	"gobotworld/src/world/terrain"
	"math/rand"
)

const (
	Width  = 200
	Height = 200
)

type Point struct {
	X int
	Y int
}

type Cell struct {
	Terrain terrain.Terrain
	Light   int
}

type World struct {
	Geography [][]*Cell
	Beings    map[character.Character]*Point
}

type DefaultTerrain struct {
	Type  terrain.Terrain
	Units int
}

type Config struct {
	terrainTypes []DefaultTerrain
	terrainSum   int
}

func NewConfig(t ...DefaultTerrain) Config {
	sum := 0
	internalTerrain := make([]DefaultTerrain, 0, len(t))
	for _, item := range t {
		sum += item.Units
		internalTerrain = append(internalTerrain, DefaultTerrain{item.Type, sum})
	}
	return Config{
		terrainTypes: internalTerrain,
		terrainSum:   sum,
	}
}

func (c Config) RandomTerrain() terrain.Terrain {
	rnd := rand.Intn(c.terrainSum)
	for _, v := range c.terrainTypes {
		if rnd < v.Units {
			return v.Type
		}
	}
	return c.terrainTypes[len(c.terrainTypes)-1].Type
}

func DefaultWorld() World {
	cfg := NewConfig(
		DefaultTerrain{terrain.Empty, 20},
		DefaultTerrain{terrain.Dirt, 7},
		DefaultTerrain{terrain.Rock, 3},
		DefaultTerrain{terrain.Obstacle, 1},
	)

	return InitWorld(Height, Width, cfg)
}

func InitWorld(height, width int, cfg Config) World {
	geography := make([][]*Cell, height, height)

	for i := range geography {
		geography[i] = make([]*Cell, 0, width)
		for j := 0; j < Width; j++ {
			cell := Cell{
				Terrain: cfg.RandomTerrain(),
				Light:   0,
			}
			geography[i] = append(geography[i], &cell)
		}
	}

	playerLocation := Point{Width / 2, Height / 2}

	return World{
		Geography: geography,
		Beings:    map[character.Character]*Point{character.Player: &playerLocation},
	}
}

func (world World) Height() int {
	return Height
}

func (world World) Width() int {
	return Width
}

func (world World) At(point Point) *Cell {
	if point.Y < 0 || point.X < 0 {
		return nil
	}
	if point.Y >= len(world.Geography) {
		return nil
	}
	if point.X >= len(world.Geography[point.Y]) {
		return nil
	}
	return world.Geography[point.Y][point.X]
}

func (world World) Move(char character.Character, deltaX int, deltaY int) {
	location := world.Beings[char]
	proposed := Point{X: location.X + deltaX, Y: location.Y + deltaY}

	if world.At(proposed) != nil && world.At(proposed).Terrain.Passable {
		world.Beings[char] = &proposed
	}
}
