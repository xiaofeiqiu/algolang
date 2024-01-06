package main

type TrieNode2 struct {
	children [26]*TrieNode2
	word     string
}

func (r *TrieNode2) AddWord(word string) {
	node := r
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			node.children[index] = &TrieNode2{children: [26]*TrieNode2{}}
		}
		node = node.children[index]
	}
	node.word = word
}

func buildTrie(words []string) *TrieNode2 {
	root := &TrieNode2{children: [26]*TrieNode2{}}
	for _, word := range words {
		root.AddWord(word)
	}
	return root
}

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return []string{}
	}

	res := make([]string, 0)
	m := len(board)
	n := len(board[0])
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	var dfs func(node *TrieNode2, i, j int)
	dfs = func(node *TrieNode2, i, j int) {
		index := board[i][j] - 'a'
		if node.children[index] == nil {
			return
		}

		node = node.children[index]
		if node.word != "" {
			res = append(res, node.word)
			node.word = ""
		}

		seen[i][j] = true
		for _, v := range dirs {
			x, y := i+v[0], j+v[1]
			if x >= 0 && x < m && y >= 0 && y < n && !seen[x][y] {
				seen[x][y] = true
				dfs(node, x, y)
				seen[x][y] = false
			}
		}
		seen[i][j] = false
	}

	root := buildTrie(words)
	for i := range board {
		for j := range board[0] {
			dfs(root, i, j)
		}
	}

	return res
}
