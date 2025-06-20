package generate

import (
	"time"

	"github.com/Galdoba/t5/pkg/dice"
)

type Generator struct {
	seed             int64
	dp               *dice.Dicepool
	targetStructType string
	rules            map[string]string
}

func NewGenerator(genOpts ...GeneratorOption) *Generator {
	gn := Generator{}
	gn.seed = time.Now().UnixMilli()
	gn.rules = make(map[string]string)
	for _, set := range genOpts {
		set(&gn)
	}
	gn.dp = dice.NewDicepool(dice.WithSeed(gn.seed))
	return &gn
}

type GeneratorOption func(*Generator)

func WithSeed(seed int64) GeneratorOption {
	return func(g *Generator) {
		g.seed = seed
	}
}

func WithRule(ruleName, ruleValue string) GeneratorOption {
	return func(g *Generator) {
		g.rules[ruleName] = ruleValue
	}
}
