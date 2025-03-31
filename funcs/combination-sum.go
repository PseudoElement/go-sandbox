package funcs

func CombinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	dfs(0, 0, []int{}, &result, target, candidates)

	return result
}

func dfs(i, total int, arr []int, result *[][]int, target int, candidates []int) {
	if total == target {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*result = append(*result, tmp)
		return
	}
	if total > target || i == len(candidates) {
		return
	}
	dfs(i, total+candidates[i], append(arr, candidates[i]), result, target, candidates)
	dfs(i+1, total, arr, result, target, candidates)
}

func CombinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var cand []int

	var dfs func(int)

	dfs = func(i int) {
		total := sum(cand)
		if i > len(candidates)-1 {
			return
		}
		if total > target {
			return
		} else if total == target {
			temp := make([]int, len(cand))
			copy(temp, cand)
			res = append(res, temp)
			return
		}

		cand = append(cand, candidates[i])
		dfs(i)
		cand = cand[:len(cand)-1]
		dfs(i + 1)
	}
	dfs(0)
	return res
}
func sum(nums []int) int {
	res := 0
	for _, val := range nums {
		res += val
	}
	return res
}
