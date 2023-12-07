package main

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	res := make([][]int, n+1)

	// Initialize the last row of res to 0
	res[n] = make([]int, n+1)

	// Start from the bottom row of the triangle and move upwards
	for row := n - 1; row >= 0; row-- {
		res[row] = make([]int, row+1) // Initialize each row of res according to the number of columns it should have
		for col := 0; col <= row; col++ {
			res[row][col] = triangle[row][col] + min(res[row+1][col], res[row+1][col+1])
		}
	}

	// Return the minimum path sum starting from the top of the triangle
	return res[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
