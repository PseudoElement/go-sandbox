package funcs

import (
	"log"
	"strconv"
)

func NumDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		digit := int(s[i-1] - '0')
		if digit != 0 {
			dp[i] += dp[i-1]
		}

		digit, _ = strconv.Atoi(s[i-2 : i])
		if digit >= 10 && digit <= 26 {
			dp[i] += dp[i-2]
		}
	}
	log.Println(dp[n])
	return dp[n]
}
