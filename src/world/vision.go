// Package provides a library for calculating vision and light in a 2D grid-based world.
package world

import (
	"gobotworld/src/geometry"
	"gobotworld/src/world/object"
	"image"
)

// Vision calculates the light intensity at a specific point based on
// surrounding light sources and the current day/night cycle.
//
// Parameters:
//   - pt: The target point for light calculation.
//   - viewable: The visible region as a geometry.Window.
//   - world: The world instance containing light sources and time information.
//
// Returns:
//   - An object.LightBlock containing the calculated light intensity (Lumen)
//     and the time of day.
func Vision(pt image.Point, viewable geometry.Window, world World) object.LightBlock {
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

	cycle, _ := object.Time(*world.Time)

	return object.LightBlock{Time: cycle, Lumen: lumen}
}
