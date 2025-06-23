package coordinates

import (
	"fmt"

	"github.com/Galdoba/t5/internal/grid/coordinates/cube"
	"github.com/Galdoba/t5/internal/grid/coordinates/global"
	"github.com/Galdoba/t5/internal/grid/coordinates/local"
	"github.com/Galdoba/t5/internal/grid/coordinates/sector"
)

type SpaceCoordinates struct {
	hex    cube.Hex
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
		sc.hex = cube.NewCube(global_to_cube(sc.globalValues()))
		sc.local = local.NewLocal(global_to_local(sc.globalValues()))
	case 3:
		sc.hex = cube.NewCube(values[0], values[1], values[2])
		sc.global = global.NewSpaceGlobal(cube_to_global(sc.hexValues()))
		sc.local = local.NewLocal(cube_to_local(sc.hexValues()))
	case 4:
		sc.local = local.NewLocal(values[0], values[1], values[2], values[3])
		sc.hex = cube.NewCube(local_to_cube(sc.localValues()))
		sc.global = global.NewSpaceGlobal(local_to_global(sc.localValues()))
	default:
		panic(fmt.Sprintf("unsupported values quantity (%v)", len(values)))
	}

	return sc
}

func (sc SpaceCoordinates) StringSectorNameHex() string {
	return fmt.Sprintf("%v %v", sector.Name(sc.local.SectorX, sc.local.SectorY), sector.Hex(sc.local.X, sc.local.Y))
}

func (sc SpaceCoordinates) hexValues() (int, int, int) {
	return sc.hex.Q, sc.hex.R, sc.hex.S
}

func (sc SpaceCoordinates) globalValues() (int, int) {
	return sc.global.X, sc.global.Y
}

func (sc SpaceCoordinates) localValues() (int, int, int, int) {
	return sc.local.SectorX, sc.local.SectorY, sc.local.X, sc.local.Y
}

func (sc SpaceCoordinates) Validate() error {
	if err := roundTrip(sc.hexValues()); err != nil {
		return err
	}
	return nil
}
