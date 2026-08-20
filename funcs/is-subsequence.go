package funcs

/**
* Example 1:
	Input: s = "abc", t = "ahbgdc"
	Output: true
  Example 2:
	Input: s = "axc", t = "ahbgdc"
	Output: false

	s = "aabc"
	t = "abaaaccba"
*/
// check if s is part of t
func IsSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	currIdxSubstr := 0
	for _, charUtf8 := range t {
		if currIdxSubstr > len(s)-1 {
			return true
		}
		if charUtf8 == rune(s[currIdxSubstr]) {
			currIdxSubstr++
		}
	}

	return currIdxSubstr == len(s)
}
