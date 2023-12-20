package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	subarraySum := nums[0] // use

	for i := 1; i < len(nums); i++ {
		subarraySum = max(nums[i], subarraySum+nums[i])
		if subarraySum > res {
			res = subarraySum
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
