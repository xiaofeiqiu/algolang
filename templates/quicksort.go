package main

// QuickSort sorts an array using the QuickSort algorithm.
func QuickSort(arr []int, low, high int) {
	if low < high {
		// Partition the array and get the pivot index
		left, right := partition(arr, low, high)

		// Recursively sort elements before and after partition
		QuickSort(arr, low, left)   // Sort left subarray
		QuickSort(arr, right, high) // Sort right subarray
	}
}

func partition(nums []int, low, high int) (int, int) {
	middle := low + (high-low)/2
	pivot := nums[middle]
	left, right := low, high

	for left <= right {
		// Move left while item < pivot
		for nums[left] < pivot {
			left++
		}
		// Move right while item > pivot
		for nums[right] > pivot {
			right--
		}

		// Swap and move pointers
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	return right, left
}
