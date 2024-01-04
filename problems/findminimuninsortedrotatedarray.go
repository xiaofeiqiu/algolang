package main

func findMin(nums []int) int {
	if nums == nil {
		return 0
	}

	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid
		} else {
			right = mid
		}
	}

	return min(nums[left], nums[right])
}
