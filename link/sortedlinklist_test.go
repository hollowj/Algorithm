package link

import (
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T) {
	skipList := NewSortedLinkList()
	skipList.Add(4)
	skipList.Add(1)
	skipList.Add(3)
	skipList.Add(0)
	fmt.Println(skipList.ToArray())
	skipList.Add(3)
	fmt.Println(skipList.ToArray())
	//skipList.Erase(3)
	//fmt.Println(skipList.ToArray())
	result := skipList.Range(1, 3)
	fmt.Println(result)
}
