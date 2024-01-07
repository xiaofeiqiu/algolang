package main

func alienOrder(words []string) string {
	if len(words) == 0 {
		return ""
	}

	graph, indegree := buildGraph(words)
	return toposort(graph, indegree)
}

func toposort(graph map[byte]set, inDegreeMap map[byte]int) string {
	if graph == nil || inDegreeMap == nil {
		return ""
	}
	n := len(inDegreeMap)
	res := make([]byte, 0)
	q := make([]byte, 0)

	for char, indegree := range inDegreeMap {
		if indegree == 0 {
			q = append(q, char)
		}
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		res = append(res, curr)

		for adj := range graph[curr] {
			inDegreeMap[adj]--
			if inDegreeMap[adj] == 0 {
				q = append(q, adj)
			}
		}
	}

	if n == len(res) {
		return string(res)
	}
	return ""
}

type set map[byte]bool

func buildGraph(words []string) (map[byte]set, map[byte]int) {

	graph := make(map[byte]set)
	inDegree := make(map[byte]int)

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			char := word[i]
			if _, exists := graph[char]; !exists {
				graph[char] = make(set)
			}
			inDegree[char] = 0 // Initialize inDegree for all characters
		}
	}

	for i := 1; i < len(words); i++ {
		w1, w2 := words[i-1], words[i]
		for j := 0; j < len(w1) && j < len(w2); j++ {
			c1, c2 := w1[j], w2[j]
			if c1 != c2 {
				if !graph[c1][c2] {
					graph[c1][c2] = true
					inDegree[c2]++
				}
				break
			}

			if len(w1) > len(w2) && j == len(w2)-1 {
				return nil, nil
			}
		}

	}
	return graph, inDegree
}
