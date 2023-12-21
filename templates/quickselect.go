package main

// quickSelect finds the k-th smallest element in an array. k is 0 based index
func quickSelect(nums []int, start, end, k int) int {
	// if only one element left, return that element
	if start == end {
		return nums[start]
	}

	// left points to the last element of the left subarray
	// right points to the first element of the right subarray
	left, right := partition(nums, start, end)

	// if k <= left, meaning kth element in on the left subarray
	// continue to process left subarray
	if k <= left {
		return quickSelect(nums, start, left, k)
	}

	// if k >= right, meaning kth element is on the right subarray
	// continue to process right subarray
	if k >= right {
		return quickSelect(nums, right, end, k)
	}

	// if k is between left, and right, then k is the pivot index
	// so kth element is at its correct position in sorted array
	// return k
	return nums[k]
}
