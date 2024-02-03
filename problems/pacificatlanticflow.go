package main

func pacificAtlantic(heights [][]int) [][]int {
	if heights == nil || heights[0] == nil {
		return [][]int{}
	}

	m, n := len(heights), len(heights[0])

	canReachPacific := make([][]bool, m)
	canReachAtlantic := make([][]bool, m)

	for i := range heights {
		canReachPacific[i] = make([]bool, n)
		canReachAtlantic[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		dfs(heights, i, 0, canReachPacific)
		dfs(heights, i, n-1, canReachAtlantic)
	}

	for i := 0; i < n; i++ {
		dfs(heights, 0, i, canReachPacific)
		dfs(heights, m-1, i, canReachAtlantic)
	}

	res := make([][]int, 0)
	for i := range canReachPacific {
		for j := range canReachPacific[i] {
			if canReachPacific[i][j] && canReachAtlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func dfs(heights [][]int, x, y int, res [][]bool) {

	m, n := len(heights), len(heights[0])
	res[x][y] = true

	for i := range dirs {
		newX, newY := x+dirs[i][0], y+dirs[i][1]
		if newX < 0 || newX >= m || newY < 0 || newY >= n {
			continue
		}

		if res[newX][newY] {
			continue
		}

		if heights[newX][newY] < heights[x][y] {
			continue
		}

		dfs(heights, newX, newY, res)
	}

}
