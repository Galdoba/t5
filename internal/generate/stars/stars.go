package stars

import (
	"github.com/Galdoba/t5/internal/cosmology/star"
	"github.com/Galdoba/t5/internal/generate/stars/basic"
	"github.com/Galdoba/t5/pkg/dice"
)

const (
	GenerationMethod_T5basic      = "basic"
	GenerationMethod_Worldbuilder = "worldbuilder"
)

type StarGenerator struct {
	seed      int64
	dp        *dice.Dicepool
	method    string
	generated *star.Star
}

func (gn *StarGenerator) Generate(options ...star.StarOption) error {
	switch gn.method {
	case GenerationMethod_T5basic:
		st, err := basic.NewStar(gn.dp, options...)
		if err != nil {
			return err
		}
		gn.generated = st
	}
	return nil
}
