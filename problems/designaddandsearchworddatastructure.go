package main

type WordDictionary struct {
	root *TrieNode
}

func Constructor1() WordDictionary {
	return WordDictionary{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isWord = true
}

func (this *WordDictionary) searchInNode(word string, node *TrieNode) bool {
	for i, ch := range word {
		if ch == '.' {
			for _, childNode := range node.children {
				if this.searchInNode(word[i+1:], childNode) {
					return true
				}
			}
			return false
		}

		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return node.isWord
}

func (this *WordDictionary) Search(word string) bool {
	return this.searchInNode(word, this.root)
}
