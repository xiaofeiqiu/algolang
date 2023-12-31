package main

func search(nums []int, target int) int {
	if nums == nil {
		return -1
	}

	left, right := 0, len(nums)-1

	for left+1 < right {
		mid := left + (right-left)/2

		if nums[mid] >= nums[0] {
			if target >= nums[0] && target < nums[mid] {
				right = mid
			} else {
				left = mid
			}
		} else {
			if target <= nums[right] && target > nums[mid] {
				left = mid
			} else {
				right = mid
			}
		}
	}

	if nums[left] == target {
		return left
	}

	if nums[right] == target {
		return right
	}
	return -1
}
