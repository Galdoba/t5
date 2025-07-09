package secondsurvey

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
1 Hex   Name    UWP     Comments        N       B       Z       PBG     A       Stellar
1 Sector        SS      Hex     Name    UWP     Bases   Remarks Zone    PBG     Allegiance      Stars   {Ix}    (Ex)    [Cx]    Nobility        W
72 Sector       SS      Hex     Name    UWP     Bases   Remarks Zone    PBG     Allegiance      Stars   {Ix}    (Ex)    [Cx]    Nobility        W       RU
4 Hex   Name    UWP     Bases   Remarks Zone    PBG     Allegiance      Stars   {Ix}    (Ex)    [Cx]    Nobility        W
1 Hex   Name    UWP     Remarks { Ix }  ( Ex )  [ Cx ]  N       B       Z       PBG     W       A       Stellar
5 Hex   Name    UWP     Remarks {Ix}    (Ex)    [Cx]    N       B       Z       PBG     W       A       Stellar
1 Sector        SS      Hex     Name    UWP     Bases   Remarks Zone    PBG     Allegiance      Stars   {Ix}    (Ex)    [Cx]    Noblity W       RU
*/

const (
	format1 = `Sector	SS	Hex	Name	UWP	Bases	Remarks	Zone	PBG	Allegiance	Stars	{Ix}	(Ex)	[Cx]	Nobility	W	RU`
	format2 = `Hex	Name	UWP	Comments	N	B	Z	PBG	A	Stellar`
	format3 = `Hex	Name	UWP	Bases	Remarks	Zone	PBG	Allegiance	Stars	{Ix}	(Ex)	[Cx]	Nobility	W`
	format4 = `Sector	SS	Hex	Name	UWP	Bases	Remarks	Zone	PBG	Allegiance	Stars	{Ix}	(Ex)	[Cx]	Nobility	W`
	format5 = `Hex	Name	UWP	Remarks	{ Ix }	( Ex )	[ Cx ]	N	B	Z	PBG	W	A	Stellar`
	format6 = `Hex	Name	UWP	Remarks	{Ix}	(Ex)	[Cx]	N	B	Z	PBG	W	A	Stellar`
	format7 = `Sector	SS	Hex	Name	UWP	Bases	Remarks	Zone	PBG	Allegiance	Stars	{Ix}	(Ex)	[Cx]	Noblity	W	RU`
)

var fillSurvey = map[string]func([]string) SecondSurveyExtended{
	format1: fillWithFormat1,
	format2: fillWithFormat2,
	format3: fillWithFormat3,
	format4: fillWithFormat4,
	format5: fillWithFormat5,
	format6: fillWithFormat6,
	format7: fillWithFormat7,
}

func ParseFile(file string) ([]SecondSurveyExtended, error) {
	bt, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	lines := strings.Split(string(bt), "\n")
	surveyed := []SecondSurveyExtended{}
	format, fmtLineNum := parseFormat(lines)
	format = strings.TrimSuffix(format, "\r")
	sectorName := strings.TrimSuffix(filepath.Base(file), ".tab")

	switch format {
	case format1, format2, format3, format4, format5, format6, format7:
		for _, line := range lines[fmtLineNum+1:] {
			line = strings.TrimSuffix(line, "\r")
			if line == "" {
				continue
			}
			data := strings.Split(line, "\t")
			surveyed = append(surveyed, fillSurvey[format](data))
		}

	default:
		fmt.Println("file", file)
		fmt.Println("format unknowm", format)

	}
	for i := range surveyed {
		surveyed[i].Sector = sectorName
	}
	return surveyed, nil
}

func fillWithFormat1(data []string) SecondSurveyExtended {
	if len(data) != 17 {
		return SecondSurveyExtended{Name: "[PARSE_ERROR_FORMAT_1]", UWP: "line=" + strings.Join(data, "|")}
	}
	sse := FillSurvey(
		AddData(SK_SectorAbb, data[0]),
		AddData(SK_SubSectorIndex, data[1]),
		AddData(SK_Hex, data[2]),
		AddData(SK_Name, data[3]),
		AddData(SK_UWP, data[4]),
		AddData(SK_Bases, data[5]),
		AddData(SK_Remarks, data[6]),
		AddData(SK_Zone, data[7]),
		AddData(SK_PBG, data[8]),
		AddData(SK_Allegiance, data[9]),
		AddData(SK_Stellar, data[10]),
		AddData(SK_Importance, data[11]),
		AddData(SK_Economic, data[12]),
		AddData(SK_Culture, data[13]),
		AddData(SK_Nobility, data[14]),
		AddData(SK_Worlds, data[15]),
		AddData(SK_RU, data[16]),
	)
	return sse
}

// `Hex	Name	UWP	Comments	N	B	Z	PBG	A	Stellar`
func fillWithFormat2(data []string) SecondSurveyExtended {
	if len(data) != 10 {
		return SecondSurveyExtended{Name: "[PARSE_ERROR_FORMAT_2]", UWP: "line=" + strings.Join(data, "|")}
	}
	sse := FillSurvey(
		AddData(SK_Hex, data[0]),
		AddData(SK_Name, data[1]),
		AddData(SK_UWP, data[2]),
		AddData(SK_Remarks, data[3]),
		AddData(SK_Nobility, data[4]),
		AddData(SK_Bases, data[5]),
		AddData(SK_Zone, data[6]),
		AddData(SK_PBG, data[7]),
		AddData(SK_Allegiance, data[8]),
		AddData(SK_Stellar, data[9]),
	)
	return sse
}

