package stellar

import "github.com/Galdoba/t5/internal/grid/coordinates"

type StellarHex struct {
	Coordiantes     coordinates.GlobalWorldCoordinates
	Density         int
	StarSystem      string
	D               string
	BD              string
	LGG             string
	MGG             []string //2
	SGG             []string //3
	RogueWorldBig   []string //4
	RogueWorldSmall []string //10
	Planetoids      []string //20
	PSR             string
	NS              string
	BH              string
	NB              string
}
