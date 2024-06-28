package funcs

import (
	"math"
	"strconv"
)

func AllCombsForNLengthArray(length float64) [][]bool {
	bools := [][]bool{}

	for i := 0; float64(i) < math.Pow(float64(2), length); i++ {
		bin := strconv.FormatInt(int64(i), 2)
		for float64(len(bin)) < length {
			bin = "0" + bin
		}

		boolArray := []bool{}
		var count = 0
		for _, ch := range bin {
			boolArray = append(boolArray, string(ch) == "0")
			count++
		}
		bools = append(bools, boolArray)
	}

	return bools
}
