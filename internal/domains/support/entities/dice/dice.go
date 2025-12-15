package roll

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"sync"
	"time"

	"github.com/Galdoba/t5/internal/domains/support/valueobjects/rollresult"
)

const (
	diceEdges           = 6
	even1to9SpecialCase = 5
	minD66Value         = 0
	maxD66Value         = 9
)

type roller struct {
	mu   sync.Mutex
	seed string
	rng  *rand.Rand
	// rollResults []int
}

var defaultRoller *roller

func init() {
	defaultRoller = &roller{
		mu:   sync.Mutex{},
		seed: "",
		rng:  rand.New(rand.NewSource(time.Now().UnixNano())),
		// rollResults: []int{},
	}
}

func New(seed fmt.Stringer) *roller {
	dp := roller{}
	dp.seed = seed.String()
	dp.rng = rand.New(rand.NewSource(stringToInt64(seed.String())))
	return &dp
}

func (dp *roller) Result(n int) rollresult.RollResult {
	dp.mu.Lock()
	defer dp.mu.Unlock()

	return rollresult.New(dp.roll(n)...)
}

func (dp *roller) Seed() string {
	return dp.seed
}

func bounded(value, min, max int) int {
	if min > max {
		panic("bounded: min > max")
	}
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// D66 return code by D66 rules. Expect 0-2 mods for first and second die respectivly
func (dp *roller) D66(mods ...int) string {
	dp.mu.Lock()
	defer dp.mu.Unlock()
	r1 := dp.rng.Intn(diceEdges) + 1
	r2 := dp.rng.Intn(diceEdges) + 1
	for len(mods) < 2 {
		mods = append(mods, 0)
	}
	r1 = bounded(r1+mods[0], minD66Value, maxD66Value)
	r2 = bounded(r2+mods[1], minD66Value, maxD66Value)
	return fmt.Sprintf("%v%v", r1, r2)
}

func (dp *roller) D3(mods ...int) int {
	dp.mu.Lock()
	defer dp.mu.Unlock()
	r := dp.rng.Intn(3) + 1
	for _, mod := range mods {
		r += mod
	}
	return r
}

func (dp *roller) Flux(mods ...int) int {
	dp.mu.Lock()
	defer dp.mu.Unlock()

	r1 := dp.rng.Intn(diceEdges) + 1
	r2 := dp.rng.Intn(diceEdges) + 1
	fl := r1 - r2
	for _, mod := range mods {
		fl += mod
	}
	return fl

}

func (dp *roller) GoodFlux(mods ...int) int {
	dp.mu.Lock()
	defer dp.mu.Unlock()

	r1 := dp.rng.Intn(diceEdges) + 1
	r2 := dp.rng.Intn(diceEdges) + 1
	fl := goodFlux(r1, r2)
	for _, mod := range mods {
		fl += mod
	}
	return fl
}

func (dp *roller) BadFlux(mods ...int) int {
	dp.mu.Lock()
	defer dp.mu.Unlock()

	r1 := dp.rng.Intn(diceEdges) + 1
	r2 := dp.rng.Intn(diceEdges) + 1
	fl := badFlux(r1, r2)
	for _, mod := range mods {
		fl += mod
	}
	return fl
}

func (dp *roller) Even0To9() int {
	dp.mu.Lock()
	defer dp.mu.Unlock()
	return dp.rng.Intn(10)
}

func (dp *roller) Even1To9() int {
	dp.mu.Lock()
	defer dp.mu.Unlock()
	return dp.rng.Intn(9) + 1
}

func (dp *roller) Percentile() int {
	dp.mu.Lock()
	defer dp.mu.Unlock()
	return dp.rng.Intn(100)
}

func (dp *roller) Variation() float64 {
	dp.mu.Lock()
	defer dp.mu.Unlock()
	return float64(dp.rng.Intn(100)) / 100
}

func stringToInt64(s string) int64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return int64(h.Sum64())
}

func (dp *roller) roll(n int) []int {
	rollResults := make([]int, n)
	if n <= 0 {
		return []int{}
	}
	for i := range n {
		rollResults[i] = dp.rng.Intn(diceEdges) + 1
	}
	return rollResults
}

func Result(n int) rollresult.RollResult {
	return defaultRoller.Result(n)
}

func D66(mods ...int) string {
	return defaultRoller.D66(mods...)
}

func Flux(mods ...int) int {
	return defaultRoller.Flux(mods...)
}

func GoodFlux(mods ...int) int {
	return defaultRoller.GoodFlux(mods...)
}

func BadFlux(mods ...int) int {
	return defaultRoller.BadFlux(mods...)
}

func Even1To9() int {
	return defaultRoller.Even1To9()
}

func Even0To9() int {
	return defaultRoller.Even0To9()
}

func Percentile() int {
	return defaultRoller.Percentile()
}

func Variation() float64 {
	return defaultRoller.Variation()
}

func goodFlux(r1, r2 int) int {
	max := max(r1, r2)
	min := min(r1, r2)
	return max - min
}

func badFlux(r1, r2 int) int {
	max := max(r1, r2)
	min := min(r1, r2)
	return min - max
}
