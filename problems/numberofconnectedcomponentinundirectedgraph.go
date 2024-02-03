package main

func countComponents(n int, edges [][]int) int {
	if n == 0 || edges == nil {
		return 0
	}

	graph := make(map[int][]int)
	for _, edge := range edges {
		w, v := edge[0], edge[1]
		graph[w] = append(graph[w], v)
		graph[v] = append(graph[v], w)
	}

	seen := make(map[int]bool)
	count := 0
	for i := 0; i < n; i++ {
		if !seen[i] {
			bfs2(graph, i, seen)
			count++
		}
	}

	return count
}

func bfs2(graph map[int][]int, node int, seen map[int]bool) {

	q := []int{node}
	seen[node] = true

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, adj := range graph[curr] {
			if !seen[adj] {
				q = append(q, adj)
				seen[adj] = true
			}
		}
	}
}
