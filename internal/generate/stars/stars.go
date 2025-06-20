package stars

import (
	"fmt"

	"github.com/Galdoba/t5/internal/cosmology/star"
	"github.com/Galdoba/t5/internal/generate/stars/basic"
	"github.com/Galdoba/t5/pkg/dice"
)

type StarGenerator struct {
	dp        *dice.Dicepool
	method    string
	generated *star.Star
}

func NewStarGenerator(dp *dice.Dicepool, rules map[string]string, options ...star.StarOption) *StarGenerator {
	sg := StarGenerator{
		dp:        dp,
		method:    rules["method"],
		generated: &star.Star{},
	}
	return &sg
}

func (gn *StarGenerator) Generate(options ...star.StarOption) (*star.Star, error) {
	switch gn.method {
	case "":
		return nil, fmt.Errorf("star generation method is not set")
	default:
		return nil, fmt.Errorf("star generation method %v is not supported", gn.method)
	case star.GenerationMethod_T5basic:
		st, err := basic.NewStar(gn.dp, options...)
		if err != nil {
			return nil, err
		}
		gn.generated = st
	}
	return gn.generated, nil
}
