package secondsurvey

import (
	"fmt"
	"regexp"
	"strings"
)

//              Hex	Name	UWP	Rem	{ Ix }	( Ex )	[ Cx ]	N	B	Z	PBG	W	A	Stellar

const (
	t5ss = 14
	tmss = 17
)

type SecondSurveyReport struct {
	Hex        string
	Name       string
	UWP        string
	Remarks    []string
	Importance string
	Economic   string
	Culture    string
	Nobility   []string
	Bases      []string
	Zone       string
	PBG        string
	Worlds     string
	Allegiance string
	Stellar    string
}

func (ssr SecondSurveyReport) Format() string {
	s := ""
	s += fmt.Sprintf("%s\t", ssr.Hex)
	s += fmt.Sprintf("%s\t", ssr.Name)
	s += fmt.Sprintf("%s\t", ssr.UWP)
	s += fmt.Sprintf("%s\t", strings.Join(ssr.Remarks, " "))
	s += fmt.Sprintf("%s\t", ssr.Importance)
	s += fmt.Sprintf("%s\t", ssr.Economic)
	s += fmt.Sprintf("%s\t", ssr.Culture)
	s += fmt.Sprintf("%s\t", strings.Join(ssr.Nobility, ""))
	s += fmt.Sprintf("%s\t", strings.Join(ssr.Bases, ""))
	s += fmt.Sprintf("%s\t", ssr.Zone)
	s += fmt.Sprintf("%s\t", ssr.PBG)
	s += fmt.Sprintf("%s\t", ssr.Worlds)
	s += fmt.Sprintf("%s\t", ssr.Allegiance)
	s += fmt.Sprintf("%s", ssr.Stellar)
	return s
}

func Parse(report string) (SecondSurveyReport, error) {
	ssr := SecondSurveyReport{}
	dataSet := strings.Split(report, "\t")
	switch len(dataSet) {
	default:
		return ssr, fmt.Errorf("unknown data format")
	case t5ss:
		return parse_t5ss(report)
	case tmss:
		return parse_tmss(report)
	}

	return ssr, nil
}

func parse_t5ss(str string) (SecondSurveyReport, error) {
	ssr := SecondSurveyReport{}
	dataSet := strings.Split(str, "\t")
	for i, data := range dataSet {
		err := fmt.Errorf("dataset[%v] not parsed", i)
		switch i {
		case 0:
			ssr.Hex, err = parseExpression(data, hex)
		case 1:
			ssr.Name, err = parseExpression(data, anyText)
		case 2:
			ssr.UWP, err = parseExpression(data, uwp)
		case 3:
			ssr.Remarks, err = parseSlice(data, " ")
		case 4:
			ssr.Importance, err = parseExpression(data, ix)
		case 5:
			ssr.Economic, err = parseExpression(data, ex)
		case 6:
			ssr.Culture, err = parseExpression(data, cx)
		case 7:
			ssr.Nobility, err = parseSlice(data, "")
		case 8:
			ssr.Bases, err = parseSlice(data, "")
		case 9:
			ssr.Zone, err = parseExpression(data, zone)
		case 10:
			ssr.PBG, err = parseExpression(data, pbg)
		case 11:
			ssr.Worlds, err = parseExpression(data, worlds)
		case 12:
			ssr.Allegiance, err = parseExpression(data, anyText)
		case 13:
			ssr.Stellar, err = parseExpression(data, anyText)
		}
		if err != nil {
			return ssr, fmt.Errorf("parse dataset[%v]: %v", i, err)
		}

	}
	return ssr, nil
}

// 0         1   2   3       4   5       6       7       8   9           10      11      12      13      14          15  16
// Sector	SS	Hex	Name	UWP	Bases	Remarks	Zone	PBG	Allegiance	Stars	{Ix}	(Ex)	[Cx]	Nobility	W	RU
func parse_tmss(str string) (SecondSurveyReport, error) {
	ssr := SecondSurveyReport{}
	dataSet := strings.Split(str, "\t")
	for i, data := range dataSet {
		err := fmt.Errorf("dataset[%v] not parsed", i)
		switch i {
		case 2:
			ssr.Hex, err = parseExpression(data, hex)
		case 3:
			ssr.Name, err = parseExpression(data, anyText)
		case 4:
			ssr.UWP, err = parseExpression(data, uwp)
		case 6:
			ssr.Remarks, err = parseSlice(data, " ")
		case 11:
			ssr.Importance, err = parseExpression(data, ix)
		case 12:
			ssr.Economic, err = parseExpression(data, ex)
		case 13:
			ssr.Culture, err = parseExpression(data, cx)
		case 14:
			ssr.Nobility, err = parseSlice(data, "")
		case 5:
			ssr.Bases, err = parseSlice(data, "")
		case 7:
			ssr.Zone, err = parseExpression(data, zone)
		case 8:
			ssr.PBG, err = parseExpression(data, pbg)
		case 15:
			ssr.Worlds, err = parseExpression(data, worlds)
		case 9:
			ssr.Allegiance, err = parseExpression(data, anyText)
		case 10:
			ssr.Stellar, err = parseExpression(data, anyText)
		}
		if err != nil {
			return ssr, fmt.Errorf("parse dataset[%v]: %v", i, err)
		}

	}
	return ssr, nil
}

func parseSlice(s, separator string) ([]string, error) {
	return strings.Split(s, separator), nil
}

var hex = `^(\d){3,4}$`
var uwp = `^[0-Z]{7}-[0-Z]$`
var ix = `^\{[+|-]?[0-9]\}$`
var ex = `^\([0-Z]{3}[+|-][0-Z]\)$`
var cx = `^\[[0-Z]{4}\]$`
var zone = `^[-|A|R]$`
var pbg = `^[1-9][0-3][0-6]$`
var worlds = `^[0-9]+$`
var anyText = `^.*$`

func parseExpression(s, exp string) (string, error) {
	if s == "" {
		return "", nil
	}
	re := regexp.MustCompile(fmt.Sprintf("%v", exp))
	found := re.FindString(s)
	if found == "" {
		return "", fmt.Errorf("failed to parse segment (%v)", s)
	}
	return found, nil

}
