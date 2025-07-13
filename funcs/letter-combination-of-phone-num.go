package funcs

import (
	"log"
	"time"
)

func LetterCombinations(digits string) []string {
	startTime := time.Now()

	lettersMap := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	if len(digits) == 0 {
		return []string{}
	}

	var result []string
	var backtrack func(int, string)

	backtrack = func(index int, path string) {
		if len(path) == len(digits) {
			result = append(result, path)
			return
		}

		// 2 or 3 or 4 ...
		currentDigit := string(digits[index])
		// {a b c} {d e f}
		lettersByDigit := lettersMap[currentDigit]
		for _, letter := range lettersByDigit {
			backtrack(index+1, path+letter)
		}
	}

	backtrack(0, "")
	// digit - 4, lettersByDigit - {"g", "h", "i"}

	log.Println("It took ", time.Since(startTime).Milliseconds(), "ms.")

	return result
}
