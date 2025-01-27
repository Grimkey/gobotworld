package object

import (
	"gobotworld/src/geometry"
	"image"
)

const (
	ticksPerCount   = 4
	countPerDay     = 20
	countPerHalfDay = countPerDay / 2
)

type LightBlock struct {
	Time  DayCycle
	Lumen int
}

type DayCycle int

const (
	NightTime = DayCycle(-1)
	DayTime   = DayCycle(1)
)

func Time(time int) (DayCycle, int) {
	now := (time / ticksPerCount) % countPerDay
	cycle := DayTime
	if now > countPerHalfDay {
		cycle = NightTime
	}
	return cycle, now % countPerHalfDay
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

type Lights []*image.Point

func (lts Lights) NearestLight(p image.Point) image.Point {
	if len(lts) == 0 {
		return image.Point{X: -1, Y: -1} // Indicate no lights are available
	}

	var nearest image.Point
	var nearestDist = 0 // math.MaxInt
	found := false
	for _, light := range lts {
		d := geometry.Distance(*light, p)

		if d < nearestDist {
			nearest = *light
			nearestDist = d
			found = true
		}
	}

	if !found {
		return image.Point{X: 0, Y: 0} // No light found
	}
	return nearest
}
