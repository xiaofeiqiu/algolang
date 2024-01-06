package main

func robii(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(robFromTo(nums, 0, n-2), robFromTo(nums, 1, n-1))
}

func robFromTo(nums []int, start, end int) int {

	dp := make([]int, end-start+2)
	dp[1] = nums[start]

	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[start+i-1])
	}
	return dp[len(dp)-1]
}
