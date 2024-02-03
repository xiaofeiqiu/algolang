package main

import "math"

func increasingTriplet(nums []int) bool {
	if nums == nil {
		return false
	}

	small := math.MaxInt
	secondSmall := math.MaxInt

	for i := range nums {
		if nums[i] <= small {
			small = nums[i]
			continue
		}

		if nums[i] <= secondSmall {
			secondSmall = nums[i]
			continue
		}

		return true
	}
	return false
}
