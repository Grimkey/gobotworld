package geometry_test

import (
	"gobotworld/src/geometry"
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWindowWithin(t *testing.T) {
	wnd := geometry.Window{Left: 0, Top: 0, Width: 10, Height: 10}

	assert.True(t, wnd.Within(image.Point{X: 5, Y: 5}), "Point should be within the window")
	assert.False(t, wnd.Within(image.Point{X: -1, Y: 5}), "Point with negative X should not be within the window")
	assert.False(t, wnd.Within(image.Point{X: 5, Y: -1}), "Point with negative Y should not be within the window")
	assert.False(t, wnd.Within(image.Point{X: 11, Y: 5}), "Point beyond the width should not be within the window")
	assert.False(t, wnd.Within(image.Point{X: 5, Y: 11}), "Point beyond the height should not be within the window")
}

func TestWindowOverlap(t *testing.T) {
	wnd1 := geometry.Window{Left: 0, Top: 0, Width: 10, Height: 10}
	wnd2 := geometry.Window{Left: 5, Top: 5, Width: 10, Height: 10}
	wnd3 := geometry.Window{Left: 20, Top: 20, Width: 10, Height: 10}

	assert.True(t, wnd1.Overlap(wnd2), "Windows with overlapping regions should return true")
	assert.False(t, wnd1.Overlap(wnd3), "Non-overlapping windows should return false")
}

func TestCircle(t *testing.T) {
	origin := image.Point{X: 5, Y: 5}
	size := 3

	circle := geometry.Circle(&origin, size)

	expected := geometry.Window{
		Left:   1,
		Top:    2,
		Width:  7,
		Height: 7,
	}
	assert.Equal(t, expected, circle, "Circle function should return the correct window")
}

func TestDistance(t *testing.T) {
	p1 := image.Point{X: 0, Y: 0}
	p2 := image.Point{X: 3, Y: 4}

	distance := geometry.Distance(p1, p2)
	assert.Equal(t, 5, distance, "Distance between (0,0) and (3,4) should be 5")
}

func TestQuickSqrt(t *testing.T) {
	assert.Equal(t, 5, geometry.QuickSqrt(25), "Square root of 25 should be 5")
	assert.Equal(t, 10, geometry.QuickSqrt(100), "Square root of 100 should be 10")
	assert.Equal(t, 0, geometry.QuickSqrt(0), "Square root of 0 should be 0")
	assert.Equal(t, 1, geometry.QuickSqrt(1), "Square root of 1 should be 1")
	assert.Equal(t, 2, geometry.QuickSqrt(4), "Square root of 4 should be 2")
}
