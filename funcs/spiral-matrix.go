package funcs

import "strconv"

// log.Println(funcs.SpiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))
// log.Println(funcs.SpiralOrder([][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}}))
// log.Println(funcs.SpiralOrder([][]int{
// 	[]int{1, 2, 3, 4},
// 	[]int{5, 6, 7, 8},
// 	[]int{9, 10, 11, 12},
// 	[]int{13, 14, 15, 16},
// 	[]int{17, 18, 19, 20},
// 	[]int{21, 22, 23, 24},
// }))

func SpiralOrder(matrix [][]int) []int {
	top := 0
	left := 0
	right := len(matrix[0]) - 1
	bottom := len(matrix) - 1

	m := make(map[string]bool, right*bottom)

	var createSpiral func(x, y int, spiral []int) []int
	createSpiral = func(x, y int, spiral []int) []int {
		cell := matrix[y][x]
		spiral = append(spiral, cell)

		strX := strconv.Itoa(x)
		strY := strconv.Itoa(y)
		coord := strX + strY
		m[coord] = true

		if len(spiral) == len(matrix)*len(matrix[0]) {
			return spiral
		}

		if (x == left || hasCheckedNeighbour(x-1, y, m)) && (y == top || hasCheckedNeighbour(x, y-1, m)) && x != right && !hasCheckedNeighbour(x+1, y, m) {
			x++
		} else if (x == right || hasCheckedNeighbour(x+1, y, m)) && (y == top || hasCheckedNeighbour(x, y-1, m)) && y != bottom && !hasCheckedNeighbour(x, y+1, m) {
			y++
		} else if (x == right || hasCheckedNeighbour(x+1, y, m)) && (y == bottom || hasCheckedNeighbour(x, y+1, m)) && x != left && !hasCheckedNeighbour(x-1, y, m) {
			x--
		} else if (x == left || hasCheckedNeighbour(x-1, y, m)) && (y == bottom || hasCheckedNeighbour(x, y+1, m)) && y != top && !hasCheckedNeighbour(x, y-1, m) {
			y--
		}

		return createSpiral(x, y, spiral)
	}

	return createSpiral(0, 0, []int{})
}

func hasCheckedNeighbour(x, y int, m map[string]bool) bool {
	strX := strconv.Itoa(x)
	strY := strconv.Itoa(y)
	coord := strX + strY
	_, ok := m[coord]

	return ok
}
