// grid.go
package hexgrid

// // Grid представляет сетку гексагонов с привязанными объектами
// type Grid struct {
// 	desc  string
// 	hexes map[Hex]any
// }

// // NewGrid создает новую пустую сетку
// func NewGrid(desc string) *Grid {
// 	return &Grid{
// 		desc:  desc,
// 		hexes: make(map[Hex]any),
// 	}
// }

// // Add добавляет объект в указанный гекс
// func (g *Grid) Add(h Hex, obj any) {
// 	g.hexes[h] = obj
// }

// // Get возвращает объект по указанному гексу
// func (g *Grid) Get(h Hex) (any, bool) {
// 	obj, exists := g.hexes[h]
// 	return obj, exists
// }

// // Remove удаляет объект из указанного гекса
// func (g *Grid) Remove(h Hex) {
// 	delete(g.hexes, h)
// }

// // Contains проверяет наличие гекса в сетке
// func (g *Grid) Contains(h Hex) bool {
// 	_, exists := g.hexes[h]
// 	return exists
// }

// // Objects возвращает все объекты сетки
// func (g *Grid) Objects() []any {
// 	result := make([]any, 0, len(g.hexes))
// 	for _, obj := range g.hexes {
// 		result = append(result, obj)
// 	}
// 	return result
// }

// // Hexes возвращает все гексы сетки
// func (g *Grid) Hexes() []Hex {
// 	keys := make([]Hex, 0, len(g.hexes))
// 	for k := range g.hexes {
// 		keys = append(keys, k)
// 	}
// 	return keys
// }

// // Desc возвращает описание сетки
// func (g *Grid) Desc() string { return g.desc }
