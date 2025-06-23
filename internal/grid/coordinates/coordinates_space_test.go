package coordinates

import (
	"reflect"
	"testing"

	"github.com/Galdoba/t5/internal/grid/coordinates/cube"
	"github.com/Galdoba/t5/internal/grid/coordinates/global"
	"github.com/Galdoba/t5/internal/grid/coordinates/local"
)

func TestNewSpaceCoordinates(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want SpaceCoordinates
	}{
		{name: "start global", args: args{values: []int{0, 0}}, want: SpaceCoordinates{hex: cube.Hex{0, 0, 0}, global: global.SpaceGlobal{0, 0}, local: local.SpaceSectorLocal{0, 0, 1, 40}}},
		{name: "start cube", args: args{values: []int{0, 0, 0}}, want: SpaceCoordinates{hex: cube.Hex{0, 0, 0}, global: global.SpaceGlobal{0, 0}, local: local.SpaceSectorLocal{0, 0, 1, 40}}},
		{name: "start local", args: args{values: []int{0, 0, 1, 40}}, want: SpaceCoordinates{hex: cube.Hex{0, 0, 0}, global: global.SpaceGlobal{0, 0}, local: local.SpaceSectorLocal{0, 0, 1, 40}}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSpaceCoordinates(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpaceCoordinates(%v) = %v, want %v", tt.args.values, got, tt.want)
			}

		})
	}
}
