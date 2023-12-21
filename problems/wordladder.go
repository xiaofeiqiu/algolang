package main

func ladderLength(beginWord string, endWord string, wordList []string) int {

	if beginWord == "" || endWord == "" || wordList == nil {
		return 0
	}

	wordSet := make(map[string]bool)
	for _, v := range wordList {
		wordSet[v] = true
	}

	if _, exist := wordSet[endWord]; !exist {
		return 0
	}

	queue := make([]string, 0)
	visited := make(map[string]bool)

	queue = append(queue, beginWord)
	visited[beginWord] = true

	res := 1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr == endWord {
				return res
			}

			adjs := getAdjs(curr, wordSet)
			for _, adj := range adjs {
				if visited[adj] {
					continue
				}
				queue = append(queue, adj)
				visited[adj] = true
			}
		}
		res++
	}
	return 0
}

func getAdjs(curr string, wordSet map[string]bool) []string {

	res := make([]string, 0)
	chars := []rune(curr)

	// for each char in current word
	for i, char := range chars {
		// try to replace it with all possible chars
		for c := 'a'; c <= 'z'; c++ {
			if char == c {
				continue
			}

			chars[i] = c
			// check if the new word in the word set
			// if yes, add to the res list
			newWord := string(chars)
			if wordSet[newWord] {
				res = append(res, newWord)
			}
			chars[i] = char
		}
	}
	return res
}
