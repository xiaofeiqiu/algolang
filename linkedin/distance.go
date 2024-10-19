package linkedin

/*
找两个会员之间的connection距离，其实就是abstract成graph，求两个node直接edge的个数。
十分钟BFS秒掉以后，给了个follow-up，问如果一个member的connection特别多怎么办，写了个两头BFS，面试官满意。
*/

func bidirectionalBFS(graph map[int][]int, start, target int) int {
	if start == target {
		return 0
	}

	startSeen := make(map[int]int)
	targetSeen := make(map[int]int)

	startQueue := []int{start}
	targetQueue := []int{target}

	for len(startQueue) > 0 && len(targetQueue) > 0 {
		res := bfs(graph, &startQueue, startSeen, targetSeen)
		if res != -1 {
			return res
		}

		res = bfs(graph, &targetQueue, targetSeen, startSeen)
		if res != -1 {
			return res
		}
	}
	return -1
}

func bfs(graph map[int][]int, startQueue *[]int, startSeen map[int]int, targetSeen map[int]int) int {
	curr := (*startQueue)[0]
	*startQueue = (*startQueue)[1:]

	for _, neighbor := range graph[curr] {
		if _, found := startSeen[neighbor]; found {
			continue
		}

		startSeen[neighbor] = startSeen[curr] + 1
		if dist, found := targetSeen[neighbor]; found {
			return dist + startSeen[neighbor]
		}

		*startQueue = append(*startQueue, neighbor)
	}

	return -1
}
