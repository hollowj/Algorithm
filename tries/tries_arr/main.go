package main

import (
	"fmt"
)

func main() {
	trieTree := Constructor()
	trieTree.Insert("apple")
	search := trieTree.Search("apple")
	fmt.Println(search)
	search = trieTree.Search("app")
	fmt.Println(search)
	search = trieTree.StartsWith("app")
	fmt.Println(search)
	trieTree.Insert("app")
	search = trieTree.Search("app")
	fmt.Println(search)
}
