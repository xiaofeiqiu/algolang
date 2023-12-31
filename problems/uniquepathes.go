package main

// dp[0][j] = 1
// dp[i][0] = 1
// dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
// return dp[m - 1][n - 1]
func uniquePaths(m int, n int) int {

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := range dp[0] {
		dp[0][i] = 1
	}
	for i := range dp {
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
