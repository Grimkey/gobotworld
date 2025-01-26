package world

import (
	"gobotworld/src/world/object"
	"image"
)

type Direction int

const (
	North Direction = 1
	West  Direction = 2
	South Direction = 3
	East  Direction = 4
)

type Character struct {
	ident     object.Object
	Location  *image.Point
	index     int
	objType   object.ObjectType
	Direction Direction
}

func (ch *Character) Ident() object.Object {
	return ch.ident
}

func (ch *Character) Passable(_o object.Thing) bool {
	return false
}

func newCharacter(id int, t object.ObjectType, start image.Point) *Character {
	c := Character{
		ident:     object.Object{Index: id, Type: t},
		Location:  &start,
		Direction: North,
	}

	return &c
}

func NewPlayer(start image.Point) *Character {
	return newCharacter(99, object.PlayerType, start)
}

func NewNPC(start image.Point) *Character {
	return newCharacter(20, object.EnemyType, start)
}
