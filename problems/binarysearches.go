package main

import (
	"fmt"
)

// Implementations

// Standard Binary Search
func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Leftmost Target Search
func findLeftmost(nums []int, target int) int {
	low, high := 0, len(nums)-1
	result := -1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			result = mid
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}

// Rightmost Target Search
func findRightmost(nums []int, target int) int {
	low, high := 0, len(nums)-1
	result := -1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			result = mid
			low = mid + 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}

// Closest Smaller Element
func findClosestSmaller(nums []int, target int) int {
	low, high := 0, len(nums)-1
	result := -1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}

// Closest Larger Element
func findClosestLarger(nums []int, target int) int {
	low, high := 0, len(nums)-1
	result := -1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result
}

// Test Functions
func testBinarySearch() {
	fmt.Println("Testing binarySearch...")

	tests := []struct {
		nums    []int
		target  int
		expect  int
	}{
		{[]int{}, 3, -1}, // Empty array
		{[]int{1}, 1, 0}, // Single element array, target found
		{[]int{1}, 2, -1}, // Single element array, target not found
		{[]int{1, 2, 3, 4, 5}, 3, 2}, // Target in the middle
		{[]int{1, 2, 3, 4, 5}, 1, 0}, // Target at the beginning
		{[]int{1, 2, 3, 4, 5}, 5, 4}, // Target at the end
		{[]int{1, 2, 3, 4, 5}, 6, -1}, // Target not found
	}

	for _, test := range tests {
		result := binarySearch(test.nums, test.target)
		if result != test.expect {
			fmt.Printf("FAIL: binarySearch(%v, %d) => %d, expected %d\n", test.nums, test.target, result, test.expect)
		} else {
			fmt.Printf("PASS: binarySearch(%v, %d) => %d\n", test.nums, test.target, result)
		}
	}
}

func testFindLeftmost() {
	fmt.Println("Testing findLeftmost...")

	tests := []struct {
		nums    []int
		target  int
		expect  int
	}{
		{[]int{}, 3, -1}, // Empty array
		{[]int{1}, 1, 0}, // Single element array, target found
		{[]int{1}, 2, -1}, // Single element array, target not found
		{[]int{1, 2, 3, 4, 5}, 3, 2}, // Target in the middle
		{[]int{1, 2, 3, 4, 5}, 1, 0}, // Target at the beginning
		{[]int{1, 2, 3, 4, 5}, 5, 4}, // Target at the end
		{[]int{1, 2, 3, 4, 5}, 6, -1}, // Target not found
	}

	for _, test := range tests {
		result := findLeftmost(test.nums, test.target)
		if result != test.expect {
			fmt.Printf("FAIL: findLeftmost(%v, %d) => %d, expected %d\n", test.nums, test.target, result, test.expect)
		} else {
			fmt.Printf("PASS: findLeftmost(%v, %d) => %d\n", test.nums, test.target, result)
		}
	}
}

func testFindRightmost() {
	fmt.Println("Testing findRightmost...")

	tests := []struct {
		nums    []int
		target  int
		expect  int
	}{
		{[]int{}, 3, -1}, // Empty array
		{[]int{1}, 1, 0}, // Single element array, target found
		{[]int{1, 1, 1, 1}, 1, 3}, // All elements are target
		{[]int{1, 2, 2, 3, 3, 3, 4}, 3, 5}, // Multiple occurrences
		{[]int{1, 2, 3, 4, 5}, 2, 1}, // Target appears once
		{[]int{1, 2, 3, 4, 5}, 6, -1}, // Target not found
	}

	for _, test := range tests {
		result := findRightmost(test.nums, test.target)
		if result != test.expect {
			fmt.Printf("FAIL: findRightmost(%v, %d) => %d, expected %d\n", test.nums, test.target, result, test.expect)
		} else {
			fmt.Printf("PASS: findRightmost(%v, %d) => %d\n", test.nums, test.target, result)
		}
	}
}

func testFindClosestSmaller() {
	fmt.Println("Testing findClosestSmaller...")

	tests := []struct {
		nums    []int
		target  int
		expect  int
	}{
		{[]int{}, 3, -1}, // Empty array
		{[]int{1}, 1, -1}, // Single element array, no smaller element
		{[]int{1, 2, 3, 4, 5}, 3, 1}, // Closest smaller exists
		{[]int{1, 2, 3, 4, 5}, 1, -1}, // No smaller element
		{[]int{1, 2, 3, 4, 5}, 6, 4}, // Target is larger than the largest element
	}

	for _, test := range tests {
		result := findClosestSmaller(test.nums, test.target)
		if result != test.expect {
			fmt.Printf("FAIL: findClosestSmaller(%v, %d) => %d, expected %d\n", test.nums, test.target, result, test.expect)
		} else {
			fmt.Printf("PASS: findClosestSmaller(%v, %d) => %d\n", test.nums, test.target, result)
		}
	}
}

func testFindClosestLarger() {
	fmt.Println("Testing findClosestLarger...")

	tests := []struct {
		nums    []int
		target  int
		expect  int
	}{
		{[]int{}, 3, -1}, // Empty array
		{[]int{1}, 1, -1}, // Single element array, no larger element
		{[]int{1, 2, 3, 4, 5}, 3, 3}, // Closest larger exists
		{[]int{1, 2, 3, 4, 5}, 5, -1}, // No larger element
		{[]int{1, 2, 3, 4, 5}, 0, 0}, // Target is smaller than the smallest element
	}

	for _, test := range tests {
		result := findClosestLarger(test.nums, test.target)
		if result != test.expect {
			fmt.Printf("FAIL: findClosestLarger(%v, %d) => %d, expected %d\n", test.nums, test.target, result, test.expect)
		} else {
			fmt.Printf("PASS: findClosestLarger(%v, %d) => %d\n", test.nums, test.target, result)
		}
	}
}

func main() {
	testBinarySearch()
	testFindLeftmost()
	testFindRightmost()
	testFindClosestSmaller()
	testFindClosestLarger()
}