package hexgrid

import (
	"math"
)

const (
	north     = 0
	northEast = 1
	southEast = 2
	south     = 3
	southWest = 4
	northWest = 5
)

// Distance возвращает расстояние между двумя гексами в шагах
func Distance(a, b Hex) int {
	return (abs(a.q-b.q) + abs(a.r-b.r) + abs(a.s-b.s)) / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var directions = []Hex{
	{q: 0, r: -1, s: 1}, // 0: north
	{q: 1, r: -1, s: 0}, // 1: northEast
	{q: 1, r: 0, s: -1}, // 2: southEast
	{q: 0, r: 1, s: -1}, // 3: south
	{q: -1, r: 1, s: 0}, // 4: southWest
	{q: -1, r: 0, s: 1}, // 5: northWest
}

// Neighbors возвращает 6 соседних гексов в порядке направлений
func Neighbors(h Hex) []Hex {
	neighbors := make([]Hex, 0, 6)
	for _, d := range directions {
		neighbors = append(neighbors, Hex{
			q: h.q + d.q,
			r: h.r + d.r,
			s: h.s + d.s,
		})
	}
	return neighbors
}

// LineDrawing возвращает гексы на прямой линии между двумя точками
func LineDrawing(a, b Hex) []Hex {
	N := Distance(a, b)
	results := make([]Hex, 0, N+1)

	// Добавляем начальную точку
	results = append(results, a)

	// Линейная интерполяция для N сегментов
	for i := 1; i <= N; i++ {
		t := float64(i) / float64(N)
		q := lerp(float64(a.q), float64(b.q), t)
		r := lerp(float64(a.r), float64(b.r), t)
		s := lerp(float64(a.s), float64(b.s), t)

		// Округляем до ближайшего гекса
		hex := roundHex(q, r, s)
		results = append(results, hex)
	}

	return results
}

func lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

func roundHex(q, r, s float64) Hex {
	rq := math.Round(q)
	rr := math.Round(r)
	rs := math.Round(s)

	// Коррекция округления для соблюдения q + r + s = 0
	dq := math.Abs(rq - q)
	dr := math.Abs(rr - r)
	ds := math.Abs(rs - s)

	if dq > dr && dq > ds {
		rq = -rr - rs
	} else if dr > ds {
		rr = -rq - rs
	} else {
		rs = -rq - rr
	}

	return Hex{
		q: int(rq),
		r: int(rr),
		s: int(rs),
	}
}

// Ring возвращает гексы на заданном расстоянии от центра
func Ring(center Hex, radius int) []Hex {
	if radius < 0 {
		return []Hex{}
	}
	if radius == 0 {
		return []Hex{center}
	}

	// Начинаем с направления north и двигаемся на radius шагов
	// после чего выбираем следующее направление
	ring := make([]Hex, 0, 6*radius)
	current := move(center, 0, radius)
	directionOptimization := 2
	for direction := range 6 {
		for range radius {
			ring = append(ring, current)
			current = neighbor(current, (direction+directionOptimization)%6)
		}
	}
	return ring
}

// Вспомогательная функция для перемещения на несколько шагов
func move(h Hex, direction, steps int) Hex {
	current := h
	for range steps {
		current = neighbor(current, direction)
	}
	return current
}

// Получение соседа в определенном направлении
func neighbor(h Hex, direction int) Hex {
	d := directions[direction]
	return Hex{
		q: h.q + d.q,
		r: h.r + d.r,
		s: h.s + d.s,
	}
}

// Spiral возвращает все гексы в пределах заданного радиуса
func Spiral(center Hex, radius int) []Hex {
	results := []Hex{center}

	for k := 1; k <= radius; k++ {
		ring := Ring(center, k)
		results = append(results, ring...)
	}

	return results
}

// SpiralMaps возвращает две карты для спирального обхода
func SpiralMaps(center Hex, radius int) (map[int]Hex, map[Hex]int) {
	indexToHex := make(map[int]Hex)
	hexToIndex := make(map[Hex]int)
	if radius < 0 {
		return indexToHex, hexToIndex
	}

	// Центральный гекс
	counter := 0
	indexToHex[counter] = center
	hexToIndex[center] = counter
	counter++

	// Обходим кольца от 1 до заданного радиуса
	for r := 1; r <= radius; r++ {
		ringHexes := Ring(center, r)
		for _, hex := range ringHexes {
			indexToHex[counter] = hex
			hexToIndex[hex] = counter
			counter++
		}
	}

	return indexToHex, hexToIndex
}

// Vector выполняет покоординатное сложение двух гексов
func Vector(start, change Hex) Hex {
	return Hex{
		q: start.q + change.q,
		r: start.r + change.r,
		s: start.s + change.s,
	}
}

func Rotate(start, center Hex, steps int) Hex {
	// Если start совпадает с центром или нулевые шаги
	if start == center || steps == 0 {
		return start
	}

	// Вычисляем радиус кольца
	r := Distance(center, start)

	// Получаем кольцо с нужным радиусом
	ring := Ring(center, r)

	// Находим позицию start в кольце
	pos := -1
	for i, hex := range ring {
		if hex == start {
			pos = i
			break
		}
	}

	// Если точка не найдена в кольце
	if pos == -1 {
		return start
	}

	// Вычисляем новую позицию с учетом направления
	n := len(ring)
	if steps > 0 {
		// Вращение по часовой стрелке
		return ring[(pos+steps)%n]
	} else {
		// Вращение против часовой стрелки
		newPos := (pos + steps) % n
		if newPos < 0 {
			newPos += n
		}
		return ring[newPos]
	}
}
