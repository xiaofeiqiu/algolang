package main

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	// Initialize adjacency list
	graph := make(map[int][]int)
	for _, edge := range edges {
		v, w := edge[0], edge[1]
		graph[v] = append(graph[v], w)
		graph[w] = append(graph[w], v)
	}

	// Initialize visited set and queue
	visited := make(map[int]bool)
	queue := []int{0}
	visited[0] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}

	// Check if all nodes are visited
	return len(visited) == n
}
