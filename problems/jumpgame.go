package main

func canJump(nums []int) bool {
	if nums == nil {
		return false
	}

	furthest := 0

	for i := range nums {
		if i > furthest {
			return false
		}
		furthest = max(furthest, nums[i]+i)
		if furthest >= len(nums)-1 {
			return true
		}
	}

	return false
}
