package starsystem

type StarSystem struct {
	Profile     string
	OrbitLayout map[float64]Orbit
}

type Orbit struct {
	Index        int
	AU           float64
	OrbitN       float64
	Status       string
	IsRogue      bool
	Parent       string
	Occupant     string
	Eccentricity float64
}

func NewOrbit(index int) *Orbit {
	orb := Orbit{
		Index:        index,
		AU:           0,
		OrbitN:       0,
		Status:       "",
		IsRogue:      false,
		Parent:       "",
		Occupant:     "",
		Eccentricity: 0,
	}
	return &orb
}
