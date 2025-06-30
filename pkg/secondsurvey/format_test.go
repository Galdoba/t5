package secondsurvey

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		report string
	}
	tests := []struct {
		name    string
		args    args
		want    SecondSurveyReport
		wantErr bool
	}{
		{
			name: "Regina",
			args: args{
				report: "1910	Regina	A788899-C	Ri Pa Ph An Cp	{+4}	(D7E+4)	[9C6D]	BcCeF	NS	-	703	8	Im	F7 V DM M3 V",
			},
			want: SecondSurveyReport{
				Hex:        "1910",
				Name:       "Regina",
				UWP:        "A788899-C",
				Remarks:    []string{"Ri", "Pa", "Ph", "An", "Cp"},
				Importance: "{+4}",
				Economic:   "(D7E+4)",
				Culture:    "[9C6D]",
				Nobility:   []string{"B", "c", "C", "e", "F"},
				Bases:      []string{"N", "S"},
				Zone:       "-",
				PBG:        "703",
				Worlds:     "8",
				Allegiance: "Im",
				Stellar:    "F7 V DM M3 V",
			},
			wantErr: false,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.report)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
			fmt.Println(got.Format())
		})
	}
}
