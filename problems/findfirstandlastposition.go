package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	first := findFirst(nums, target)
	last := findLast(nums, target)

	return []int{first, last}
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid
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

func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid
		}
	}

	if nums[right] == target {
		return right
	}

	if nums[left] == target {
		return left
	}
	return -1
}
