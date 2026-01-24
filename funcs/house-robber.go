package funcs

func Rob(treasure []int) int {
	if len(treasure) == 0 {
		return 0
	}

	// initialize dp array
	dp := make([]int, len(treasure)+1)

	// fill in base cases (dp[0] = 0 already)
	dp[1] = treasure[0]

	// iterate to fill in the rest of the array
	for i := 2; i <= len(treasure); i++ {
		// fill in dp[i] using the recurrence relation
		take := dp[i-2] + treasure[i-1]
		skip := dp[i-1]
		dp[i] = max(take, skip)
	}

	return dp[len(treasure)]
}
