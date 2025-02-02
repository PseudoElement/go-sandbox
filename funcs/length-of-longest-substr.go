package funcs

import "math"

func lengthOfLongestSubstring(s string) int {
	left := 0
	maxLen := 0
	set := NewSet()

	for right, char := range s {
		for set.Has(string(char)) {
			leftChar := string(s[left])
			set.Remove(leftChar)
			left++
		}
		set.Add(string(char))

		maxLen = int(math.Max(float64(maxLen), float64(right-left+1)))
	}

	return maxLen
}
