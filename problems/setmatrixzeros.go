package main

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	firstRowZero, firstColZero := false, false

	// Check if first row has zero
	for _, v := range matrix[0] {
		if v == 0 {
			firstRowZero = true
			break
		}
	}

	// Check if first column has zero
	for _, row := range matrix {
		if row[0] == 0 {
			firstColZero = true
			break
		}
	}

	// Use first row and column as flags
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// Zero out cells based on flags
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Zero out first row if needed
	if firstRowZero {
		for j := range matrix[0] {
			matrix[0][j] = 0
		}
	}

	// Zero out first column if needed
	if firstColZero {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
}
