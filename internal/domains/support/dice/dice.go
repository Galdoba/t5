package dice

import (
	"hash/fnv"
	"math/rand"
	"sync"
	"time"
)

const (
	diceEdges = 6
)

// Pool использует пул для переиспользования срезов
var slicePool = sync.Pool{
	New: func() interface{} {
		return make([]int, 0, 12) // начальная емкость 12 (оптимизация для частых бросков)
	},
}

type Pool struct {
	rng        *rand.Rand
	seed       string
	lastRolled []int
}

// newRng создает генератор случайных чисел
func newRng(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

// New создает новый пул кубиков с указанным сидом
func New(seed string) *Pool {
	if seed == "" {
		seed = time.Now().String()
	}

	dp := &Pool{
		seed: seed,
		rng:  newRng(stringToInt64(seed)),
	}

	return dp
}

// stringToInt64 конвертирует строку в int64 используя FNV-1a
func stringToInt64(s string) int64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return int64(h.Sum64())
}

// Roll выполняет бросок n кубиков и возвращает сумму
func (dp *Pool) Roll(n int) int {
	if n <= 0 {
		dp.lastRolled = nil
		return 0
	}

	// Получаем срез из пула и настраиваем его размер
	rolled := slicePool.Get().([]int)
	if cap(rolled) < n {
		rolled = make([]int, n)
	} else {
		rolled = rolled[:n]
	}

	// Заполняем случайными значениями и вычисляем сумму
	sum := 0
	for i := 0; i < n; i++ {
		val := dp.rng.Intn(diceEdges) + 1
		rolled[i] = val
		sum += val
	}

	// Возвращаем старый срез в пул
	if dp.lastRolled != nil {
		slicePool.Put(dp.lastRolled[:0])
	}

	dp.lastRolled = rolled
	return sum
}

// Results возвращает копию результатов последнего броска
func (dp *Pool) Results() []int {
	if dp.lastRolled == nil {
		return nil
	}

	// Возвращаем копию для безопасности
	result := make([]int, len(dp.lastRolled))
	copy(result, dp.lastRolled)
	return result
}

// Roll выполняет единичный бросок n кубиков с новым генератором
func Roll(n int) int {
	if n <= 0 {
		return 0
	}

	// Используем быстрый способ без создания полного пула
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	sum := 0

	for i := 0; i < n; i++ {
		sum += rng.Intn(diceEdges) + 1
	}

	return sum
}
