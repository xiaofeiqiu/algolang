package main

func rotate(matrix [][]int) {
	n := len(matrix)

	// for each layer
	for layer := 0; layer < n/2; layer++ {
		// define the first index
		first := layer

		// define the last index, exclude the last element for each edge
		last := n - 1 - layer

		// init i from first, j from last
		// i < last, meaning we do not inlcude the last element from left to right or from top to bottom
		// j > first meaning we do not include the first element from right to left or from bottom to top
		for i, j := first, last; i < last && j > first; i, j = i+1, j-1 {

			// list cells need to be swapped
			topCell := &matrix[first][i]   // fixed row, left to right
			rightCell := &matrix[i][last]  // fixed col, top to bottom
			bottomCell := &matrix[last][j] // fixed row, right to left
			leftCell := &matrix[j][first]  // fixed col, bottom to top

			// left -> top
			// top -> right
			// right -> bottom
			// bottom -> top
			*topCell, *rightCell, *bottomCell, *leftCell = *leftCell, *topCell, *rightCell, *bottomCell
		}
	}
}
