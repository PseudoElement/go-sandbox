package funcs

import (
	"slices"
)

func MaxProfit(prices []int) int {
	// key - value
	// value - index
	order := make(map[int]int)
	for i := 0; i < len(prices); i++ {
		value := prices[i]
		order[value] = i
	}

	slices.Sort(prices)

	start := 0
	end := len(prices) - 1

	profit := findMaxProfit(start, end, prices, order)

	return profit
}

func findMaxProfit(start int, end int, sortedPrices []int, order map[int]int) int {
	min := sortedPrices[start]
	max := sortedPrices[end]
	idxMin := order[min]
	idxMax := order[max]

	if idxMin < idxMax {
		return max - min
	}

	nextMin := sortedPrices[start+1]
	nextDeltaMin := nextMin - min
	nextMax := sortedPrices[end-1]
	nextDeltaMax := max - nextMax

	if nextDeltaMax > nextDeltaMin {
		start++
	} else {
		end--
	}

	return findMaxProfit(start, end, sortedPrices, order)
}
