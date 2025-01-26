package terminal

import (
	"gobotworld/src/world/object"
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
)

func TestTint(t *testing.T) {
	originalColor := tcell.NewRGBColor(100, 150, 200)
	tintedColor := Tint(originalColor, 0.5)

	r, g, b := tintedColor.RGB()
	assert.Equal(t, int32(50), r, "Red component should be 50")
	assert.Equal(t, int32(75), g, "Green component should be 75")
	assert.Equal(t, int32(100), b, "Blue component should be 100")
}

func TestTintStyleBackground(t *testing.T) {
	originalStyle := tcell.StyleDefault.Background(tcell.NewRGBColor(100, 100, 100))
	tintedStyle := TintStyleBackground(originalStyle, 0.5)

	_, bg, _ := tintedStyle.Decompose()
	r, g, b := bg.RGB()
	assert.Equal(t, int32(50), r, "Red component of background should be tinted to 50")
	assert.Equal(t, int32(50), g, "Green component of background should be tinted to 50")
	assert.Equal(t, int32(50), b, "Blue component of background should be tinted to 50")
}

func TestFindRuneStyle_DayTime(t *testing.T) {
	obj := object.NewObject(1, object.PlayerType, true)
	light := object.LightBlock{Time: object.DayTime, Lumen: 2}

	runeStyle := FindRuneStyle(obj, light)

	assert.Equal(t, 'M', runeStyle.Symbol, "Symbol for PlayerType should be 'M'")
	assert.Equal(t, fire3Style.Background(tcell.ColorBlack), runeStyle.Style.Background(tcell.ColorBlack), "Style should match fire3Style for Lumen > 1")
}

func TestFindRuneStyle_NightTime(t *testing.T) {
	obj := object.NewObject(1, object.PlayerType, true)
	light := object.LightBlock{Time: object.NightTime, Lumen: 0}

	runeStyle := FindRuneStyle(obj, light)

	assert.Equal(t, 'M', runeStyle.Symbol, "Symbol for PlayerType should be 'M'")
	assert.Equal(t, nightPlayerStyle.Background(tcell.ColorBlack), runeStyle.Style.Background(tcell.ColorBlack), "Style should match nightPlayerStyle for StyleTypePlayer at NightTime")
}

func TestFindRuneStyle_NightTime_Light(t *testing.T) {
	obj := object.NewObject(1, object.TorchType, true)
	light := object.LightBlock{Time: object.NightTime, Lumen: 3}

	runeStyle := FindRuneStyle(obj, light)

	assert.Equal(t, '^', runeStyle.Symbol, "Symbol for TorchType should be '^'")
	assert.Equal(t, fire4Style.Background(tcell.ColorBlack), runeStyle.Style.Background(tcell.ColorBlack), "Style should match fire4Style for Lumen > 3")
}

func TestDayRuneStyle(t *testing.T) {
	obj := object.NewObject(1, object.Dirt2Type, true)
	light := object.LightBlock{Time: object.DayTime, Lumen: 0}

	runeStyle := dayRuneStyle(obj, light)

	assert.Equal(t, '.', runeStyle.Symbol, "Symbol for Dirt2Type should be '.'")
	assert.Equal(t, dayDefaultStyle.Background(DayGreen), runeStyle.Style.Background(DayGreen), "Style should match dayDefaultStyle for StyleTypeDefault at DayTime")
}

func TestNightRuneStyle(t *testing.T) {
	obj := object.NewObject(1, object.ObstacleType, false)
	light := object.LightBlock{Time: object.NightTime, Lumen: 0}

	runeStyle := nightRuneStyle(obj, light)

	assert.Equal(t, '@', runeStyle.Symbol, "Symbol for ObstacleType should be '@'")
	assert.Equal(t, nightObstacleStyle.Background(tcell.ColorBlack), runeStyle.Style.Background(tcell.ColorBlack), "Style should match nightObstacleStyle for StyleTypeObstacle at NightTime")
}
