package terrain

type Type int

type Terrain struct {
	Index    int
	Passable bool
}

var Empty = Terrain{Index: 0, Passable: true}
var Dirt = Terrain{Index: 1, Passable: true}
var Rock = Terrain{Index: 2, Passable: true}
var Obstacle = Terrain{Index: 3, Passable: false}
