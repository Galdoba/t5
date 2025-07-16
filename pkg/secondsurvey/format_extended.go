package secondsurvey

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/Galdoba/t5/pkg/grid/coordinates"
)

const (
	SK_Sector         = "Sector"
	SK_Subsector      = "Subsector"
	SK_Hex            = "Hex"
	SK_Name           = "Name"
	SK_UWP            = "UWP"
	SK_Remarks        = "Remarks"
	SK_Importance     = "Importance"
	SK_Economic       = "Economic"
	SK_Culture        = "Culture"
	SK_Nobility       = "Nobility"
	SK_Bases          = "Bases"
	SK_Zone           = "Zone"
	SK_PBG            = "PBG"
	SK_Worlds         = "Worlds"
	SK_Allegiance     = "Allegiance"
	SK_Stellar        = "Stellar"
	SK_SectorAbb      = "SectorAbb"
	SK_SubSectorAbb   = "SubSectorAbb"
	SK_SubSectorIndex = "SubSectorIndex"
	SK_GlobalX        = "GlobalX"
	SK_GlobalY        = "GlobalY"
	SK_SectorX        = "SectorX"
	SK_SectorY        = "SectorY"
	SK_LocalX         = "LocalX"
	SK_LocalY         = "LocalY"
	SK_CubeQ          = "CubeQ"
	SK_CubeR          = "CubeR"
	SK_CubeS          = "CubeS"
	SK_RU             = "RU"
)

type SecondSurveyExtended struct {
	Sector            string
	Subsector         string
	Hex               string
	Name              string
	UWP               string
	Remarks           []string
	Importance        string
	Economic          string
	Culture           string
	Nobility          []string
	Bases             []string
	Zone              string
	PBG               string
	Worlds            string
	Allegiance        string
	Stellar           string
	SectorAbb         string
	SubSectorAbb      string
	SubSectorIndex    string
	CoordinatesFilled bool
	GlobalX           int
	GlobalY           int
	SectorX           int
	SectorY           int
	LocalX            int
	LocalY            int
	CubeQ             int
	CubeR             int
	CubeS             int
	RU                int
}

func (sse SecondSurveyExtended) String() string {
	s := ""
	switch sse.CoordinatesFilled {
	case true:
		s += fmt.Sprintf("{%v,%v}", sse.GlobalX, sse.GlobalY)
	case false:
		s += fmt.Sprintf("{???,???}")
	}
	s += "\t"
	s += fmt.Sprintf("%v\t", sse.Sector)
	s += fmt.Sprintf("%v\t", sse.Hex)
	s += fmt.Sprintf("%v\t", sse.Name)
	s += fmt.Sprintf("%v\t", sse.UWP)
	s += fmt.Sprintf("%v\t", sse.Remarks)
	s += fmt.Sprintf("%v\t", sse.Bases)
	s += fmt.Sprintf("%v\t", sse.Zone)
	s += fmt.Sprintf("%v\t", sse.Subsector)
	s += fmt.Sprintf("%v\t", sse.Importance)
	s += fmt.Sprintf("%v\t", sse.Economic)
	s += fmt.Sprintf("%v\t", sse.Culture)
	s += fmt.Sprintf("%v\t", sse.Nobility)
	s += fmt.Sprintf("%v\t", sse.PBG)
	s += fmt.Sprintf("%v\t", sse.Worlds)
	s += fmt.Sprintf("%v\t", sse.Allegiance)
	s += fmt.Sprintf("%v", sse.Stellar)
	return s
}

func FillSurvey(data ...SurveyData) SecondSurveyExtended {
	sse := SecondSurveyExtended{}
	for _, fill := range data {
		fill(&sse)
	}
	return sse
}

type SurveyData func(*SecondSurveyExtended)

