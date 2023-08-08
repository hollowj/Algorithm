package tanxin

import (
	"sort"
	"testing"

	"github.com/emirpasic/gods/lists/singlylinkedlist"
	"github.com/stretchr/testify/assert"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		} else if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return false
	})
	l := singlylinkedlist.New()
	for _, person := range people {
		k := person[1]
		iterator := l.Iterator()
		index := 0
		for iterator.Next() {
			ints := iterator.Value().([]int)
			if ints[0] >= person[0] {
				k--
			}
			if k == 0 {
				index = iterator.Index() + 1
				break
			}
		}
		l.Insert(index, person)

	}
	l.Each(func(index int, value interface{}) {
		people[index] = value.([]int)
	})
	return people
}

func TestT406(t *testing.T) {
	assert.Equal(t, [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}, reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
	assert.Equal(t, [][]int{{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0}}, reconstructQueue([][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}))
}
