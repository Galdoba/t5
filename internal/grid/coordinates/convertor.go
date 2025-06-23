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

func local_to_global(sx, sy, lx, ly int) (int, int) {
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

func global_to_local(x, y int) (int, int, int, int) {
	//reverse offset
	adj_x := x - loc2Glo_Offset_Width
	adj_y := y - loc2Glo_Offset_Height

	//x-axis
	temp := adj_x - 1
	sx := temp / sectorWidth
	rx := temp % sectorWidth //negative remainder
	//adjust negative remainder
	if rx < 0 {
		rx += sectorWidth
		sx--
	}
	lx := rx + 1

	//y-axis
	temp = adj_y - 1
	sy := temp / sectorHeight
	ry := temp % sectorHeight
	if ry < 0 {
		ry += sectorHeight
		sy--
	}
	ly := ry + 1
	return sx, sy, lx, ly
}

func local_to_cube(sx, sy, lx, ly int) (int, int, int) {
	// sx_adj := 0
	// sy_adj := 0
	// for lx > sectorWidth {
	// 	lx -= sectorWidth
	// 	sx_adj++
	// }
	// for lx < 1 {
	// 	lx += sectorWidth
	// 	sx_adj++
	// }
	// for ly > sectorWidth {
	// 	ly -= sectorHeight
	// 	sy_adj++
	// }
	// for ly < 1 {
	// 	ly += sectorHeight
	// 	sy_adj++
	// }
	// sx += sx_adj
	// sy += sy_adj

	x := lx + loc2Glo_Offset_Width + (sx * sectorWidth)
	y := ly + loc2Glo_Offset_Height + (sy * sectorHeight)

	// q-axis
	q := x

	// r-axis
	s := -(y + (x / 2))

	// s-axis
	r := -q - s
	return q, r, s
}

func cube_to_local(q, r, s int) (int, int, int, int) {
	x := q
	y := -((q / 2) + s)

	adj_x := x - loc2Glo_Offset_Width
	adj_y := y - loc2Glo_Offset_Height
	//x-axis
	temp := adj_x - 1
	sx := temp / sectorWidth
	rx := temp % sectorWidth //negative remainder
	//adjust negative remainder
	if rx < 0 {
		rx += sectorWidth
		sx--
	}
	lx := rx + 1

	//y-axis
	temp = adj_y - 1
	sy := temp / sectorHeight
	ry := temp % sectorHeight
	if ry < 0 {
		ry += sectorHeight
		sy--
	}
	ly := ry + 1

	return sx, sy, lx, ly
}

func cube_to_global(q, r, s int) (int, int) {
	x := q
	y := ((q / 2) + s) * -1
	return x, y
}

func global_to_cube(x, y int) (int, int, int) {
	q := x
	s := -(y + (x / 2))
	r := -q - s
	return q, r, s
}

func roundTrip_Cube(q, r, s int) error {
	gx, gy := cube_to_global(q, r, s)
	sx, sy, lx, ly := global_to_local(gx, gy)
	rq, rr, rs := local_to_cube(sx, sy, lx, ly)
	if q != rq || r != rr || s != rs {
		return fmt.Errorf("round trip failed: cube(%v, %v, %v)->global(%v, %v)->local(%v, %v, %v, %v)->return(%v, %v, %v)", q, r, s, gx, gy, sx, sy, lx, ly, rq, rr, rs)
	}
	return nil
}
func roundTrip_CubeBack(q, r, s int) error {
	sx, sy, lx, ly := cube_to_local(q, r, s)
	gx, gy := local_to_global(sx, sy, lx, ly)
	rq, rr, rs := global_to_cube(gx, gy)
	if q != rq || r != rr || s != rs {
		return fmt.Errorf("round trip failed: cube(%v, %v, %v)->local(%v, %v, %v, %v)->global(%v, %v)->return(%v, %v, %v)", q, r, s, sx, sy, lx, ly, gx, gy, rq, rr, rs)
	}
	return nil
}

func roundTrip(q, r, s int) error {
	gx, gy := cube_to_global(q, r, s)
	sx, sy, lx, ly := global_to_local(gx, gy)
	rq, rr, rs := local_to_cube(sx, sy, lx, ly)
	if q != rq || r != rr || s != rs {
		return fmt.Errorf("round trip failed: cube(%v, %v, %v)->global(%v, %v)->local(%v, %v, %v, %v)->return(%v, %v, %v)", q, r, s, gx, gy, sx, sy, lx, ly, rq, rr, rs)
	}
	sx, sy, lx, ly = cube_to_local(q, r, s)
	gx, gy = local_to_global(sx, sy, lx, ly)
	rq, rr, rs = global_to_cube(gx, gy)
	if q != rq || r != rr || s != rs {
		return fmt.Errorf("round trip failed: cube(%v, %v, %v)->local(%v, %v, %v, %v)->global(%v, %v)->return(%v, %v, %v)", q, r, s, sx, sy, lx, ly, gx, gy, rq, rr, rs)
	}
	return nil
}
