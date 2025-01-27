// Package terminal provides utilities for sensory perception and lighting effects
// in a 2D game environment. This includes determining visibility of objects
// and calculating light intensity at specific points on the map.
package terminal

import (
	"gobotworld/src/geometry"
	"gobotworld/src/world/object"
	"image"
	"math"
)

// SenseValue calculates the visibility of a point relative to a character's
// location and direction. Visibility is influenced by distance and whether
// the point is within the character's field of view.
//
// Parameters:
//   - pt: The target point to sense.
//   - loc: The location of the character.
//   - direction: The character's facing direction.
//
// Returns:
//   - A float32 value representing the visibility of the point:
//     1.0 indicates full visibility, and 0.75 indicates reduced visibility
//     (e.g., if the point is outside the field of view or too far away).
func SenseValue(pt image.Point, loc image.Point, direction object.Direction) float32 {
	dist := geometry.Distance(loc, pt)
	if dist > 15 {
		return .75
	}

	arc := northVision
	switch direction {
	case object.South:
		arc = southVision
	case object.East:
		arc = eastVision
	case object.West:
		arc = westVision
	}

	if arc(pt, loc) {
		return 1
	}
	return .75
}

func northVision(pt image.Point, char image.Point) bool {
	currentAngle := math.Atan2(float64(char.Y+2-pt.Y), float64(char.X-pt.X))
	return currentAngle < 7*math.Pi/8 && currentAngle > math.Pi/8
}

func southVision(pt image.Point, char image.Point) bool {
	currentAngle := math.Atan2(float64(char.Y-2-pt.Y), float64(char.X-pt.X))
	return currentAngle < -math.Pi/8 && currentAngle > -7*math.Pi/8
}

func westVision(pt image.Point, char image.Point) bool {
	currentAngle := math.Atan2(float64(char.Y-pt.Y), float64(char.X+2-pt.X))
	return currentAngle > -1*math.Pi/8 && currentAngle < 1*math.Pi/8
}

func eastVision(pt image.Point, char image.Point) bool {
	currentAngle := math.Atan2(float64(char.Y-pt.Y), float64(pt.X-(char.X-2)))
	return currentAngle > -math.Pi/8 && currentAngle < math.Pi/8
}

// LightValue calculates the light intensity at a specific point based on
// surrounding light sources and the current day/night cycle.
//
// Parameters:
//   - pt: The target point for light calculation.
//   - viewable: The visible region as a geometry.Window.
//   - lights: A collection of light sources.
//   - cycle: The current day/night cycle.
//
// Returns:
//   - An object.LightBlock containing the calculated light intensity (Lumen)
//     and the time of day.
func LightValue(pt image.Point, viewable geometry.Window, lights object.Lights, cycle object.DayCycle) object.LightBlock {
	lumen := 0
	for _, light := range lights {
		lightWnd := geometry.Circle(light, object.TorchArea)
		if viewable.Overlap(lightWnd) {
			result := object.LightAt(pt, *light, object.TorchArea)
			if result > lumen {
				lumen = result
			}
		}
	}

	return object.LightBlock{Time: cycle, Lumen: lumen}
}
