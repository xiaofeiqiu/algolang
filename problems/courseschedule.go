package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 {
		return true
	}

	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	for _, v := range prerequisites {
		from, to := v[1], v[0]
		graph[from] = append(graph[from], to)
		inDegree[to]++
	}

	q := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	count := 0
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		count++
		adjs := graph[curr]
		for _, adj := range adjs {
			inDegree[adj]--
			if inDegree[adj] == 0 {
				q = append(q, adj)
			}
		}
	}

	return count == numCourses
}
