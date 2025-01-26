package object

type LightBlock struct {
	Time  DayCycle
	Lumen int
}

type DayCycle int

const (
	NightTime = DayCycle(-1)
	DayTime   = DayCycle(1)
)
