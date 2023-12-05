package main

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	left, right := start, end
	pivot := nums[start+(end-start)/2]

	for left <= right {
		for left <= right && nums[left] < pivot {
			left++
		}

		for left <= right && nums[right] > pivot {
			right--
		}

		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	quickSort(nums, start, right)
	quickSort(nums, left, end)
}
