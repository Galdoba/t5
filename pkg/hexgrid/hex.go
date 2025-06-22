// hex.go
package hexgrid

// Hex - неизменяемая структура гекса
type Hex struct {
	q, r, s int // кубические координаты
}

// NewHex создает гекс с указанными координатами
func NewHex(q, r, s int) Hex {
	h := Hex{
		q: q,
		r: r,
		s: s,
	}
	return h
}

// Геттеры для координат
func (h Hex) Q() int { return h.q }
func (h Hex) R() int { return h.r }
func (h Hex) S() int { return h.s }
