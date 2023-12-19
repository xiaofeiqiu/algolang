package main

import "math"

func find132pattern(nums []int) bool {
	if nums == nil {
		return false
	}

	if len(nums) < 3 {
		return false
	}

	stack := make([]int, 0) // mainting a stack to num_k candidates
	numK := math.MinInt     // mainting a max num_k

	// from right to left
	for i := len(nums) - 1; i >= 0; i-- {

		// if num_k found, and num_i < num_k, pattern found
		if nums[i] < numK {
			return true
		}

		// identify a upward slop from right to left
		// num_j candidate found, look for max num_k in the stack
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			numK = stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop top to maintain the mono decreasing stack
		}

		// also added this num to num_k candidate
		stack = append(stack, nums[i])
	}
	return false
}
