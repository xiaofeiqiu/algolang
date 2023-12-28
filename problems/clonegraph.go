package main

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	res := &Node{node.Val, nil}

	queue := make([]*Node, 0)
	queue = append(queue, node)

	seen := make(map[int]*Node)
	seen[res.Val] = res

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, adj := range curr.Neighbors {
			if _, exist := seen[adj.Val]; !exist {
				queue = append(queue, adj)
				seen[adj.Val] = &Node{adj.Val, nil}
			}

			seen[curr.Val].Neighbors = append(seen[curr.Val].Neighbors, seen[adj.Val])
		}
	}

	return res
}
