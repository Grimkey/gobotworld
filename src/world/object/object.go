package object

import (
	"gobotworld/src/geometry"
	"image"
	"slices"
)

type ObjectType int

const TorchArea = 4

const (
	Dirt1Type    = ObjectType(0)
	Dirt2Type    = ObjectType(1)
	RockType     = ObjectType(2)
	ObstacleType = ObjectType(3)
	PlayerType   = ObjectType(4)
	EnemyType    = ObjectType(5)
	TorchType    = ObjectType(6)
)

type Thing interface {
	Ident() Object
	Passable(o Thing) bool
}

type Object struct {
	Index int
	Type  ObjectType
}

type BasicObject struct {
	ident    Object
	passable bool
}

var NullThing Thing = NewObject(-1, Dirt1Type, false)

func NewObject(idx int, objType ObjectType, passable bool) BasicObject {
	return BasicObject{ident: Object{idx, objType}, passable: passable}
}

func (bo BasicObject) Ident() Object {
	return bo.ident
}

func (bo BasicObject) Passable(_o Thing) bool {
	return bo.passable
}

type ThingList []Thing

func (tl ThingList) DeleteItem(t Thing) ThingList {
	return slices.DeleteFunc(tl, func(oth Thing) bool {
		return oth.Ident() == t.Ident()
	})
}

func (tl ThingList) Top() Thing {
	if len(tl) == 0 {
		return NullThing
	}

	return tl[0]
}

type Light struct {
	ident Object
	Area  int
}

func NewLight(area int) Light {
	return Light{ident: Object{10, TorchType}, Area: area}
}

func (lt Light) Ident() Object {
	return lt.ident
}

func (lt Light) Passable(_o Thing) bool {
	return true
}

func LightAt(origin image.Point, target image.Point, area int) int {
	distance := geometry.Distance(origin, target)
	lumen := 0
	if distance <= area {
		lumen = distance
	}
	return lumen
}
