package basic

import (
	"github.com/Galdoba/t5/internal/cosmology/star"
	"github.com/Galdoba/t5/pkg/dice"
)

const (
	DM_SpectralOther = -1
	DM_SizeOther     = 2
	DM_SubType       = -1
)

func NewStar(dp *dice.Dicepool, options ...star.StarOption) (*star.Star, error) {
	st := star.Star{
		Position:     star.Primary,
		SpectralType: "?",
		Size:         "?",
		SubType:      "?",
		OrbitOffset:  -1,
		HZO:          -1,
		JumpShadow:   -1,
		OrbitIndex:   -1, //-1 - undefined; -2 - companion; -3 - rogue
	}
	for _, set := range options {
		set(&st)
	}
	if err := star.AssertPosition(st.Position); err != nil {
		return nil, err
	}
	if err := rollSpectral(dp, &st); err != nil {
		return nil, err
	}
	if err := rollSize(dp, &st); err != nil {
		return nil, err
	}
	if err := rollSubType(dp, &st); err != nil {
		return nil, err
	}
	st.HZO = HZO(st.Size, st.SpectralType)

	return &st, nil
}
