package main

import (
	"fmt"

	"github.com/hollowj/algorithm/unionfindset/bak"
)

func main() {
	set := bak.NewUnionFindSet(10)
	set.Merge(1, 2)
	set.Merge(3, 4)
	set.Merge(5, 2)
	set.Merge(4, 6)
	set.Merge(2, 6)
	set.Merge(8, 7)
	set.Merge(9, 7)
	set.Merge(1, 6)
	set.Merge(2, 4)
	fmt.Println(set.Sum())
	//fmt.Println(set.arr)
}
