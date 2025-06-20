package star

import "fmt"

func AssertPosition(pos SystemPosition) error {
	switch pos {
	default:
		return fmt.Errorf("invalid star position '%v'", pos)
	case Primary, PrimaryCompanion, Close, CloseCompanion, Near, NearCompanion, Far, FarCompanion, Rogue:
	}
	return nil
}

func AssertSpectralType(spectralType SpectralType) error {
	switch spectralType {
	default:
		return fmt.Errorf("invalid star spectral type '%v'", spectralType)
	case SpectalType_O, SpectalType_B, SpectalType_A, SpectalType_F, SpectalType_G, SpectalType_K, SpectalType_M, SpectalType_BD:
	}
	return nil
}

func AssertSizeClass(size SizeClass) error {
	switch size {
	default:
		return fmt.Errorf("invalid star size '%v'", size)
	case Size_Ia, Size_Ib, Size_II, Size_III, Size_IV, Size_V, Size_VI, Size_D, Size_NUL:
	}
	return nil
}

func AssertSubType(subtype SubType) error {
	switch subtype {
	default:
		return fmt.Errorf("invalid star subtype '%v'", subtype)
	case SubType_0, SubType_1, SubType_2, SubType_3, SubType_4, SubType_5, SubType_6, SubType_7, SubType_8, SubType_9, SubType_NUL:
	}
	return nil
}
