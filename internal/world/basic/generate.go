package basic

import (
	"fmt"

	"github.com/Galdoba/t5/pkg/dice"
	"github.com/Galdoba/t5/pkg/ehex"
)

var generationStep int

const (
	stepCS = iota
	stepSA
)

func (w *World) Generate(dp *dice.Dicepool) error {
	if err := w.generateSize(dp); err != nil {
		fmt.Println(err)
	}
	// w.Size = ehex.FromInt(dp.Sum("2d6-2"), "auto generated")
	// if w.Size.Value() == 10 {
	// 	w.Size = ehex.FromInt(dp.Sum("1d6+9"), "auto generated")
	// }
	// w.Atmo = ehex.FromInt(dp.Sum("flux", dice.DM(w.Size.Value())))
	// if w.Size.Value() == 0 || w.Atmo.Value() < 0 {
	// 	w.Atmo = ehex.FromInt(0)
	// }
	// if w.Atmo.Value() > 15 {
	// 	w.Atmo = ehex.FromInt(15)
	// }
	// hDM := w.Atmo.Value()
	// if w.Atmo.Value() < 2 || w.Atmo.Value() > 9 {
	// 	hDM += -4
	// }
	// w.Hydr = ehex.FromInt(dp.Sum("flux", dice.DM(hDM)))
	// if w.Size.Value() < 2 || w.Hydr.Value() < 0 {
	// 	w.Hydr = ehex.FromInt(0)
	// }
	// if w.Hydr.Value() > 10 {
	// 	w.Atmo = ehex.FromInt(10)
	// }
	// //w.Pops = ehex.FromInt(dp.Sum("2d6-2"))
	return nil
}

func (w *World) generateSize(dp *dice.Dicepool) error {
	if w.Size != nil {
		return fmt.Errorf("size was generated or injected")
	}
	if w.Template == "" {

	}
	switch w.Template {
	case "garden":
		w.Size = ehex.FromInt(dp.Sum("2d6-2"), "auto generated")
		if w.Size.Value() == 10 {
			w.Size = ehex.FromInt(dp.Sum("1d6+9"), "auto generated")
		}
	}
	return nil
}

/*

 */
