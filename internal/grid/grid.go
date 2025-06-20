package grid

import (
	"github.com/Galdoba/t5/internal/grid/coordinates"
	"github.com/Galdoba/t5/internal/grid/stellarhex"
)

type Grid struct {
	Cell map[coordinates.GlobalWorldCoordinates]*stellarhex.StellarHex
}
