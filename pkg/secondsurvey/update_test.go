package secondsurvey

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Galdoba/t5/internal/grid/coordinates"
	"github.com/Galdoba/t5/internal/grid/coordinates/sector"
)

func TestFetchOUT(t *testing.T) {
	err := FetchOTU()
	if err != nil {
		fmt.Println(err)
		panic(0)
	}

	sect := ReadFile(`c:\Users\pemaltynov\travellermap\res\Sectors\sectors.json`)
	fmt.Println(len(sect.Sectors))
	crd := make(map[string]coordinates.SpaceCoordinates)

	for _, sector := range sect.Sectors {
		for _, name := range sector.Names {
			crd[name.Text] = coordinates.NewSpaceCoordinates(sector.X, sector.Y, 1, 1)
		}
	}
	// for k, v := range crd {
	// 	fmt.Printf("sector %v = cor: %v\n", k, v)
	// }
	fmt.Println(len(crd))
	// fileMap := MapSectorDataFiles()
	// connector := sector.NewCommector()
	// i := 0
	// worlds := 0
	// for sectorName, sectorFilePath := range fileMap {
	// 	//fmt.Printf("sector %v = file: %v\n", sectorName, sectorFilePath)
	// 	for _, sector := range sect.Sectors {
	// 		for _, name := range sector.Names {
	// 			if name.Text == sectorName && sector.Milieu == "M1105" {
	// 				//fmt.Printf("sector %v = coords: %v [%v]\n", sectorName, crd[sectorName], sector.Milieu)
	// 				connector.Add(sector.X, sector.Y, sectorName, sectorFilePath)
	// 				if err := connector.Save(); err != nil {
	// 					panic(err)
	// 				}
	// 				i++
	// 				for localX := 1; localX <= 32; localX++ {
	// 					for localY := 1; localY <= 40; localY++ {
	// 						coors := coordinates.NewSpaceCoordinates(sector.X, sector.Y, localX, localY)
	// 						for _, line := range FileLines(sectorFilePath) {
	// 							tabVals := strings.Split(line, "\t")
	// 							for _, val := range tabVals {
	// 								if val == coors.SectorHex() {
	// 									worlds++
	// 									fmt.Println(worlds, line)

	// 								}
	// 							}
	// 						}
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}

	// }
	// fmt.Println(len(fileMap))
	// fmt.Println(len(crd))
	// fmt.Println(i)
	c, err := sector.LoadConnector()
	if err != nil {
		panic(err)
	}
	// worlds := 0
	sm := NewSpaceMap()
	unparsed := []string{}
	worldCound := 0
	sseMap := make(map[string]SecondSurveyExtended)
	sseKeys := []string{}
	for scs, currentSector := range c.Map {
		scs = strings.TrimPrefix(scs, "{")
		scs = strings.TrimSuffix(scs, "{")
		// scsP := strings.Split(scs, ",")
		// sectorX, _ := strconv.Atoi(scsP[0])
		// sectorY, _ := strconv.Atoi(scsP[1])
		file := currentSector.File

		surv, err := ParseFile(file)
		if err != nil {
			unparsed = append(unparsed, fmt.Sprintf("unparsed %v: %v", file, err))
			continue
		}
		for _, sse := range surv {

			lx, ly := sector.ParseLocalFromHex(sse.Hex)
			//			fmt.Println(sse.Hex, currentSector.X, currentSector.Y, lx, ly)
			crd := coordinates.NewSpaceCoordinates(currentSector.X, currentSector.Y, lx, ly)
			sse = InjectCoordinates(sse, crd)
			worldCound++
			sseKeys = append(sseKeys, sse.String())
			sseMap[sse.String()] = sse
			fmt.Println(sse.String())
		}

		// for lNum, line := range FileLines(file) {
		// 	fmt.Println(lNum, parseFormat(file))
		// for localX := 1; localX <= 32; localX++ {
		// 	for localY := 1; localY <= 40; localY++ {
		// 		hex := sector.Hex(localX, localY)
		// 		crd := coordinates.NewSpaceCoordinates(sectorX, sectorY, localX, localY)
		// 		tabValues := strings.Split(line, "\t")
		// 		for _, tVal := range tabValues {
		// 			if hex != tVal {
		// 				continue
		// 			}
		// 			ssr, err := Parse(line)
		// 			if err != nil {
		// 				fmt.Println(file)
		// 				fmt.Println(lNum, line)
		// 				fmt.Println(err)
		// 				time.Sleep(time.Second * 2)
		// 				continue
		// 			}
		// 			sm.Add(crd, AddString(SK_Hex, hex),
		// 				AddString(SK_Sector, currentSector.Name),
		// 				AddString(SK_UWP, ssr.UWP),
		// 				AddString(SK_Remarks, strings.Join(ssr.Remarks, " ")),
		// 				AddString(SK_Importance, ssr.Importance),
		// 				AddString(SK_Economic, ssr.Economic),
		// 				AddString(SK_Culture, ssr.Culture),
		// 				AddString(SK_Nobility, strings.Join(ssr.Nobility, "")),
		// 				AddString(SK_Bases, strings.Join(ssr.Bases, "")),
		// 				AddString(SK_Zone, ssr.Zone),
		// 				AddString(SK_PBG, ssr.PBG),
		// 				AddString(SK_Worlds, ssr.Worlds),
		// 				AddString(SK_Allegiance, ssr.Allegiance),
		// 				AddString(SK_Stellar, ssr.Stellar),
		// 				AddCoordinates(crd))

		// 			worlds++
		// 			fmt.Printf("worlds parsed: %v\r", worlds)
		// 		}
		// 	}
		// }
		// }
	}
	for k, v := range unparsed {
		fmt.Println(k, v)
	}
	if err := sm.Save(`c:\Users\pemaltynov\travellermap\res\merge\compiledSurveyData.json`); err != nil {
		panic(err)
	}

}

func FileLines(path string) []string {
	bt, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(bt), "\n")
}
