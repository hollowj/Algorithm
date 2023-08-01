package bak

import (
	"fmt"
	"sort"
	"testing"
)

func sum(split []int) int {
	sum := 0
	for _, count := range split {
		sum += count
	}
	return sum
}
func CutGold(split []int) int {
	iTotal := 0
	for _, count := range split {
		iTotal += count
	}
	sort.Slice(split, func(i, j int) bool {
		return i > j
	})
	iConsume := 0
	for i := 0; i < len(split)-1; i++ {
		iConsume += sum(split[i:])
	}
	return iConsume
}

func TestT1(t *testing.T) {
	arr := []int{10, 20, 30}
	gold := CutGold(arr)
	fmt.Println(gold)
}
