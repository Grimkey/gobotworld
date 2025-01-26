package terminal

import (
	"gobotworld/src/geometry"
	"gobotworld/src/world"
	"gobotworld/src/world/object"
	"image"
	"math"
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
	currentAngle := math.Atan2(float64(char.Y-pt.Y), float64(pt.X-(char.X-2)))
	return currentAngle > -math.Pi/8 && currentAngle < math.Pi/8
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
