package funcs

import (
	"math"
	"strconv"
)

func MajorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	m := make(map[string]int)
	for _, val := range nums {
		strVal := strconv.Itoa(val)
		_, ok := m[strVal]
		if !ok {
			m[strVal] = 0
		}
		m[strVal] += 1
	}

	var maxValue float64
	var maxKey string
	for key, val := range m {
		max := math.Max(maxValue, float64(val))
		maxValue = max
		if max == float64(val) {
			maxKey = key
		}
	}

	res, _ := strconv.Atoi(maxKey)

	return res
}
