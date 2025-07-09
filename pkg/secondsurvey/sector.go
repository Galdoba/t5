package secondsurvey

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type SectorInfo struct {
	Sectors []SectorData `json:"Sectors,omitempty"`
}

type SectorData struct {
	X            int          `json:"X,omitempty"`
	Y            int          `json:"Y,omitempty"`
	Milieu       string       `json:"Milieu,omitempty"`
	Abbreviation string       `json:"Abbreviation,omitempty"`
	Tags         string       `json:"Tags,omitempty"`
	Names        []SectorName `json:"Names,omitempty"`
}

type SectorName struct {
	Text string `json:"Text,omitempty"`
	Lang string `json:"Lang,omitempty"`
}

func ReadFile(path string) SectorInfo {
	sectors := SectorInfo{}
	bt, err := os.ReadFile(path)
	fmt.Println(len(bt), "bytes")
	if err != nil {
		fmt.Println(err)
		return sectors
	}

	if err := json.Unmarshal(bt, &sectors); err != nil {
		fmt.Println(err)
		return sectors
	}
	return sectors
}

func MapSectorDataFiles() map[string]string {
	files := make(map[string]string)
	fi, err := os.ReadDir(`c:\Users\pemaltynov\travellermap\res\Sectors\M1105\`)
	if err != nil {
		panic(err)
	}
	for _, f := range fi {
		name := f.Name()
		if !strings.HasSuffix(name, ".tab") {
			continue
		}
		files[strings.TrimSuffix(name, ".tab")] = `c:\Users\pemaltynov\travellermap\res\Sectors\M1105\` + name
	}
	return files
}
