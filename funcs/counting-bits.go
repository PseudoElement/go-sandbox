package funcs

func CountBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}

	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1

	for i := 2; i < n+1; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i/2] + 1
		}
	}

	return res
}
