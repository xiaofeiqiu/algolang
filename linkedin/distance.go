package linkedin

/*
找两个会员之间的connection距离，其实就是abstract成graph，求两个node直接edge的个数。
十分钟BFS秒掉以后，给了个follow-up，问如果一个member的connection特别多怎么办，写了个两头BFS，面试官满意。
*/

func bidirectionalBFS(graph map[int][]int, start, goal int) int {
	if start == goal {
		return 0
	}

	// 初始化队列和访问记录
	startQueue := []int{start}
	goalQueue := []int{goal}

	startVisited := map[int]int{start: 0}
	goalVisited := map[int]int{goal: 0}

	// 双向BFS
	for len(startQueue) > 0 && len(goalQueue) > 0 {
		if res := bfsStep(graph, &startQueue, startVisited, goalVisited); res != -1 {
			return res
		}

		if res := bfsStep(graph, &goalQueue, goalVisited, startVisited); res != -1 {
			return res
		}
	}

	return -1 // 没有路径
}

func bfsStep(graph map[int][]int, queue *[]int, visited, otherVisited map[int]int) int {
	current := (*queue)[0]
	*queue = (*queue)[1:]

	for _, neighbor := range graph[current] {
		if _, found := visited[neighbor]; found {
			continue
		}

		visited[neighbor] = visited[current] + 1

		if dist, found := otherVisited[neighbor]; found {
			return visited[neighbor] + dist
		}

		*queue = append(*queue, neighbor)
	}

	return -1
}
