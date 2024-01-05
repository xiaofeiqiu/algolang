package main

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, ch := range word {
		if _, exist := node.children[ch]; !exist {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isWord = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, ch := range word {
		if _, exist := node.children[ch]; !exist {
			return false
		}
		node = node.children[ch]
	}
	return node.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, ch := range prefix {
		if _, exist := node.children[ch]; !exist {
			return false
		}
		node = node.children[ch]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
