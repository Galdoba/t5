package stars

import (
	"fmt"
	"testing"

	"github.com/Galdoba/t5/internal/cosmology/star"
	"github.com/Galdoba/t5/pkg/dice"
)

func TestStarGenerator_Generate(t *testing.T) {
	var seed int64 = 4253622
	gn := StarGenerator{
		seed:      seed,
		dp:        dice.NewDicepool(dice.WithSeed(seed)),
		method:    GenerationMethod_T5basic,
		generated: &star.Star{},
	}
	fmt.Println(gn.Generate(star.WithPosition(star.Near)))
	fmt.Println(gn.generated)
}