// Hex	Name	UWP	Bases	Remarks	Zone	PBG	Allegiance	Stars	{Ix}	(Ex)	[Cx]	Nobility	W
func fillWithFormat3(data []string) SecondSurveyExtended {
	if len(data) != 14 {
		return SecondSurveyExtended{Name: "[PARSE_ERROR_FORMAT_3]", UWP: "line=" + strings.Join(data, "|")}
	}
	sse := FillSurvey(
		AddData(SK_Hex, data[0]),
		AddData(SK_Name, data[1]),
		AddData(SK_UWP, data[2]),
		AddData(SK_Bases, data[3]),
		AddData(SK_Remarks, data[4]),
		AddData(SK_Zone, data[5]),
		AddData(SK_PBG, data[6]),
		AddData(SK_Allegiance, data[7]),
		AddData(SK_Stellar, data[8]),
		AddData(SK_Importance, data[9]),
		AddData(SK_Economic, data[10]),
		AddData(SK_Culture, data[11]),
		AddData(SK_Nobility, data[12]),
		AddData(SK_Worlds, data[13]),
	)
	return sse
}

func fillWithFormat4(data []string) SecondSurveyExtended {
	if len(data) != 16 {
		return SecondSurveyExtended{Name: "[PARSE_ERROR_FORMAT_4]", UWP: "line=" + strings.Join(data, "|")}
	}
	sse := FillSurvey(
		AddData(SK_SectorAbb, data[0]),
		AddData(SK_SubSectorIndex, data[1]),
		AddData(SK_Hex, data[2]),
		AddData(SK_Name, data[3]),
		AddData(SK_UWP, data[4]),
		AddData(SK_Bases, data[5]),
		AddData(SK_Remarks, data[6]),
		AddData(SK_Zone, data[7]),
		AddData(SK_PBG, data[8]),
		AddData(SK_Allegiance, data[9]),
		AddData(SK_Stellar, data[10]),
		AddData(SK_Importance, data[11]),
		AddData(SK_Economic, data[12]),
		AddData(SK_Culture, data[13]),
		AddData(SK_Nobility, data[14]),
		AddData(SK_Worlds, data[15]),
	)
	return sse
}

// format5 = `Hex	Name	UWP	Remarks	{ Ix }	( Ex )	[ Cx ]	N	B	Z	PBG	W	A	Stellar`
func fillWithFormat5(data []string) SecondSurveyExtended {
	if len(data) != 14 {
		return SecondSurveyExtended{Name: "[PARSE_ERROR_FORMAT_5]", UWP: "line=" + strings.Join(data, "|")}
	}
	sse := FillSurvey(
		AddData(SK_Hex, data[0]),
		AddData(SK_Name, data[1]),
		AddData(SK_UWP, data[2]),
		AddData(SK_Remarks, data[3]),
		AddData(SK_Importance, data[4]),
		AddData(SK_Economic, data[5]),
		AddData(SK_Culture, data[6]),
		AddData(SK_Nobility, data[7]),
		AddData(SK_Bases, data[8]),
		AddData(SK_Zone, data[9]),
		AddData(SK_PBG, data[10]),
		AddData(SK_Worlds, data[11]),
		AddData(SK_Allegiance, data[12]),
		AddData(SK_Stellar, data[13]),
	)
	return sse
}

// format5 = `Hex	Name	UWP	Remarks	{ Ix }	( Ex )	[ Cx ]	N	B	Z	PBG	W	A	Stellar`
func fillWithFormat6(data []string) SecondSurveyExtended {
	if len(data) != 14 {
		return SecondSurveyExtended{Name: "[PARSE_ERROR_FORMAT_6]", UWP: "line=" + strings.Join(data, "|")}
	}
	sse := FillSurvey(
		AddData(SK_Hex, data[0]),
		AddData(SK_Name, data[1]),
		AddData(SK_UWP, data[2]),
		AddData(SK_Remarks, data[3]),
		AddData(SK_Importance, data[4]),
		AddData(SK_Economic, data[5]),
		AddData(SK_Culture, data[6]),
		AddData(SK_Nobility, data[7]),
		AddData(SK_Bases, data[8]),
		AddData(SK_Zone, data[9]),
		AddData(SK_PBG, data[10]),
		AddData(SK_Worlds, data[11]),
		AddData(SK_Allegiance, data[12]),
		AddData(SK_Stellar, data[13]),
	)
	return sse
}

func fillWithFormat7(data []string) SecondSurveyExtended {
	if len(data) != 17 {
		return SecondSurveyExtended{Name: "[PARSE_ERROR_FORMAT_7]", UWP: "line=" + strings.Join(data, "|")}
	}
	sse := FillSurvey(
		AddData(SK_SectorAbb, data[0]),
		AddData(SK_SubSectorIndex, data[1]),
		AddData(SK_Hex, data[2]),
		AddData(SK_Name, data[3]),
		AddData(SK_UWP, data[4]),
		AddData(SK_Bases, data[5]),
		AddData(SK_Remarks, data[6]),
		AddData(SK_Zone, data[7]),
		AddData(SK_PBG, data[8]),
		AddData(SK_Allegiance, data[9]),
		AddData(SK_Stellar, data[10]),
		AddData(SK_Importance, data[11]),
		AddData(SK_Economic, data[12]),
		AddData(SK_Culture, data[13]),
		AddData(SK_Nobility, data[14]),
		AddData(SK_Worlds, data[15]),
		AddData(SK_RU, data[16]),
	)
	return sse
}
