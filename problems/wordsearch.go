package main

func exist(board [][]byte, word string) bool {
	if board == nil || board[0] == nil {
		return false
	}

	m := len(board)
	n := len(board[0])

	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var dfs func(i, j, curr int) bool
	dfs = func(i, j, curr int) bool {
		if curr == len(word)-1 {
			return true
		}

		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}

			if seen[x][y] {
				continue
			}

			if board[x][y] != word[curr+1] {
				continue
			}

			seen[x][y] = true
			if dfs(x, y, curr+1) {
				return true
			}
			seen[x][y] = false
		}

		return false
	}

	for i := range board {
		for j := range board[0] {
			if board[i][j] == word[0] {
				seen[i][j] = true
				if dfs(i, j, 0) {
					return true
				}
				seen[i][j] = false
			}
		}
	}

	return false
}
