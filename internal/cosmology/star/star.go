package star

import "fmt"

type Star struct {
	Position     SystemPosition //P Pc C Cc N Nc F Fc
	SpectralType SpectralType   //O B A F G K M BD
	Size         SizeClass      //Ia Ib II III IV V VI D ""
	SubType      SubType        //0 1 2 3 4 5 6 7 8 9 ""
	OrbitOffset  int            //planets in orbits lower than this are consumed/burned
	HZO          int            //Habitable Orbit
	JumpShadow   int            //orbits equal or lower are in star's shadow
	OrbitIndex   int            //position relative to hex center (-2 for rogue)
}

func (st *Star) String() string {
	if st.Size == Size_D {
		return fmt.Sprintf("%v%v", st.Size, st.SpectralType)
	}
	if st.SpectralType == SpectalType_BD {
		return fmt.Sprintf("%v", st.SpectralType)
	}
	return fmt.Sprintf("%v%v %v", st.SpectralType, st.SubType, st.Size)
}

func (st *Star) DebugText() string {
	return fmt.Sprintf("\toffset: %v\tHZ: %v\tJS: %v", st.OrbitOffset, st.HZO, st.JumpShadow)
}
