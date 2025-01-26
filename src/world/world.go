package world

import (
	"gobotworld/src/geometry"
	"gobotworld/src/world/object"
	"image"
	"iter"
	"log"
	"math/rand"
)

const (
	Width  = 200
	Height = 200

	ticksPerCount   = 4
	countPerDay     = 20
	countPerHalfDay = countPerDay / 2
)

type Lights []*image.Point

func (lts Lights) NearestLight(p image.Point) image.Point {
	var nearest image.Point
	var nearestDist = 0
	for _, light := range lts {
		d := geometry.Distance(*light, p)
		if d < nearestDist {
			nearest = *light
		}
	}

	return nearest
}

type Map [][]object.ThingList

func (m Map) At(point image.Point) object.ThingList {
	if point.Y < 0 || point.X < 0 {
		return nil
	}
	if point.Y >= len(m) {
		return nil
	}
	if point.X >= len(m[point.Y]) {
		return nil
	}
	return m[point.Y][point.X]
}

func (m Map) SetLoc(p image.Point, things object.ThingList) {
	m[p.Y][p.X] = things
}

func (m Map) RemoveLoc(p image.Point, thing object.Thing) {
	m.SetLoc(p, m.At(p).DeleteItem(thing))
}

func (m Map) AddLoc(p image.Point, thing object.Thing) {
	m.SetLoc(p, append(m.At(p), thing))
}

func (m Map) Height() int {
	return len(m)
}

func (m Map) Width() int {
	return len(m[0])
}

func (m Map) CanPass(point image.Point, thing object.Thing) bool {
	objs := m.At(point)
	for _, obj := range objs {
		if !obj.Passable(thing) {
			return false
		}
	}
	return true
}

func RandomMap(height, width int, cfg Config) (Map, Lights) {
	geography := make(Map, height)
	var lights Lights
	for i := range geography {
		geography[i] = make([]object.ThingList, 0, width)
		for j := 0; j < width; j++ {
			rndObj := cfg.RandomObject()
			if rndObj.Ident().Type == object.TorchType {
				lights = append(lights, &image.Point{X: j, Y: i})
			}
			geography[i] = append(geography[i], object.ThingList{rndObj})
		}
	}
	return geography, lights
}

type World struct {
	logger    *log.Logger
	Geography Map
	Player    *Character
	Beings    map[*Character]bool
	Lights    Lights
	time      *int
}

func EmptyWorld(logger *log.Logger) World {
	cfg := NewConfig(
		DefaultTerrain{object.NewObject(0, object.Dirt1Type, true), 400},
		DefaultTerrain{object.NewObject(0, object.Dirt2Type, true), 250},
		DefaultTerrain{object.NewObject(0, object.RockType, true), 50},
	)
	return InitWorld(logger, Height, Width, cfg)
}

func DefaultWorld(logger *log.Logger) World {
	cfg := NewConfig(
		DefaultTerrain{object.NewObject(0, object.Dirt1Type, true), 400},
		DefaultTerrain{object.NewObject(0, object.Dirt2Type, true), 250},
		DefaultTerrain{object.NewObject(0, object.RockType, true), 50},
		DefaultTerrain{object.NewObject(0, object.ObstacleType, false), 5},
		DefaultTerrain{object.NewObject(0, object.TorchType, false), 1},
	)

	return InitWorld(logger, Height, Width, cfg)
}

func InitWorld(logger *log.Logger, height, width int, cfg Config) World {
	geography, lights := RandomMap(height, width, cfg)

	playerLocation := image.Point{X: width / 2, Y: height / 2}
	enemyLocation := image.Point{X: 10, Y: 10}
	start := 0
	player := NewPlayer(playerLocation)
	enemy := NewNPC(enemyLocation)

	geography.AddLoc(playerLocation, player)
	geography.AddLoc(enemyLocation, enemy)

	logger.Print("Working with a map of size ", geography.Height(), "x", geography.Width())
	return World{
		logger:    logger,
		Geography: geography,
		Lights:    lights,
		Player:    player,
		Beings: map[*Character]bool{
			player: true,
			enemy:  false,
		},
		time: &start,
	}
}

func (world World) Time() (object.DayCycle, int) {
	now := (*world.time / ticksPerCount) % countPerDay
	cycle := object.DayTime
	if now > countPerHalfDay {
		cycle = object.NightTime
	}
	return cycle, now % countPerHalfDay
}

func (world World) Tick() {
	*world.time += 1
}

var directions = []Direction{North, South, East, West}

func (world World) NpcMove() {
	for being, isPlayer := range world.Beings {
		if isPlayer {
			continue
		}

		// Shuffle the directions
		rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })

		// Try to move in each direction
		for _, direction := range directions {
			if moved := world.Move(being, direction); moved {
				// If the move was successful, stop trying to move
				world.logger.Printf("NPC %d moved %s", being.Ident().Index, direction.String())
				break
			}
		}
	}
}

func (world World) Neighbours(p image.Point) iter.Seq[image.Point] {
	return func(yield func(image.Point) bool) {
		offsets := []image.Point{
			{0, -1}, // North
			{1, 0},  // East
			{0, 1},  // South
			{-1, 0}, // West
		}

		for _, off := range offsets {
			q := p.Add(off)
			if world.Geography.CanPass(q, world.Player) {
				// I find iterators a little tricky. This means keep yielding until we get back a false which means the caller is done iterating.
				if !yield(q) {
					return
				}
			}
		}
	}
}

var moveTransform = map[Direction]image.Point{
	North: {X: 0, Y: -1},
	West:  {X: -1, Y: 0},
	South: {X: 0, Y: 1},
	East:  {X: 1, Y: 0},
}

func (world World) Move(char *Character, direction Direction) bool {
	location := char.Location
	char.Direction = direction
	delta := moveTransform[direction]
	proposed := image.Point{X: location.X + delta.X, Y: location.Y + delta.Y}

	if world.Geography.At(proposed) == nil {
		return false
	}

	for being := range world.Beings {
		if char.Ident().Index == being.Ident().Index { // Ignore if we are the same being
			continue
		}
		if being.Location.X == proposed.X && being.Location.Y == proposed.Y && !being.Passable(char) {
			return false
		}
	}

	if world.Geography.CanPass(proposed, char) {
		world.Geography.RemoveLoc(*location, char)
		world.Geography.AddLoc(proposed, char)
		char.Location = &proposed

		return true
	}

	return false
}
