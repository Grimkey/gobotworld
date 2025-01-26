package terminal

import (
	"github.com/gdamore/tcell/v2"
	"gobotworld/src/world/object"
)

var (
	DayGreen = tcell.NewRGBColor(0x66, 0x66, 0x00)
	DayGray  = tcell.NewRGBColor(0x00, 0x99, 0x99)
	fire1    = tcell.NewRGBColor(0xFA, 0xC0, 0x00)
	fire2    = tcell.NewRGBColor(0xFF, 0x75, 0x00)
	fire3    = tcell.NewRGBColor(0xFC, 0x64, 0x00)
	fire4    = tcell.NewRGBColor(0xD7, 0x35, 0x02)
	fire5    = tcell.NewRGBColor(0xB6, 0x22, 0x03)
	fire6    = tcell.NewRGBColor(0x80, 0x11, 0x00)
)

func Tint(c tcell.Color, factor float32) tcell.Color {
	r, g, b := c.RGB()
	return tcell.NewRGBColor(int32(float32(r)*factor), int32(float32(g)*factor), int32(float32(b)*factor))
}

func TintStyleBackground(c tcell.Style, factor float32) tcell.Style {
	fg, bg, _ := c.Decompose()
	return tcell.StyleDefault.Foreground(fg).Background(Tint(bg, factor))
}

var (
	displayStyle = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
	borderStyle  = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	dayDefaultStyle  = tcell.StyleDefault.Foreground(DayGray).Background(DayGreen)
	dayObstacleStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(DayGreen)

	nightObstacleStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	nightDefaultStyle  = tcell.StyleDefault.Foreground(tcell.ColorGray).Background(tcell.ColorBlack)

	dayPlayerStyle   = tcell.StyleDefault.Foreground(tcell.ColorDarkRed).Background(DayGreen)
	nightPlayerStyle = tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)

	fire1Style = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(fire1)
	fire2Style = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(fire2)
	fire3Style = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(fire3)
	fire4Style = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(fire4)
	fire5Style = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(fire5)
	fire6Style = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(fire6)

	pathStyle = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
)

type StyleType int

var (
	StyleTypeDefault  = StyleType(0)
	StyleTypeObstacle = StyleType(1)
	StyleTypePlayer   = StyleType(2)
	StyleTypeLight    = StyleType(3)
)

type display struct {
	Symbol    rune
	StyleType StyleType
}

var terrainSymbols = map[object.ObjectType]display{
	object.Dirt1Type:    {' ', StyleTypeDefault},
	object.Dirt2Type:    {'.', StyleTypeDefault},
	object.RockType:     {'o', StyleTypeDefault},
	object.ObstacleType: {'@', StyleTypeObstacle},
	object.PlayerType:   {'M', StyleTypePlayer},
	object.EnemyType:    {'E', StyleTypePlayer},
	object.TorchType:    {'^', StyleTypeLight},
}

type RuneStyle struct {
	Symbol rune
	Style  tcell.Style
}

func FindRuneStyle(obj object.Thing, light object.LightBlock) RuneStyle {
	if light.Time == object.DayTime {
		return dayRuneStyle(obj, light)
	}

	return nightRuneStyle(obj, light)
}

func nightRuneStyle(obj object.Thing, light object.LightBlock) RuneStyle {
	symbol := terrainSymbols[obj.Ident().Type]
	style := borderStyle
	switch {
	case light.Lumen > 4:
		style = fire5Style
	case light.Lumen > 3:
		style = fire4Style
	case light.Lumen > 2:
		style = fire3Style
	case light.Lumen > 1:
		style = fire2Style
	case light.Lumen > 0:
		style = fire1Style
	case symbol.StyleType == StyleTypeLight:
		style = fire1Style
	case symbol.StyleType == StyleTypeDefault:
		style = nightDefaultStyle
	case symbol.StyleType == StyleTypePlayer:
		style = nightPlayerStyle
	case symbol.StyleType == StyleTypeObstacle:
		style = nightObstacleStyle
	}

	return RuneStyle{Symbol: symbol.Symbol, Style: style}
}

func dayRuneStyle(obj object.Thing, light object.LightBlock) RuneStyle {
	symbol := terrainSymbols[obj.Ident().Type]
	style := borderStyle
	switch {
	case light.Lumen > 2:
		style = fire4Style
	case light.Lumen > 1:
		style = fire3Style
	case light.Lumen > 0:
		style = fire2Style
	case symbol.StyleType == StyleTypeLight:
		style = fire1Style
	case symbol.StyleType == StyleTypeDefault:
		style = dayDefaultStyle
	case symbol.StyleType == StyleTypePlayer:
		style = dayPlayerStyle
	case symbol.StyleType == StyleTypeObstacle:
		style = dayObstacleStyle
	}

	return RuneStyle{Symbol: symbol.Symbol, Style: style}
}
