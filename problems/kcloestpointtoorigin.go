package main

func kClosest(points [][]int, k int) [][]int {

	index := quickSelect(points, 0, len(points)-1, k-1)

	res := make([][]int, 0)
	for index >= 0 {
		res = append(res, points[index])
		index--
	}

	return res
}

func quickSelect(points [][]int, start int, end int, k int) int {
	if start == end {
		return start
	}

	mid := start + (end-start)/2
	pivot := getDist(points[mid])

	left, right := start, end
	for left <= right {
		for getDist(points[left]) < pivot {
			left++
		}

		for getDist(points[right]) > pivot {
			right--
		}

		if left <= right {
			points[left], points[right] = points[right], points[left]
			left++
			right--
		}
	}

	if k <= right {
		return quickSelect(points, 0, right, k)
	}

	if k >= left {
		return quickSelect(points, left, end, k)
	}

	return k
}

func getDist(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}
