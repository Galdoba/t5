package coordinates

import "fmt"

const (
	sectorWidth           = 32
	sectorHeight          = 40
	loc2Glo_Offset_Width  = -1
	loc2Glo_Offset_Height = -40
)

type Combined struct {
	Q, R, S          int
	SectorX, SectorY int
	LocalX, LocalY   int
	GlobalX, GlobalY int
}

func localToGlobal(sx, sy, lx, ly int) (int, int) {
	if lx > sectorWidth || lx < 1 {
		panic(fmt.Sprintf("lx (%v) not in range 1-32", lx))
	}
	if ly > sectorHeight || ly < 1 {
		panic(fmt.Sprintf("ly (%v) not in range 1-40", ly))
	}
	x := lx + loc2Glo_Offset_Width + (sx * sectorWidth)
	y := ly + loc2Glo_Offset_Height + (sy * sectorHeight)
	return x, y
}

func localToCube(sx, sy, lx, ly int) (int, int, int) {
	if lx > sectorWidth || lx < 1 {
		panic(fmt.Sprintf("lx (%v) not in range 1-32", lx))
	}
	if ly > sectorHeight || ly < 1 {
		panic(fmt.Sprintf("ly (%v) not in range 1-40", ly))
	}
	// q = x - 1 (индексация с 0)
	q := lx + loc2Glo_Offset_Width + (sx * sectorWidth)

	// Вычисление r с учетом смещения столбцов
	r_orig := (ly) - ((q - (q & 1)) / 2) + (sx * sectorHeight * -1) + loc2Glo_Offset_Height

	// s вычисляется из кубического условия
	s_orig := -q - r_orig

	// Применяем вертикальный сдвиг: перемещаем начало координат в нижний левый угол
	r := r_orig //localToCube+ (sx * sectorHeight)
	s := s_orig //+ (sx * sectorHeight * -1)
	return q, r, s
}
