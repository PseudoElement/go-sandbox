package funcs

import "strings"

func NumDecodings(s string) int {
	if len(s) < 3 && strings.HasPrefix(s, "0") {
		return 0
	}

	dict := map[string]string{
		"1":  "A",
		"2":  "B",
		"3":  "C",
		"4":  "D",
		"5":  "E",
		"6":  "F",
		"7":  "G",
		"8":  "H",
		"9":  "I",
		"10": "J",
		"11": "K",
		"12": "L",
		"13": "M",
		"14": "N",
		"15": "O",
		"16": "P",
		"17": "Q",
		"18": "R",
		"19": "S",
		"20": "T",
		"21": "U",
		"22": "V",
		"23": "W",
		"24": "X",
		"25": "Y",
		"26": "Z",
	}

	// 1110   -> [1 1 10] [11 10]
	// 11106  -> [11 10 6] [1 1 10 6]
	// 111062 -> [11 10 6 2] [1 1 10 6 2]
	// 111060

	var dp func(rest string) int
	dp = func(rest string) int {
		if len(rest) <= 2 {
			if strings.Contains(rest, "0") {
				return 1
			} else {
				return 2
			}
		}
		return 0
	}

	return 1
}
