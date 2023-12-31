package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, 0)

	top, right, bottom, left := 0, n-1, m-1, 0

	for top <= bottom && left <= right {
		// top row
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		// right column
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		// check if topRow has passed bottomRow
		if top <= bottom {
			// bottom row
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		// check if leftCol has passed rightCol
		if left <= right {
			// left column
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}

	return res
}
