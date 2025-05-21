package stellarhex

import (
	"fmt"
	"testing"

	"github.com/Galdoba/t5/internal/grid/coordinates"
	"github.com/Galdoba/t5/pkg/dice"
)

func TestNew(t *testing.T) {
	// type args struct {
	// 	crd     coordinates.GlobalWorldCoordinates
	// 	options []StellarHexOption
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want *StellarHex
	// }{
	// 	{
	// 		name: "",
	// 		args: args{
	// 			crd:     coordinates.GlobalWorldCoordinates{1, 1},
	// 			options: []StellarHexOption{},
	// 		},
	// 		want: &StellarHex{},
	// 	}, // TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := New(tt.args.crd, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("New() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
	sh := New(coordinates.GlobalWorldCoordinates{1, 1})
	err := sh.GenerateMissingDetails(dice.NewDicepool())
	fmt.Println(err)
	fmt.Println(sh)
}
