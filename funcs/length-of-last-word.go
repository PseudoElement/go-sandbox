package funcs

import "strings"

func lengthOfLastWord(s string) int {
	trimmed := strings.TrimSpace(s)
	splitted := strings.Split(trimmed, " ")
	lastWord := splitted[len(splitted)-1]

	return len(lastWord)
}

func lengthOfLastWordFast(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		strChar := string(s[i])
		if strChar == " " {
			if length == 0 {
				continue
			} else {
				return length
			}
		} else {
			length++
		}
	}

	return length
}
