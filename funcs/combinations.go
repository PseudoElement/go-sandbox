package funcs

func Combine(n int, k int) [][]int {
	var ans [][]int
	var solve func(start int, temp []int)

	solve = func(start int, temp []int) {
		if len(temp) == k {
			comb := make([]int, len(temp))
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		for i := start; i <= n; i++ {
			solve(i+1, append(temp, i))
		}
	}

	solve(1, []int{})
	return ans
}
