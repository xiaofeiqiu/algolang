package main

import "sort"

func findClosestElements(arr []int, k int, x int) []int {
	if arr == nil {
		return nil
	}

	// binary search to find the left bound
	index := findJustSmallerOrEqualIndex(arr, x)

	// use two poiter expand from that index
	left := index
	right := index + 1

	res := make([]int, 0)
	for k > 0 {
		if isLeftCloser(arr, left, right, x) {
			res = append(res, arr[left])
			left--
		} else {
			res = append(res, arr[right])
			right++
		}
		k--
	}

	sort.Ints(res)
	return res
}

func findJustSmallerOrEqualIndex(arr []int, x int) int {
	left, right := 0, len(arr)-1
	for left+1 < right {
		mid := left + (right-left)/2

		if arr[mid] >= x {
			right = mid
		} else {
			left = mid
		}
	}

	if arr[right] <= x {
		return right
	}

	if arr[left] <= x {
		return left
	}

	return -1
}

func isLeftCloser(arr []int, left int, right int, x int) bool {
	if left < 0 {
		return false
	}

	if right >= len(arr) {
		return true
	}

	return abs(arr[left]-x) <= abs(arr[right]-x)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
