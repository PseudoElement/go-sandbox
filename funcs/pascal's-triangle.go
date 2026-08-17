package funcs

func PascalsTriangle(numRows int) [][]int {
	res := make([][]int, numRows)
	for idx := range numRows {
		res[idx] = make([]int, idx+1)
	}
	res[0][0] = 1

	for i := range numRows {
		if i == 0 {
			continue
		}

		for j, _ := range res[i] {
			itemValue := res[i][j]
			if inBoundLeft(j) && inBoundRight(j, res[i]) {
				itemValue = res[i-1][j-1] + res[i-1][j]
			} else if !inBoundLeft(j) {
				itemValue = res[i-1][j]
			} else if !inBoundRight(j, res[i]) {
				itemValue = res[i-1][j-1]
			}
			res[i][j] = itemValue
		}
	}

	return res
}

func inBoundLeft(itemIdx int) bool {
	return itemIdx != 0
}

func inBoundRight(itemIdx int, row []int) bool {
	return itemIdx != len(row)-1
}
