package starsystem

import (
	"fmt"
	"testing"
)

func Test_closest(t *testing.T) {
	bt, err := OrbitNConversionCache()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bt))
	for i := 0; i <= 2000; i += 10 {
		o := float64(i) / 100
		a := Convert_OrbitN_AU(o)
		//fmt.Printf("%v\t-> %v\n", o, a)
		b := Convert_AU_OrbitN(a)
		//fmt.Printf("%v\t-> %v\n", a, b)

		fmt.Printf("%v\t-> %v\t%v", o, a, b)
		switch o == b {
		case true:
			fmt.Printf("\tmatch!!\n")
		case false:
			fmt.Printf("\tBAD!!\n")
		}

	}

}
