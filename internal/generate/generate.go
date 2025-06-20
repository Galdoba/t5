package generate

import (
	"github.com/Galdoba/t5/internal/cosmology/star"
	"github.com/Galdoba/t5/internal/generate/stars"
)

func (gn *Generator) GenerateStar(options ...star.StarOption) (*star.Star, error) {
	starGen := stars.NewStarGenerator(gn.dp, gn.rules, options...)
	return starGen.Generate(options...)
}
