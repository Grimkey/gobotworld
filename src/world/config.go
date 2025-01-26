package world

import (
	"gobotworld/src/world/object"
	"math/rand"
)

type DefaultTerrain struct {
	Type  object.Thing
	Units int
}

type Config struct {
	terrainTypes []DefaultTerrain
	terrainSum   int
}

func NewConfig(t ...DefaultTerrain) Config {
	sum := 0
	internalTerrain := make([]DefaultTerrain, 0, len(t))
	for _, item := range t {
		sum += item.Units
		internalTerrain = append(internalTerrain, DefaultTerrain{item.Type, sum})
	}
	return Config{
		terrainTypes: internalTerrain,
		terrainSum:   sum,
	}
}

func (c Config) RandomObject() object.Thing {
	rnd := rand.Intn(c.terrainSum)
	for _, v := range c.terrainTypes {
		if rnd < v.Units {
			return v.Type
		}
	}
	panic("randomization incorrect")
}
