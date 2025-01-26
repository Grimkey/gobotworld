package object_test

import (
	"gobotworld/src/world/object"
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewObject(t *testing.T) {
	obj := object.NewObject(1, object.PlayerType, true)

	assert.Equal(t, 1, obj.Ident().Index, "Object index should be 1")
	assert.Equal(t, object.PlayerType, obj.Ident().Type, "Object type should be PlayerType")
	assert.True(t, obj.Passable(nil), "Passable should be true for this object")
}

func TestNullThing(t *testing.T) {
	assert.Equal(t, -1, object.NullThing.Ident().Index, "NullThing should have an index of -1")
	assert.False(t, object.NullThing.Passable(nil), "NullThing should not be passable")
}

func TestThingListDeleteItem(t *testing.T) {
	obj1 := object.NewObject(1, object.PlayerType, true)
	obj2 := object.NewObject(2, object.EnemyType, false)

	list := object.ThingList{obj1, obj2}
	newList := list.DeleteItem(obj1)

	assert.Len(t, newList, 1, "List should have 1 item after deletion")
	assert.Equal(t, obj2.Ident(), newList[0].Ident(), "Remaining item should be obj2")
}

func TestThingListTop(t *testing.T) {
	obj1 := object.NewObject(1, object.PlayerType, true)
	obj2 := object.NewObject(2, object.EnemyType, false)

	list := object.ThingList{obj1, obj2}
	top := list.Top()

	assert.Equal(t, obj1.Ident(), top.Ident(), "Top should return the first item in the list")

	emptyList := object.ThingList{}
	assert.Equal(t, object.NullThing.Ident(), emptyList.Top().Ident(), "Top should return NullThing for an empty list")
}

func TestNewLight(t *testing.T) {
	light := object.NewLight(10)

	assert.Equal(t, object.TorchType, light.Ident().Type, "Light type should be TorchType")
	assert.True(t, light.Passable(nil), "Light should always be passable")
	assert.Equal(t, 10, light.Area, "Light area should be 10")
}

func TestLightAt(t *testing.T) {
	origin := image.Point{X: 5, Y: 5}
	target := image.Point{X: 6, Y: 6}
	area := 3

	lumen := object.LightAt(origin, target, area)
	assert.Equal(t, 1, lumen, "Lumen should equal distance when within area")

	targetOutside := image.Point{X: 10, Y: 10}
	lumen = object.LightAt(origin, targetOutside, area)
	assert.Equal(t, 0, lumen, "Lumen should be 0 when outside the light area")
}

func TestBasicObjectPassable(t *testing.T) {
	obj := object.NewObject(1, object.RockType, false)
	assert.False(t, obj.Passable(nil), "Passable should be false for this object")

	obj = object.NewObject(2, object.PlayerType, true)
	assert.True(t, obj.Passable(nil), "Passable should be true for this object")
}
