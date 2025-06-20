package basic

import "github.com/Galdoba/t5/internal/cosmology/star"

const (
	min_index = -6
	max_index = 6
)

var spectralMap = map[int]string{
	-6: "OB",
	-5: "A",
	-4: "A",
	-3: "F",
	-2: "F",
	-1: "G",
	0:  "G",
	1:  "K",
	2:  "K",
	3:  "M",
	4:  "M",
	5:  "M",
	6:  "BD",
}

var mapSizeO = map[int]star.SizeClass{
	-6: star.Size_Ia,
	-5: star.Size_Ia,
	-4: star.Size_Ib,
	-3: star.Size_II,
	-2: star.Size_III,
	-1: star.Size_III,
	0:  star.Size_III,
	1:  star.Size_V,
	2:  star.Size_V,
	3:  star.Size_V,
	4:  star.Size_IV,
	5:  star.Size_D,
	6:  star.Size_IV,
}

var mapSizeB = map[int]star.SizeClass{
	-6: star.Size_Ia,
	-5: star.Size_Ia,
	-4: star.Size_Ib,
	-3: star.Size_II,
	-2: star.Size_III,
	-1: star.Size_III,
	0:  star.Size_III,
	1:  star.Size_III,
	2:  star.Size_V,
	3:  star.Size_V,
	4:  star.Size_IV,
	5:  star.Size_D,
	6:  star.Size_IV,
}

var mapSizeA = map[int]star.SizeClass{
	-6: star.Size_Ia,
	-5: star.Size_Ia,
	-4: star.Size_Ib,
	-3: star.Size_II,
	-2: star.Size_III,
	-1: star.Size_IV,
	0:  star.Size_V,
	1:  star.Size_V,
	2:  star.Size_V,
	3:  star.Size_V,
	4:  star.Size_V,
	5:  star.Size_D,
	6:  star.Size_IV,
}

var mapSizeF = map[int]star.SizeClass{
	-6: star.Size_II,
	-5: star.Size_II,
	-4: star.Size_III,
	-3: star.Size_IV,
	-2: star.Size_V,
	-1: star.Size_V,
	0:  star.Size_V,
	1:  star.Size_V,
	2:  star.Size_V,
	3:  star.Size_V,
	4:  star.Size_VI,
	5:  star.Size_D,
	6:  star.Size_VI,
}

var mapSizeG = map[int]star.SizeClass{
	-6: star.Size_II,
	-5: star.Size_II,
	-4: star.Size_III,
	-3: star.Size_IV,
	-2: star.Size_V,
	-1: star.Size_V,
	0:  star.Size_V,
	1:  star.Size_V,
	2:  star.Size_V,
	3:  star.Size_V,
	4:  star.Size_VI,
	5:  star.Size_D,
	6:  star.Size_VI,
}

var mapSizeK = map[int]star.SizeClass{
	-6: star.Size_II,
	-5: star.Size_II,
	-4: star.Size_III,
	-3: star.Size_IV,
	-2: star.Size_V,
	-1: star.Size_V,
	0:  star.Size_V,
	1:  star.Size_V,
	2:  star.Size_V,
	3:  star.Size_V,
	4:  star.Size_VI,
	5:  star.Size_D,
	6:  star.Size_VI,
}

var mapSizeM = map[int]star.SizeClass{
	-6: star.Size_II,
	-5: star.Size_II,
	-4: star.Size_II,
	-3: star.Size_II,
	-2: star.Size_III,
	-1: star.Size_V,
	0:  star.Size_V,
	1:  star.Size_V,
	2:  star.Size_V,
	3:  star.Size_V,
	4:  star.Size_VI,
	5:  star.Size_D,
	6:  star.Size_VI,
}

func sizeMapByType(spectral star.SpectralType) map[int]star.SizeClass {
	switch spectral {
	default:
		return nil
	case star.SpectalType_O:
		return mapSizeO
	case star.SpectalType_B:
		return mapSizeB
	case star.SpectalType_A:
		return mapSizeA
	case star.SpectalType_F:
		return mapSizeF
	case star.SpectalType_G:
		return mapSizeG
	case star.SpectalType_K:
		return mapSizeK
	case star.SpectalType_M:
		return mapSizeM
	}
}

func HZO(size star.SizeClass, spec star.SpectralType) int {
	switch spec {
	case star.SpectalType_O:
		switch size {
		case star.Size_Ia:
			return 15
		case star.Size_Ib:
			return 15
		case star.Size_II:
			return 14
		case star.Size_III:
			return 13
		case star.Size_IV:
			return 12
		case star.Size_V:
			return 11
		case star.Size_D:
			return 1
		}
	case star.SpectalType_B:
		switch size {
		case star.Size_Ia:
			return 13
		case star.Size_Ib:
			return 13
		case star.Size_II:
			return 12
		case star.Size_III:
			return 11
		case star.Size_IV:
			return 10
		case star.Size_V:
			return 9
		case star.Size_D:
			return 0
		}
	case star.SpectalType_A:
		switch size {
		case star.Size_Ia:
			return 12
		case star.Size_Ib:
			return 11
		case star.Size_II:
			return 9
		case star.Size_III:
			return 7
		case star.Size_IV:
			return 7
		case star.Size_V:
			return 7
		case star.Size_D:
			return 0
		}
	case star.SpectalType_F:
		switch size {
		case star.Size_Ia:
			return 11
		case star.Size_Ib:
			return 10
		case star.Size_II:
			return 9
		case star.Size_III:
			return 6
		case star.Size_IV:
			return 6
		case star.Size_V:
			return 4
		case star.Size_VI:
			return 3
		case star.Size_D:
			return 0
		}
	case star.SpectalType_G:
		switch size {
		case star.Size_Ia:
			return 12
		case star.Size_Ib:
			return 10
		case star.Size_II:
			return 9
		case star.Size_III:
			return 7
		case star.Size_IV:
			return 5
		case star.Size_V:
			return 3
		case star.Size_VI:
			return 2
		case star.Size_D:
			return 0
		}
	case star.SpectalType_K:
		switch size {
		case star.Size_Ia:
			return 12
		case star.Size_Ib:
			return 10
		case star.Size_II:
			return 9
		case star.Size_III:
			return 8
		case star.Size_IV:
			return 5
		case star.Size_V:
			return 2
		case star.Size_VI:
			return 1
		case star.Size_D:
			return 0
		}
	case star.SpectalType_M:
		switch size {
		case star.Size_Ia:
			return 12
		case star.Size_Ib:
			return 11
		case star.Size_II:
			return 10
		case star.Size_III:
			return 9
		case star.Size_V:
			return 0
		case star.Size_VI:
			return 0
		case star.Size_D:
			return 0
		}
	}
	return 0
}
