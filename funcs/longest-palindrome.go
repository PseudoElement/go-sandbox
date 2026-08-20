package funcs

/*
*
*
Example 1:
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Example 2:
Input: s = "cbbd"
Output: "bb"
*/
func LongestPalindrome(s string) string {
	palindrome := ""
	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			str := s[i:j]
			if IsPalindrom(str) && len(str) > len(palindrome) {
				palindrome = str
			}
		}
	}

	return palindrome
}

func IsPalindrom(str string) bool {
	ptrStart := 0
	ptrEnd := len(str) - 1
	for ptrStart < ptrEnd {
		if str[ptrStart] != str[ptrEnd] {
			return false
		}
		ptrStart++
		ptrEnd--
	}
	return true
}
