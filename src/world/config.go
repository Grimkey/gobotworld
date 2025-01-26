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
	terrainTypes   []DefaultTerrain
	terrainSum     int
	rndObjextIndex func(int) int
}

func NewConfig(t ...DefaultTerrain) Config {
	sum := 0
	internalTerrain := make([]DefaultTerrain, 0, len(t))
	for _, item := range t {
		sum += item.Units
		internalTerrain = append(internalTerrain, DefaultTerrain{item.Type, sum})
	}
	cfg := Config{
		terrainTypes:   internalTerrain,
		terrainSum:     sum,
		rndObjextIndex: randomObjectIndex, // Useful for deterministic testing
	}
	return cfg
}

func (c Config) getObjectType(n int) object.Thing {
	for _, v := range c.terrainTypes {
		if n < v.Units {
			return v.Type
		}
	}
	panic("randomization incorrect")
}

func (c Config) RandomObject() object.Thing {
	if c.terrainSum == 0 {
		panic("no terrain types")
	}
	rnd := c.rndObjextIndex(c.terrainSum)
	return c.getObjectType(rnd)
}

func randomObjectIndex(max int) int {
	return rand.Intn(max)
}
