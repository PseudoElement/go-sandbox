package funcs

import (
	"log"
	"strconv"
)

type Combo struct {
	HoursNums   []int
	MinutesNums []int
}

// incorrect solution but good looking :)
func ReadBinaryWatch(turnedOn int) []string {
	hours := []int{1, 2, 4, 8}
	minutes := []int{1, 2, 4, 8, 16, 32}

	if turnedOn == 0 {
		return []string{"0:00"}
	}
	if turnedOn > 8 {
		return []string{}
	}

	combs := []Combo{}

	var find func(hoursIdx int, minsIdx int, combo Combo)
	find = func(hoursIdx int, minsIdx int, combo Combo) {
		if len(combo.HoursNums)+len(combo.MinutesNums) >= turnedOn || (hoursIdx >= len(hours) && minsIdx >= len(minutes)) {
			combs = append(combs, combo)
			return
		}

		if hoursIdx >= len(hours) {
			minChar := minutes[minsIdx]
			combo.MinutesNums = append(combo.MinutesNums, minChar)
		} else if minsIdx >= len(minutes) {
			hoursChar := hours[hoursIdx]
			combo.HoursNums = append(combo.HoursNums, hoursChar)
		} else {
			hoursChar := hours[hoursIdx]
			minChar := minutes[minsIdx]
			if turnedOn == 1 {
				copiedCombo := combo
				combo.HoursNums = append(combo.HoursNums, hoursChar)
				find(hoursIdx+1, minsIdx, combo)

				copiedCombo.MinutesNums = append(copiedCombo.MinutesNums, minChar)
				find(hoursIdx, minsIdx+1, copiedCombo)

				return
			}

			combo.HoursNums = append(combo.HoursNums, hoursChar)
			combo.MinutesNums = append(combo.MinutesNums, minChar)
		}

		find(hoursIdx+1, minsIdx+1, combo)
		find(hoursIdx, minsIdx+1, combo)
		find(hoursIdx+1, minsIdx, combo)
		find(hoursIdx, minsIdx, combo)
	}

	for hoursIdx, _ := range hours {
		find(hoursIdx, 0, Combo{})
	}
	for minsIdx, _ := range minutes {
		find(0, minsIdx, Combo{})
	}

	// find(0, 0, Combo{})
	log.Println("comb ==>", combs)

	resMap := make(map[string]bool, len(combs))
	results := make([]string, 0, len(combs))
	for _, combo := range combs {
		str, valid := convertToWatchString(combo)
		if !valid {
			continue
		}

		_, ok := resMap[str]
		if ok {
			continue
		}

		results = append(results, str)
		resMap[str] = true
	}

	return results
}

func convertToWatchString(combo Combo) (string, bool) {
	var hours int
	var mins int
	for _, h := range combo.HoursNums {
		hours += h
	}
	for _, m := range combo.MinutesNums {
		mins += m
	}

	if hours > 11 || mins > 59 {
		return "", false
	}

	var hoursStr string = strconv.Itoa(hours)
	var minsStr string = strconv.Itoa(mins)
	if mins < 10 {
		minsStr = "0" + minsStr
	}

	return hoursStr + ":" + minsStr, true
}
