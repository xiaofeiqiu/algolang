package main

const x = 0
const y = 1

type Node1 [2]int

var deltas = []Node1{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	seen := make(map[Node1]bool)

	res := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' && !seen[Node1{i, j}] {
				bfs1(grid, Node1{i, j}, seen)
				res++
			}
		}
	}

	return res
}

func bfs1(grid [][]byte, node Node1, seen map[Node1]bool) {
	m, n := len(grid), len(grid[0])
	q := []Node1{node}
	seen[node] = true

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, delta := range deltas {
			adj := Node1{curr[x] + delta[x], curr[y] + delta[y]}
			if adj[x] < 0 || adj[x] >= m || adj[y] < 0 || adj[y] >= n {
				continue
			}

			if seen[adj] {
				continue
			}

			if grid[adj[x]][adj[y]] != '1' {
				continue
			}

			q = append(q, adj)
			seen[adj] = true
		}
	}
}
