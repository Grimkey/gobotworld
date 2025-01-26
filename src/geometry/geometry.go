package geometry

import (
	"image"
	"math"
)

type Window struct {
	Left   int
	Top    int
	Width  int
	Height int
}

func (wnd Window) Within(point image.Point) bool {
	if point.X < wnd.Left || point.X > wnd.Width {
		return false
	}
	if point.Y < wnd.Top || point.Y > wnd.Height {
		return false
	}
	return true
}

func (wnd Window) Overlap(other Window) bool {
	widthOverlap := intMin(wnd.Left+wnd.Width, other.Left+other.Width) > intMax(wnd.Left, other.Left)
	heightOverlap := intMin(wnd.Top+wnd.Height, other.Top+other.Height) > intMax(wnd.Top, other.Top)

	return widthOverlap && heightOverlap
}

func intMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func intMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Circle(origin *image.Point, size int) Window {
	left := origin.X - size - 1
	if left < 0 {
		left = 0
	}
	top := origin.Y - size
	if top < 0 {
		top = 0
	}
	return Window{
		Left:   left,
		Width:  size*2 + 1,
		Top:    top,
		Height: size*2 + 1,
	}
}

func Distance(p1 image.Point, p2 image.Point) int {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return QuickSqrt(dx*dx + dy*dy)
}

func QuickSqrt(x int) int {
	xfloat := float64(x)
	if xfloat == 0 {
		return 0
	}
	q := xfloat / 2
	if q == 0 {
		return 0
	}
	s := (q + xfloat/q) / 2   /*first guess*/
	for i := 1; i <= 4; i++ { /*average of guesses*/
		if s == 0 {
			break
		}
		s = (s + xfloat/s) / 2
	}
	return int(math.Round(s))
}
