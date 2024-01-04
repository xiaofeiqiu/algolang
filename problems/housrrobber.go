package main

func rob(nums []int) int {
	if nums == nil {
		return 0
	}

	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]

	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[len(nums)]
}
