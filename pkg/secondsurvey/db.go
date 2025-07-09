package secondsurvey

import "github.com/Galdoba/t5/internal/grid/coordinates"

type DB struct {
	path     string
	gridData map[coordinates.SpaceCoordinates]string
}

// c:\Users\pemaltynov\travellermap\res\Sectors\M1105\Core.tab
//
