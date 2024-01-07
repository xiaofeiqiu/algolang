package main

func missingNumber(nums []int) int {
	n := len(nums)
	totalSum := 0
	for i := 1; i <= n; i++ {
		totalSum += i
	}

	actual := 0
	for i := range nums {
		actual += nums[i]
	}

	return totalSum - actual
}
