// hex.go
package cube

// Hex - неизменяемая структура гекса
type Hex struct {
	Q, R, S int // кубические координаты
}

// NewHex создает гекс с указанными координатами
func NewCube(q, r, s int) Hex {
	h := Hex{
		Q: q,
		R: r,
		S: s,
	}
	return h
}
