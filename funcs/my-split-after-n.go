package funcs

func MySplitAfterN(str string, sep string, n int) []string {
	var matchCount int
	var sub string
	var s = make([]string, 0, 10)
	for idx, char := range str {
		sub += string(char)
		if matchCount < n-1 {
			if string(char) == sep {
				matchCount++
				s = append(s, sub)
				sub = ""
			}
		} else {
			s = append(s, str[idx:])
			break
		}
	}

	return s
}
