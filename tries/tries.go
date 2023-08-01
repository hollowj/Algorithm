package main

type TrieNode struct {
	pass  int
	end   int
	nexts map[int]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{nexts: make(map[int]*TrieNode)}
}

type Trie struct {
	root *TrieNode
}

func NewTrieTree() *Trie {
	return &Trie{root: NewTrieNode()}
}
func (this *Trie) Insert(str string) {
	node := this.root
	node.pass++
	for _, char := range str {
		charNo := int(char)
		n, ok := node.nexts[charNo]
		if !ok {
			n = NewTrieNode()
			node.nexts[charNo] = n
		}
		n.pass++
		node = n
	}
	node.end++
}
func (this *Trie) Find(str string) int {
	node := this.root
	for _, char := range str {
		charNo := int(char)
		n, ok := node.nexts[charNo]
		if !ok {
			return 0
		}
		node = n
	}
	return node.end
}
func (this *Trie) FindPrefix(str string) int {
	node := this.root
	for _, char := range str {
		charNo := int(char)
		n, ok := node.nexts[charNo]
		if !ok {
			return node.pass
		}
		node = n
	}
	return node.pass
}
func (this *Trie) Remove(str string) {
	node := this.root
	for _, char := range str {
		charNo := int(char)
		n, ok := node.nexts[charNo]
		if !ok {
			return
		} else {
			n.pass--
		}
		if n.pass == 0 {
			delete(node.nexts, charNo)
			return
		}
		node = n
	}
	node.end--
	return
}
