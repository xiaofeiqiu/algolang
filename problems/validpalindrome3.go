package main

func isValidPalindrome(s string, k int) bool {
	if s == "" {
		return true
	}

	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for len := 2; len <= n; len++ {
		for i := 0; i <= n-len; i++ {
			j := i + len - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}

	return dp[0][n-1] <= k
}