func AddData(key, value string) SurveyData {
	return func(sse *SecondSurveyExtended) {
		switch key {
		case SK_Sector:
			sse.Sector = value
		case SK_Subsector:
			sse.Subsector = value
		case SK_Hex:
			sse.Hex = value
		case SK_Name:
			sse.Name = value
		case SK_UWP:
			sse.UWP = value
		case SK_Remarks:
			sse.Remarks = strings.Split(value, " ")
		case SK_Importance:
			sse.Importance = value
		case SK_Economic:
			sse.Economic = value
		case SK_Culture:
			sse.Culture = value
		case SK_Nobility:
			sse.Nobility = strings.Split(value, "")
		case SK_Bases:
			sse.Bases = strings.Split(value, "")
		case SK_Zone:
			sse.Zone = value
		case SK_PBG:
			sse.PBG = value
		case SK_Worlds:
			sse.Worlds = value
		case SK_Allegiance:
			sse.Allegiance = value
		case SK_Stellar:
			sse.Stellar = value
		case SK_SectorAbb:
			sse.SectorAbb = value
		case SK_SubSectorAbb:
			sse.SubSectorAbb = value
		case SK_SubSectorIndex:
			sse.SubSectorIndex = value
		case SK_RU:
			ru, err := strconv.Atoi(value)
			if err != nil {
				ru = math.MinInt
			}
			sse.RU = ru
		}
	}
}

func AddInt(key string, value int) SurveyData {
	return func(sse *SecondSurveyExtended) {
		switch key {
		case SK_GlobalX:
			sse.GlobalX = value
		case SK_GlobalY:
			sse.GlobalY = value
		case SK_SectorX:
			sse.SectorX = value
		case SK_SectorY:
			sse.SectorY = value
		case SK_LocalX:
			sse.LocalX = value
		case SK_LocalY:
			sse.LocalY = value
		case SK_CubeQ:
			sse.CubeQ = value
		case SK_CubeR:
			sse.CubeR = value
		case SK_CubeS:
			sse.CubeS = value
		}
	}
}

func AddCoordinates(crd coordinates.SpaceCoordinates) SurveyData {
	return func(sse *SecondSurveyExtended) {
		q, r, s := crd.CubeValues()
		x, y := crd.GlobalValues()
		sx, sy, lx, ly := crd.LocalValues()
		sse.GlobalX = x
		sse.GlobalY = y
		sse.CubeQ = q
		sse.CubeR = r
		sse.CubeS = s
		sse.SectorX = sx
		sse.SectorY = sy
		sse.LocalX = lx
		sse.LocalY = ly

	}
}

func InjectCoordinates(sse SecondSurveyExtended, crd coordinates.SpaceCoordinates) SecondSurveyExtended {
	sse.CubeQ, sse.CubeR, sse.CubeS = crd.CubeValues()
	sse.GlobalX, sse.GlobalY = crd.GlobalValues()
	sse.SectorX, sse.SectorY, sse.LocalX, sse.LocalY = crd.LocalValues()
	sse.CoordinatesFilled = true
	return sse
}

type SpaceMap struct {
	Map map[string]SecondSurveyExtended
}

func NewSpaceMap() *SpaceMap {
	sm := SpaceMap{}
	sm.Map = make(map[string]SecondSurveyExtended)
	return &sm
}

func (sm *SpaceMap) Add(crd coordinates.SpaceCoordinates, data ...SurveyData) {
	q, r, s := crd.CubeValues()
	key := fmt.Sprintf("{%v,%v,%v}", q, r, s)
	survey := FillSurvey(data...)
	sm.Map[key] = survey
}

func (sm *SpaceMap) Save(path string) error {

	bt, err := json.MarshalIndent(sm, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, bt, 0755)
}

func (sm *SpaceMap) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, sm); err != nil {
		return fmt.Errorf("failed to unmarshal: %v", err)
	}
	return nil
}

func parseFormat(lines []string) (string, int) {
	for l, line := range lines {
		if strings.Contains(line, "UWP") && strings.Contains(line, "Name") {
			return line, l
		}
	}
	return fmt.Sprintf("format unknown"), -1
}
