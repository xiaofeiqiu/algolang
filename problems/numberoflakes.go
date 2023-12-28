package main

func countLakes(grid [][]rune) int {
	if grid == nil {
		return 0
	}

	m := len(grid)
	n := len(grid[0])
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	var dfs func(r int, c int) bool
	dfs = func(r int, c int) bool {
		if r >= len(grid) || c >= len(grid[0]) || r < 0 || c < 0 {
			return false
		}

		// seen meaning the cell was previously calculated,
		if grid[r][c] == 'x' || seen[r][c] {
			return true
		}

		seen[r][c] = true

		left := dfs(r, c-1)
		right := dfs(r, c+1)
		up := dfs(r-1, c)
		down := dfs(r+1, c)

		return left && right && up && down
	}

	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '.' && !seen[i][j] && dfs(i, j) {
				res++
			}
		}
	}

	return res
}
