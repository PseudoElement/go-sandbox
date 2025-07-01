package funcs

import (
	"log"
	"slices"
	"strconv"
	"strings"
	"time"
)

func Combine(n int, k int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	possibleNums := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		possibleNums = append(possibleNums, i)
	}
	log.Println("possibleNums==>", possibleNums)

	totalCombsCount := factorial(n) / (factorial(k) * factorial(n-k))
	combinations := make([][]int, 0, totalCombsCount)
	combsMap := make(map[string]bool, totalCombsCount)

	startPointer := 0
	endPointer := 0
	combination := make([]int, 0, k)
	for len(combinations) < totalCombsCount {
		log.Println("first -", possibleNums[startPointer], "second -", possibleNums[endPointer])
		if len(combination) == 0 {
			combination = append(combination, possibleNums[startPointer])
		}
		if !sliceContains(combination, possibleNums[endPointer]) {
			combination = append(combination, possibleNums[endPointer])
		}

		if len(combination) == k {
			key := mapKey(combination...)
			has := mapHas(combsMap, key)
			if !has {
				combinations = append(combinations, combination)
				combsMap[key] = true
				combination = []int{}
			}
		}

		if endPointer == len(possibleNums)-1 {
			endPointer = 0
			startPointer++
		} else {
			endPointer++
		}

		time.Sleep(200 * time.Millisecond)
	}

	return combinations
}

func sliceContains(combination []int, num int) bool {
	for _, n := range combination {
		if n == num {
			return true
		}
	}
	return false
}

func mapKey(keys ...int) string {
	slices.Sort(keys)
	stringKeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		stringKeys[i] = strconv.Itoa(keys[i])
	}

	str := strings.Join(stringKeys, "")
	return str
}

func mapHas(m map[string]bool, key string) bool {
	_, ok := m[key]
	return ok
}

func factorial(value int) int {
	res := 1
	for i := value; i > 0; i-- {
		res = res * i
	}
	return res
}
