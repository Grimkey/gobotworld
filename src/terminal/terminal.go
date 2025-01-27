package terminal

import (
	"fmt"
	"gobotworld/src/geometry"
	"gobotworld/src/world"
	"gobotworld/src/world/object"
	"image"
	"log"

	"github.com/gdamore/tcell/v2"
)

var DefaultDisplayLength = 15

type Terminal struct {
	CommandWidth int
	screen       tcell.Screen
	Logger       *log.Logger
}

func Init() (Terminal, error) {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	terminal := Terminal{screen: s, CommandWidth: DefaultDisplayLength}

	if e != nil {
		return terminal, e
	}
	if e = s.Init(); e != nil {
		return terminal, e
	}

	s.SetStyle(nightDefaultStyle)
	s.Clear()

	return terminal, nil
}

func (t Terminal) PollEvent() tcell.Event {
	return t.screen.PollEvent()
}

func (t Terminal) Show() {
	t.screen.Show()
}

func (t Terminal) Fini() {
	t.screen.Fini()
}

func (t Terminal) SetCell(x int, y int, runeStyle RuneStyle) {
	t.screen.SetCell(x, y, runeStyle.Style, runeStyle.Symbol)
}

func (t Terminal) PrintPath(path map[image.Point]bool) {
	t.Logger.Print("Path: ")
	for point := range path {
		t.Logger.Printf("(%d, %d)\n", point.X, point.Y)
	}
}

func (t Terminal) DrawWorld(gameWorld world.World) {
	playerLocation := *gameWorld.Player.Location
	wnd := t.drawWindow(playerLocation, gameWorld.Geography.Height(), gameWorld.Geography.Width())
	cycle, count := object.Time(*gameWorld.Time)
	pathFinder := world.PathFinder{World: gameWorld, Logger: t.Logger}

	nearestLight := gameWorld.Lights.NearestLight(playerLocation)
	if nearestLight.X == -1 && nearestLight.Y == -1 {
		if t.Logger != nil {
			t.Logger.Println("No nearest light found.")
		}
		return
	}
	t.Logger.Printf("Nearest light found at %d, %d\n", nearestLight.X, nearestLight.Y)
	path := pathFinder.Find(playerLocation, nearestLight)
	t.PrintPath(path)

	x := 0
	y := 1
	w, _ := t.screen.Size()
	display := w - t.CommandWidth
	t.SetCell(display, 0, RuneStyle{Symbol: '|', Style: borderStyle})

	for col := wnd.Top; col < wnd.Height; col++ {
		t.SetCell(w-t.CommandWidth, y, RuneStyle{Symbol: '|', Style: borderStyle})

		for row := wnd.Left; row < wnd.Width; row++ {
			pt := gameWorld.Geography[col][row]
			loc := image.Point{X: row, Y: col}

			light := LightValue(loc, wnd, gameWorld.Lights, cycle)
			sense := SenseValue(loc, *gameWorld.Player.Location, gameWorld.Player.Direction)
			runeStyle := drawCell(pt, light, sense)

			if path != nil {
				if ok := path[loc]; ok {
					runeStyle = RuneStyle{Symbol: runeStyle.Symbol, Style: pathStyle}
				}
			}

			t.SetCell(x, y, runeStyle)
			x += 1
		}

		x = 0
		y += 1
	}

	player := gameWorld.Player
	str := fmt.Sprintf("X: %d, Y: %d -- Left: %d, Top: %d, Width: %d, Height: %d", player.Location.X, player.Location.Y, wnd.Left, wnd.Top, wnd.Width, wnd.Height)
	t.screen.SetContent(0, 0, ' ', []rune(str), displayStyle)

	str = "::Time::"
	t.screen.SetContent(w-t.CommandWidth+1, 2, ' ', []rune(str), borderStyle)

	c := "Night"
	if cycle == object.DayTime {
		c = "  Day"
	}
	str = fmt.Sprintf("%s %d", c, count)
	t.screen.SetContent(w-t.CommandWidth+1, 3, ' ', []rune(str), borderStyle)
}

func drawCell(l object.ThingList, light object.LightBlock, bgFactor float32) RuneStyle {
	runeStyle := RuneStyle{Symbol: 'X', Style: borderStyle}

	idx := -1
	for _, obj := range l {
		if obj.Ident().Index > idx {
			idx = obj.Ident().Index
			runeStyle = FindRuneStyle(obj, light)
		}
	}

	runeStyle.Style = TintStyleBackground(runeStyle.Style, bgFactor)
	return runeStyle
}

func (t Terminal) drawWindow(player image.Point, worldHeight, worldWidth int) geometry.Window {
	w, h := t.screen.Size()
	w -= t.CommandWidth
	h -= 1
	if w == 0 || h == 0 {
		panic("invalid screen size")
	}

	left := player.X - (w / 2)
	width := player.X + (w / 2)
	if left < 0 {
		left = 0
		width = w
	}
	if width >= worldWidth {
		left = worldWidth - w
		width = worldWidth
	}
	top := player.Y - (h / 2)
	height := player.Y + (h / 2)
	if top < 0 {
		top = 1
		height = h
	}
	if height >= worldHeight {
		top = worldHeight - h
		height = worldHeight
	}

	return geometry.Window{Left: left, Top: top, Width: width, Height: height}
}
