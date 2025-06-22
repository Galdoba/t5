package coordinates

func hex_to_square(q, r, s int) (int, int) {
	// Проверка кубического условия (должно выполняться q + r + s == 0)
	if q+r+s != 0 {
		panic("Invalid cubic coordinates: q + r + s must be 0")
	}

	// Отмена вертикального сдвига
	r_orig := r + 39
	s_orig := s - 39

	// Проверка восстановленного условия
	if q+r_orig+s_orig != 0 {
		panic("Shift reversal failed: inconsistent coordinates")
	}

	// x = q + 1 (смещение столбца)
	x := q + 1

	// Вычисление y с учетом смещения столбцов и восстановленного r
	// ((q) - (q & 1)) / 2 - коррекция строки для "кирпичной" кладки
	y := r_orig + ((q-(q&1))/2 + 1)

	return x, y
}

func square_to_hex(x, y int) (int, int, int) {
	// Проверка границ сетки
	if x < 1 || x > 32 || y < 1 || y > 40 {
		panic("Coordinates out of grid bounds")
	}

	// q = x - 1 (индексация с 0)
	q := x - 1

	// Вычисление r с учетом смещения столбцов
	r_orig := (y - 1) - ((q - (q & 1)) / 2)

	// s вычисляется из кубического условия
	s_orig := -q - r_orig

	// Применяем вертикальный сдвиг: перемещаем начало координат в нижний левый угол
	r := r_orig - 39
	s := s_orig + 39

	return q, r, s
}
