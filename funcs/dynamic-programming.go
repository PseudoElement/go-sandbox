package funcs

/*
word1 = "horse"
word2 = "ros"
	""	r	o	s
""	0	1	2	3
h	1	1	2	3
o	2	2	1	2
r	3	2	2	2
s	4	3	3	2
e	5	4	4	3


word1 = "intention"
word2 = "execution"
	""	e	x	e	c	u	t	i	o	n
""	0	1	2	3	4	5	6	7	8	9
i	1	1	2	3	4	5	6	6	7	8
n	2	2	2	3	4	5	6	7	7	7
t	3	3	3	3	4	5	6	7	8	8
e	4	3	4	3	4	5	6	7	8	9
n	5	4	4	4	4	5	6	7	8	8
t	6	5	5	5	5	5	6	7	8	9
i	7	6	6	6	6	6	6	6	7	8
o	8	7	7	7	7	7	7	7	6	7
n	9	8	8	8	8	8	8	8	7	6

*/

func MinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	// Create DP table
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = i // Deleting all characters in word1
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // Inserting all characters of word2
	}

	// Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				// Characters match, no operation needed
				dp[i][j] = dp[i-1][j-1]
			} else {
				// Take the minimum of Insert, Delete, Replace
				dp[i][j] = 1 + Min(dp[i-1][j], Min(dp[i][j-1], dp[i-1][j-1]))
			}
		}
	}

	return dp[m][n]
}

// Helper function to find minimum of two numbers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
