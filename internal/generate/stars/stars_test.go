package stars

import (
	"fmt"
	"testing"

	"github.com/Galdoba/t5/internal/cosmology/star"
	"github.com/Galdoba/t5/pkg/dice"
)

func TestStarGenerator_Generate(t *testing.T) {
	gn := StarGenerator{
		dp:        dice.NewDicepool(dice.WithSeed(4253622)),
		method:    star.GenerationMethod_T5basic,
		generated: &star.Star{},
	}
	st, err := gn.Generate(star.WithPosition(star.Near))
	fmt.Println(err)
	fmt.Println(st)
}
