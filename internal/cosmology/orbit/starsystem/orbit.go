package starsystem

import (
	"encoding/json"
	"math"
)

type Orbit struct {
	OrbitN         int     //0 - 2000 (1= 10 )
	AU             int     //0 - 787000 (1= 150 000 000 km)
	Retrograde     bool    //if true move conterclockwise
	Essentrisity   float64 //0 - 1 for min-max calculation
	AzimuthOffset  int     //0 - 3600
	Azimuth        int     //0 - 3600
	PeriodDuration int     //in space tiks (6 min)
}

const (
	au_min = 20
	au_max = 7870000
)

var xs = [21]int{0, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400, 1500, 1600, 1700, 1800, 1900, 2000}
var ys = [21]int{1, 40, 70, 100, 160, 280, 520, 1000, 2000, 4000, 7700, 15400, 30800, 61500, 123000, 250000, 490000, 980000, 1950000, 3950000, 7870000}
var logYs [21]float64

func init() {
	for i, y := range ys {
		logYs[i] = math.Log(float64(y))
	}
}

func convertAtoB(a int) int {
	n := len(xs)
	if a <= xs[0] {
		return ys[0]
	}
	if a >= xs[n-1] {
		return ys[n-1]
	}

	for i := 0; i < n-1; i++ {
		if a >= xs[i] && a <= xs[i+1] {
			if a == xs[i] {
				return ys[i]
			}
			if a == xs[i+1] {
				return ys[i+1]
			}
			fraction := float64(a-xs[i]) / float64(xs[i+1]-xs[i])
			yLog := logYs[i] + fraction*(logYs[i+1]-logYs[i])
			return int(math.Round(math.Exp(yLog)))
		}
	}
	return 0
}

func convertBtoA(b int) int {

	first := 0
	last := len(cache.DataONtoAU) - 1

	// Граничные случаи
	if b <= cache.DataONtoAU[first] {
		return first
	}
	if b >= cache.DataONtoAU[last] {
		return last
	}

	// Бинарный поиск нижней границы
	low, high := first, last
	for high-low > 1 {
		mid := (low + high) / 2
		if cache.DataONtoAU[mid] <= b {
			low = mid
		} else {
			high = mid
		}
	}
	// Логарифмическая интерполяция
	// b1 := float64(cache.DataONtoAU[low])
	// b2 := float64(cache.DataONtoAU[high])
	// logB := math.Log(float64(b))
	// logB1 := math.Log(b1)
	// logB2 := math.Log(b2)

	// fraction := (logB - logB1) / (logB2 - logB1)
	// rounded := math.Round(float64(low) + fraction)
	// return int(rounded)
	return low
}

type OrbinNToAUConversion struct {
	DataONtoAU map[int]int `json:"Orbit# to AU conversion"`
	// ProblematicValuesAUtoON map[int]int `json:"Problematic conversion Values AU to Orbit#"`
}

var cache OrbinNToAUConversion

func OrbitNConversionCache() ([]byte, error) {
	c := OrbinNToAUConversion{}
	c.DataONtoAU = make(map[int]int)
	// c.ProblematicValuesAUtoON = make(map[int]int)
	for o := 0; o <= 2000; o++ {
		in := o
		out := convertAtoB(in)
		c.DataONtoAU[in] = out
		// back := convertBtoA(out)
		// if back != in {
		// 	fmt.Println("add", out, o, back)
		// 	c.ProblematicValuesAUtoON[out] = o
		// }
	}
	cache = c
	return json.MarshalIndent(&c, "", "  ")
}

func Convert_OrbitN_AU(on float64) float64 {
	// fmt.Println("---")
	on = limit(on, 0.0, 20.0)
	// fmt.Println("on", on)
	key := int(on * 100)
	// fmt.Println("key", key)
	out := cache.DataONtoAU[key]
	// fmt.Println("out", out, float64(out)/100)
	return float64(out) / 100
}

func Convert_AU_OrbitN(auFlt float64) float64 {
	low, high := -1.0, 80000.0
	for i := 0; i <= 200; i++ {
		try := float64(i) / 10
		catch := Convert_OrbitN_AU(try)
		if catch <= auFlt {
			low = catch
		}
		if catch >= auFlt {
			high = catch
		}
		if catch == auFlt {
			return try
		}
	}

	// // fmt.Println("+++")
	// auKey := int(auFlt * 100)
	// // fmt.Println("auKey", auFlt, auKey)
	// au := int(limit(float64(auKey), au_min, au_max))
	// // fmt.Println("au", au)
	// on := convertBtoA(au)
	// // fmt.Println("on", on, float64(on)/100)
	// return float64(on) / 100
	return (low + high) / 2
}

func limit(f, min, max float64) float64 {
	if f < min {
		return min
	}
	if f > max {
		return max
	}
	return f
}
