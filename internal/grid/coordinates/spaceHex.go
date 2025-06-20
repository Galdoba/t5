package coordinates

type GlobalWorldCoordinates struct {
	Snail  int `json:"snail"`
	Double Double
	Cube   Cube
}

var GroundZero = GlobalWorldCoordinates{
	Snail:  0,
	Double: Double{0, 0},
	Cube:   Cube{0, 0, 0},
}

type Double struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Cube struct {
	Q int `json:"q"`
	R int `json:"r"`
	S int `json:"s"`
}

func NewGlobalCoordinates(crd ...int) GlobalWorldCoordinates {
	coords := GroundZero
	switch len(crd) {
	case 0:
		return GroundZero
	case 1:
		coords.Snail = crd[0]
	case 2:
		coords.Double = Double{crd[0], crd[1]}
	default:
		coords.Cube = Cube{crd[0], crd[1], crd[2]}
	}
	return coords
}
