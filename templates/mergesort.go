package main

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {

	res := make([]int, 0)
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}

	res = append(res, left[l:]...)
	res = append(res, right[r:]...)

	return res
}
