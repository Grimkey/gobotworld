package terminal

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"gobotworld/src/world"
	"gobotworld/src/world/character"
	"gobotworld/src/world/terrain"
)

var displayStyle = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
var defaultStyle = tcell.StyleDefault.Foreground(tcell.ColorGray).Background(tcell.ColorBlack)
var obstacleStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
var playerStyle = tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)

type terrainDisplay struct {
	Symbol rune
	Style  tcell.Style
}

var terrainSymbols = map[terrain.Terrain]terrainDisplay{
	terrain.Empty:    {' ', defaultStyle},
	terrain.Dirt:     {'.', defaultStyle},
	terrain.Rock:     {'o', defaultStyle},
	terrain.Obstacle: {'@', obstacleStyle},
}

type Window struct {
	Left   int
	Top    int
	Width  int
	Height int
}

var DefaultDisplayLength = 15

type Terminal struct {
	DisplayWidth  int
	DisplayHeight int
	screen        tcell.Screen
}

func Init() (Terminal, error) {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	w, h := s.Size()
	terminal := Terminal{screen: s, DisplayWidth: w - 15, DisplayHeight: h - 1}

	if e != nil {
		return terminal, e
	}
	if e = s.Init(); e != nil {
		return terminal, e
	}

	s.SetStyle(defaultStyle)
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

func (t Terminal) SetCell(x int, y int, style tcell.Style, ch ...rune) {
	t.screen.SetCell(x, y, style, ch...)
}

func (t Terminal) DrawWorld(world world.World) {
	playerLocation := *world.Beings[character.Player]
	wnd := t.drawWindow(playerLocation, world.Height(), world.Width())

	x := 0
	y := 1
	for col := wnd.Top; col < wnd.Height; col++ {
		for row := wnd.Left; row < wnd.Width; row++ {
			cell := world.Geography[col][row]
			if cell == nil {
				print("here")
			}
			terrain := terrainSymbols[cell.Terrain]
			t.screen.SetCell(x, y, terrain.Style, terrain.Symbol)
			x += 1
		}
		x = 0
		y += 1
	}

	for _, location := range world.Beings {
		t.screen.SetCell(location.X-wnd.Left, location.Y-wnd.Top+1, playerStyle, 'M')
	}

	player := world.Beings[character.Player]
	str := fmt.Sprintf("X: %d, Y: %d -- Left: %d, Top: %d, Width: %d, Height: %d", player.X, player.Y, wnd.Left, wnd.Top, wnd.Width, wnd.Height)

	t.screen.SetContent(0, 0, ' ', []rune(str), displayStyle)
}

func (t Terminal) drawWindow(player world.Point, worldHeight, worldWidth int) Window {
	//w := t.DisplayWidth
	//h := t.DisplayHeight
	w, h := t.screen.Size()
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

	return Window{left, top, width, height}
}
