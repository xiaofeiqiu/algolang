package main

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	if beginWord == "" || endWord == "" || wordList == nil {
		return [][]string{}
	}

	wordSet := make(map[string]bool)
	for _, v := range wordList {
		wordSet[v] = true
	}

	// make sure to check end word exist in the word set or not
	if _, exist := wordSet[endWord]; !exist {
		return [][]string{}
	}

	// add begin word to the set
	wordSet[beginWord] = true

	distanceMap := bfs(beginWord, endWord, wordSet)
	res := make([][]string, 0)
	dfs(beginWord, endWord, wordSet, distanceMap, []string{beginWord}, &res)
	return res
}

func dfs(currWord string, endWord string, wordSet map[string]bool, distanceMap map[string]int, path []string, res *[][]string) {

	if currWord == endWord {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	adjs := getAdjs(currWord, wordSet)
	for _, adj := range adjs {
		if distanceMap[currWord]-1 == distanceMap[adj] {
			dfs(adj, endWord, wordSet, distanceMap, append(path, adj), res)
		}
	}
}

// bfs build distance map, starting from endword
func bfs(beginWord, endWord string, wordSet map[string]bool) map[string]int {
	queue := make([]string, 0)
	queue = append(queue, endWord)

	res := make(map[string]int)
	res[endWord] = 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == beginWord {
			return res
		}

		adjs := getAdjs1(curr, wordSet)
		for _, adj := range adjs {
			if _, seen := res[adj]; seen {
				continue
			}

			res[adj] = res[curr] + 1
			queue = append(queue, adj)
		}
	}
	return res
}

func getAdjs1(curr string, wordSet map[string]bool) []string {
	res := make([]string, 0)
	chars := []rune(curr)

	for i, char := range chars {
		for c := 'a'; c <= 'z'; c++ {
			if char == c {
				continue
			}
			chars[i] = c
			newWord := string(chars)
			chars[i] = char

			if _, exist := wordSet[newWord]; !exist {
				continue
			}
			res = append(res, newWord)
		}
	}
	return res
}
