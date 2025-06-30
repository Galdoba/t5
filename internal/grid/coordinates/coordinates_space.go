package coordinates

import (
	"fmt"

	"github.com/Galdoba/t5/internal/grid/coordinates/convert"
	"github.com/Galdoba/t5/internal/grid/coordinates/cube"
	"github.com/Galdoba/t5/internal/grid/coordinates/global"
	"github.com/Galdoba/t5/internal/grid/coordinates/local"
	"github.com/Galdoba/t5/internal/grid/coordinates/sector"
)

type SpaceCoordinates struct {
	cube   cube.Cube
	global global.SpaceGlobal
	local  local.SpaceSectorLocal
}

// NewSpaceCoordinates - cretare space hex coordinates based on values provided:
//
// 2 values: GlobalWorldCoordinates (X, Y)
// 3 values: Cube Coordinates (Q,R,S)
// 4 values: Local Sector Coordinates (SectorX, SectorY, LocalX, LocalY)
func NewSpaceCoordinates(values ...int) SpaceCoordinates {
	sc := SpaceCoordinates{}
	switch len(values) {
	case 2:
		sc.global = global.NewSpaceGlobal(values[0], values[1])
		sc.cube = convert.GlobalToCube(sc.global)
		sc.local = convert.GlobalToLocal(sc.global)
	case 3:
		sc.cube = cube.NewCube(values[0], values[1], values[2])
		sc.global = convert.CubeToGlobal(sc.cube)
		sc.local = convert.CubeToLocal(sc.cube)
	case 4:
		sc.local = local.NewLocal(values[0], values[1], values[2], values[3])
		sc.cube = convert.LocalToCube(sc.local)
		sc.global = convert.LocalToGlobal(sc.local)
	default:
		panic(fmt.Sprintf("unsupported values quantity (%v)", len(values)))
	}

	return sc
}

func (sc SpaceCoordinates) StringSectorNameHex() string {
	return fmt.Sprintf("%v %v", sector.Name(sc.local.SectorX, sc.local.SectorY), sector.Hex(sc.local.X, sc.local.Y))
}

func (sc SpaceCoordinates) hexValues() (int, int, int) {
	return sc.cube.Q, sc.cube.R, sc.cube.S
}

// func (sc SpaceCoordinates) globalValues() (int, int) {
// 	return sc.global.X, sc.global.Y
// }

// func (sc SpaceCoordinates) localValues() (int, int, int, int) {
// 	return sc.local.SectorX, sc.local.SectorY, sc.local.X, sc.local.Y
// }

func (sc SpaceCoordinates) Validate() error {
	if err := convert.RoundTrip(sc.hexValues()); err != nil {
		return err
	}
	return nil
}
