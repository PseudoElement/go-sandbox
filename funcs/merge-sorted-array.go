package funcs

import "sort"

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	var filteredNums1 []int
	for i, val := range nums1 {
		if i <= (m - 1) {
			filteredNums1 = append(filteredNums1, val)
		}
	}

	nums1 = append(filteredNums1, nums2...)

	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})

	return nums1
}
