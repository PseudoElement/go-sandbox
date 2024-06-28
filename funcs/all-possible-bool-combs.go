package funcs

import (
	"fmt"
	"math"
	"strconv"
)

func AllCombsForNLengthArray(length float64) [][]bool {
	bools := [][]bool{}

	for i := 0; float64(i) < math.Pow(float64(2), length); i++ {
		bin := strconv.FormatInt(int64(i), 2)
		fmt.Println("BIN_1 ===> ", bin)
		for float64(len(bin)) < length {
			bin = "0" + bin
		}
		fmt.Println("BIN_2 ===> ", bin)

		boolArray := []bool{}
		for _, ch := range bin {
			boolArray = append(boolArray, string(ch) == "0")
		}
		bools = append(bools, boolArray)
	}

	return bools
}
