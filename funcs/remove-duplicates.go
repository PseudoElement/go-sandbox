package funcs

func RemoveDuplicates(nums []int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); {
		num := nums[i]
		_, ok := m[num]
		if !ok {
			m[num] = 1
		} else {
			m[num] = m[num] + 1
		}

		if m[num] > 2 {
			nums = append(nums[0:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return nums
}
