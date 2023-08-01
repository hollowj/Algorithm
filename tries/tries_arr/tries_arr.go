package main

type TrieNode struct {
	pass  int
	end   int
	nexts []*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{nexts: make([]*TrieNode, 26)}
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: NewTrieNode()}
}

func (this *Trie) Insert(word string) {
	node := this.root
	node.pass++
	for _, char := range word {
		charNo := int(char) - 97
		n := node.nexts[charNo]
		if n == nil {
			n = NewTrieNode()
			node.nexts[charNo] = n
		}
		n.pass++
		node = n
	}
	node.end++
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, char := range word {
		charNo := int(char) - 97
		n := node.nexts[charNo]
		if n == nil {
			return false
		}
		node = n
	}
	return node.end > 0
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, char := range prefix {
		charNo := int(char) - 97
		n := node.nexts[charNo]
		if n == nil {
			return false
		}
		node = n
	}
	return true
}
