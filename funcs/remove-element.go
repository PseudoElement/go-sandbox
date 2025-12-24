package funcs

import "log"

func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		} else {
			return 1
		}
	}

	start, end := 0, len(nums)-1
	for start <= end {
		for nums[start] == val && start <= end {
			nums[start], nums[end] = nums[end], nums[start]
			end--
		}
		start++
	}

	var count int
	for count < len(nums) {
		if nums[count] == val {
			break
		} else {
			count++
		}
	}

	log.Printf("arr - %+v, k - %d\n\n", nums, count)

	return count
}
