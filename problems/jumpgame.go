package main

func canJump(nums []int) bool {
	if nums == nil {
		return false
	}

	maxIndex := 0
	for i := range nums {
		// maxIndex smaller than current index, return false
		if maxIndex < i {
			return false
		}

		// update maxIndex
		if nums[i]+i > maxIndex {
			maxIndex = i + nums[i]
		}

		if maxIndex >= len(nums)-1 {
			return true
		}
	}

	return false
}
