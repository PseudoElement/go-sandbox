package funcs

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price // Update the minimum price
			continue
		}

		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit // Update the maximum profit
		}
	}

	return maxProfit
}
