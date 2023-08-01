package main

import "fmt"

func main() {
	strs := []string{"abc", "abd", "bcf", "abcd"}
	trieTree := NewTrieTree()
	for _, str := range strs {
		trieTree.Insert(str)
	}
	fmt.Println(trieTree)
	count := trieTree.Find("abc")
	fmt.Println(count)
	count = trieTree.FindPrefix("abc")
	fmt.Println(count)
	//trieTree.Insert("bcf")
	find := trieTree.Find("bcf")
	trieTree.Remove("bcf")
	find = trieTree.Find("bcf")
	fmt.Println(find)
	fmt.Printf("%d", 'a')
}
