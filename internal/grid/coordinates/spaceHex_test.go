package coordinates

import (
	"fmt"
	"testing"
)

func Test_func(t *testing.T) {
	// Тестовые точки
	tests := []struct {
		q, r, s int
		x, y    int
	}{
		{0, -39, 39, 1, 1},     // Верхний-левый
		{0, 0, 0, 1, 40},       // нижний-левый
		{2, -3, 1, 3, 38},      // нижний-левый + немного отойти
		{31, -15, -16, 32, 40}, // нижний-правый
		{31, -54, 23, 32, 1},   // нижний-правый
	}

	for i, test := range tests {
		fmt.Println("test", i)
		// Проверка hex → square
		x, y := hex_to_square(test.q, test.r, test.s)
		if x != test.x || y != test.y {
			println("HEX->SQUARE ERROR:", test.q, test.r, test.s, "expected", test.x, test.y, "got", x, y)
		}

		// Проверка square → hex
		q, r, s := square_to_hex(test.x, test.y)
		if q != test.q || r != test.r || s != test.s {
			println("SQUARE->HEX ERROR:", test.x, test.y, "expected", test.q, test.r, test.s, "got", q, r, s)
		}
	}
}

// func TestGetSectorAndLocalCoords(t *testing.T) {
// 	tests := []struct {
// 		name string // description of this test case
// 		// Named input parameters for target function.
// 		global_q int
// 		global_r int
// 		global_s int
// 		want     int
// 		want2    int
// 		want3    int
// 		want4    int
// 		want5    int
// 		want6    int
// 	}{
// 		{
// 			name:     "center",
// 			global_q: 0,
// 			global_r: 0,
// 			global_s: 0,
// 			want:     0,
// 			want2:    0,
// 			want3:    1,
// 			want4:    40,
// 			want5:    0,
// 			want6:    0,
// 		}, // TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got2, got3, got4, got5, got6 := GetSectorAndLocalCoords(tt.global_q, tt.global_r, tt.global_s)
// 			// TODO: update the condition below to compare got with tt.want.
// 			if true {
// 				t.Errorf("GetSectorAndLocalCoords() = %v, want %v", got, tt.want)
// 			}
// 			if true {
// 				t.Errorf("GetSectorAndLocalCoords() = %v, want %v", got2, tt.want2)
// 			}
// 			if true {
// 				t.Errorf("GetSectorAndLocalCoords() = %v, want %v", got3, tt.want3)
// 			}
// 			if true {
// 				t.Errorf("GetSectorAndLocalCoords() = %v, want %v", got4, tt.want4)
// 			}
// 			if true {
// 				t.Errorf("GetSectorAndLocalCoords() = %v, want %v", got5, tt.want5)
// 			}
// 			if true {
// 				t.Errorf("GetSectorAndLocalCoords() = %v, want %v", got6, tt.want6)
// 			}
// 		})
// 	}
// }
