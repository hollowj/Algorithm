package link

import (
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T) {
	//skipList := NewSortedLinkList()
	for i := 0; i < 10000; i++ {
		TestSkip(t)
	}
}

func TestSkip(t *testing.T) {
	skipList := NewSkipList()
	skipList.Add(4)
	skipList.Add(1)
	skipList.Add(3)
	skipList.Add(0)
	fmt.Println(skipList.ToArray())
	skipList.Add(3)
	fmt.Println(skipList.ToArray())
	if len(skipList.head.next) > 1 {
		fmt.Println(1)
	}
	skipList.Erase(3)
	fmt.Println(skipList.ToArray())
	result := skipList.Range(1, 3)
	fmt.Println(result)
}
