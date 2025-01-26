package terminal

import (
	"gobotworld/src/geometry"
	"gobotworld/src/world"
	"gobotworld/src/world/object"
	"image"
	"math"
)

var (
	sqrt3over2 = math.Sqrt(3) / 2
	sqrt2over2 = math.Sqrt(2) / 2

	arcTan30 = math.Atan2(sqrt3over2, 0.5)
	arcTan45 = math.Atan2(sqrt2over2, sqrt2over2)
	arcTan60 = math.Atan2(0.5, sqrt3over2)
	arcTan90 = math.Atan2(0, 1)

	arcTan120 = math.Atan2(-0.5, sqrt3over2)
	arcTan135 = math.Atan2(-sqrt2over2, sqrt2over2)
	arcTan150 = math.Atan2(-sqrt3over2, 0.5)
	arcTan180 = math.Atan2(-1, 0)

	arcTan210 = math.Atan2(-sqrt3over2, -0.5)
	arcTan225 = math.Atan2(-sqrt2over2, -sqrt2over2)
	arcTan240 = math.Atan2(-0.5, -sqrt3over2)
	arcTan270 = math.Atan2(0, -1)

	arcTan300 = math.Atan2(1/2, -sqrt3over2)
	arcTan315 = math.Atan2(sqrt2over2, -sqrt2over2)
	arcTan330 = math.Atan2(sqrt3over2, -0.5)
	arcTan360 = math.Atan2(1, 0)
)

func SenseValue(pt image.Point, wld world.World) float32 {
	dist := geometry.Distance(*wld.Player.Location, pt)
	if dist > 15 {
		return .75
	}

	arc := NorthVision
	switch wld.Player.Direction {
	case world.South:
		arc = SouthVision
	case world.East:
		arc = EastVision
	case world.West:
		arc = WestVision
	}

	if arc(pt, *wld.Player.Location) {
		return 1
	}
	return .75
}

func NorthVision(pt image.Point, char image.Point) bool {
	currentAngle := math.Atan2(float64(char.Y+2-pt.Y), float64(char.X-pt.X))
	return currentAngle < 7*math.Pi/8 && currentAngle > math.Pi/8
}

func SouthVision(pt image.Point, char image.Point) bool {
	currentAngle := math.Atan2(float64(char.Y-2-pt.Y), float64(char.X-pt.X))
	return currentAngle < -math.Pi/8 && currentAngle > -7*math.Pi/8
}

func WestVision(pt image.Point, char image.Point) bool {
	currentAngle := math.Atan2(float64(char.Y-pt.Y), float64(char.X+2-pt.X))
	return currentAngle > -1*math.Pi/8 && currentAngle < 1*math.Pi/8
}

func EastVision(pt image.Point, char image.Point) bool {
	currentAngle := math.Atan2(float64(char.Y-pt.Y), float64(char.X-2-pt.X))
	return currentAngle < math.Pi/8 && currentAngle > 15*math.Pi/8
}

func LightValue(pt image.Point, viewable geometry.Window, world world.World) object.LightBlock {
	lumen := 0
	for _, light := range world.Lights {
		lightWnd := geometry.Circle(light, object.TorchArea)
		if viewable.Overlap(lightWnd) {
			result := object.LightAt(pt, *light, object.TorchArea)
			if result > lumen {
				lumen = result
			}
		}
	}

	cycle, _ := world.Time()

	return object.LightBlock{Time: cycle, Lumen: lumen}
}
