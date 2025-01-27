package object

import (
	"image"
)

type Direction int

const (
	North Direction = 1
	West  Direction = 2
	South Direction = 3
	East  Direction = 4
)

// String method for Direction
func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case West:
		return "West"
	case South:
		return "South"
	case East:
		return "East"
	default:
		return "Unknown"
	}
}

type Character struct {
	ident     Object
	Location  *image.Point
	index     int
	objType   ObjectType
	Direction Direction
}

func (ch *Character) Ident() Object {
	return ch.ident
}

func (ch *Character) Passable(_o Thing) bool {
	return false
}

func newCharacter(id int, t ObjectType, start image.Point) *Character {
	c := Character{
		ident:     Object{Index: id, Type: t},
		Location:  &start,
		Direction: North,
	}

	return &c
}

func NewPlayer(start image.Point) *Character {
	return newCharacter(99, PlayerType, start)
}

func NewNPC(start image.Point) *Character {
	return newCharacter(20, EnemyType, start)
}
