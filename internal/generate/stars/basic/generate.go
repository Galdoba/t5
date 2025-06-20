package basic

import (
	"fmt"
	"strings"

	"github.com/Galdoba/t5/internal/cosmology/star"
	"github.com/Galdoba/t5/pkg/dice"
)

func rollSpectral(dp *dice.Dicepool, st *star.Star) error {
	if st.SpectralType != star.SpectalType_Undefined {
		return star.AssertSpectralType(st.SpectralType)
	}
	index := 0
	switch st.Position {
	case star.Primary:
		index = dp.Flux() + dp.Flux()
	default:
		index = dp.Flux() + dp.Sum("2d6", dice.DM(DM_SpectralOther))
	}
	index = boundInt(index, min_index, max_index)
	spec := spectralMap[index]
	if spec == "OB" {
		spec = strings.Split(spec, "")[dp.Sum("1d2", dice.DM(-1))]
	}
	st.SpectralType = star.SpectralType(spec)
	return star.AssertSpectralType(st.SpectralType)
}

func rollSize(dp *dice.Dicepool, st *star.Star) error {
	if st.Size != "?" {
		return star.AssertSizeClass(st.Size)
	}
	if st.SpectralType == star.SpectalType_BD {
		st.Size = star.Size_NUL
		return star.AssertSizeClass(st.Size)
	}
	index := 0
	switch st.Position {
	case star.Primary:
		index = dp.Flux() + dp.Flux()
	default:
		index = dp.Flux() + dp.Sum("2d6", dice.DM(DM_SizeOther))
	}
	index = boundInt(index, min_index, max_index)
	sizeMap := sizeMapByType(st.SpectralType)
	st.Size = star.SizeClass(sizeMap[index])
	return star.AssertSizeClass(st.Size)
}

func rollSubType(dp *dice.Dicepool, st *star.Star) error {
	if st.SpectralType == star.SpectalType_BD {
		st.SubType = star.SubType_NUL
		return star.AssertSubType(st.SubType)
	}
	if st.Size == star.Size_D {
		st.SubType = star.SubType_NUL
		return star.AssertSubType(st.SubType)
	}

	subtype := fmt.Sprintf("%v", dp.Sum("1d10", dice.DM(DM_SubType)))
	switch st.Size {
	case star.Size_IV:
		switch st.SpectralType {
		case star.SpectalType_M:
			st.Size = star.Size_V
		case star.SpectalType_K:
			switch subtype {
			case "5", "6", "7", "8", "9":
				st.Size = star.Size_V
			}
		}
	case star.Size_VI:
		switch st.SpectralType {
		case star.SpectalType_A, star.SpectalType_O, star.SpectalType_B:
			st.Size = star.Size_V
		case star.SpectalType_F:
			switch subtype {
			case "0", "1", "2", "3", "4":
				st.Size = star.Size_V
			}
		}
	default:
	}
	st.SubType = star.SubType(subtype)
	return star.AssertSubType(st.SubType)
}

func boundInt(i, min, max int) int {
	if i < min {
		return min
	}
	if i > max {
		return max
	}
	return i
}
