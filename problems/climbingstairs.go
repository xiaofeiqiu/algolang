package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	// total n + 1 stairs
	// dp[i] to represent number of ways to climb to ith stairs
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
