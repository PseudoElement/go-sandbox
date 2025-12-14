package funcs

// it checks common substring from any index
func LongestCommonSubstring(strs []string) string {
	longest := ""
	commonLetters := Set{}

	for _, str := range strs {
		if len(str) == 0 {
			return ""
		}
	}

MainLoop:
	for letterIdx, letter := range strs[0] {
		for _, str := range strs {
			if letterIdx > len(str)-1 {
				break MainLoop
			}
			if string(str[letterIdx]) == string(letter) {
				commonLetters.Add(string(letter))
			} else {
				for commonLetters.Size() > 0 {
					commonLetters.Pop()
				}
				continue MainLoop
			}
		}

		if len(longest) < commonLetters.Size() {
			longest = ""
			commonLetters.ForEach(func(idx int, el string) {
				longest += el
			})
		}
	}

	return longest
}

// it check common prefix only from 0 index
func longestCommonPrefix(strs []string) string {
	longest := ""

	for _, str := range strs {
		if len(str) == 0 {
			return ""
		}
	}

MainLoop:
	for letterIdx, letter := range strs[0] {
		for _, str := range strs {
			if letterIdx > len(str)-1 {
				break MainLoop
			}
			if string(str[letterIdx]) != string(letter) {
				break MainLoop
			}
		}

		longest += string(letter)
	}

	return longest
}
