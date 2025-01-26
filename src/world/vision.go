package world

import (
	"gobotworld/src/geometry"
	"gobotworld/src/world/object"
	"image"
)

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

	cycle, _ := world.Time()

	return object.LightBlock{Time: cycle, Lumen: lumen}
}
