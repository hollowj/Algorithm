package backtracking

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BackTracking332 struct {
	result []string
}

func NewBackTracking332() *BackTracking332 {
	return &BackTracking332{result: make([]string, 0)}
}
func (b *BackTracking332) dfs(tickets [][]string, data map[string][]int, sFrom string, used map[int]bool) bool {
	if len(b.result) == len(tickets)+1 {
		return true
	}
	tos := data[sFrom]
	for _, toIndex := range tos {
		if used[toIndex] {
			continue
		}
		used[toIndex] = true
		sTo := tickets[toIndex][1]
		b.result = append(b.result, sTo)
		if b.dfs(tickets, data, sTo, used) {
			return true
		}
		delete(used, toIndex)
		b.result = b.result[:len(b.result)-1]
	}
	return false
}
func findItinerary(tickets [][]string) []string {
	data := make(map[string][]int)
	for i, ticket := range tickets {
		sFrom := ticket[0]
		arr, ok := data[sFrom]
		if !ok {
			arr = make([]int, 0)
		}
		arr = append(arr, i)
		data[sFrom] = arr

	}
	for _, ints := range data {
		sort.Slice(ints, func(i, j int) bool {
			return tickets[ints[i]][1] < tickets[ints[j]][1]
		})
	}
	tracking491 := NewBackTracking332()
	sFrom := "JFK"
	tracking491.result = append(tracking491.result, sFrom)
	used := make(map[int]bool)
	tracking491.dfs(tickets, data, sFrom, used)
	return tracking491.result
}
func TestT332(t *testing.T) {
	assert.Equal(t, []string{"JFK", "MUC", "LHR", "SFO", "SJC"}, findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))
	assert.Equal(t, []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}, findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))
	assert.Equal(t, []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}, findItinerary([][]string{{"AXA", "EZE"}, {"EZE", "AUA"}, {"ADL", "JFK"}, {"ADL", "TIA"}, {"AUA", "AXA"}, {"EZE", "TIA"}, {"EZE", "TIA"}, {"AXA", "EZE"}, {"EZE", "ADL"}, {"ANU", "EZE"}, {"TIA", "EZE"}, {"JFK", "ADL"}, {"AUA", "JFK"}, {"JFK", "EZE"}, {"EZE", "ANU"}, {"ADL", "AUA"}, {"ANU", "AXA"}, {"AXA", "ADL"}, {"AUA", "JFK"}, {"EZE", "ADL"}, {"ANU", "TIA"}, {"AUA", "JFK"}, {"TIA", "JFK"}, {"EZE", "AUA"}, {"AXA", "EZE"}, {"AUA", "ANU"}, {"ADL", "AXA"}, {"EZE", "ADL"}, {"AUA", "ANU"}, {"AXA", "EZE"}, {"TIA", "AUA"}, {"AXA", "EZE"}, {"AUA", "SYD"}, {"ADL", "JFK"}, {"EZE", "AUA"}, {"ADL", "ANU"}, {"AUA", "TIA"}, {"ADL", "EZE"}, {"TIA", "JFK"}, {"AXA", "ANU"}, {"JFK", "AXA"}, {"JFK", "ADL"}, {"ADL", "EZE"}, {"AXA", "TIA"}, {"JFK", "AUA"}, {"ADL", "EZE"}, {"JFK", "ADL"}, {"ADL", "AXA"}, {"TIA", "AUA"}, {"AXA", "JFK"}, {"ADL", "AUA"}, {"TIA", "JFK"}, {"JFK", "ADL"}, {"JFK", "ADL"}, {"ANU", "AXA"}, {"TIA", "AXA"}, {"EZE", "JFK"}, {"EZE", "AXA"}, {"ADL", "TIA"}, {"JFK", "AUA"}, {"TIA", "EZE"}, {"EZE", "ADL"}, {"JFK", "ANU"}, {"TIA", "AUA"}, {"EZE", "ADL"}, {"ADL", "JFK"}, {"ANU", "AXA"}, {"AUA", "AXA"}, {"ANU", "EZE"}, {"ADL", "AXA"}, {"ANU", "AXA"}, {"TIA", "ADL"}, {"JFK", "ADL"}, {"JFK", "TIA"}, {"AUA", "ADL"}, {"AUA", "TIA"}, {"TIA", "JFK"}, {"EZE", "JFK"}, {"AUA", "ADL"}, {"ADL", "AUA"}, {"EZE", "ANU"}, {"ADL", "ANU"}, {"AUA", "AXA"}, {"AXA", "TIA"}, {"AXA", "TIA"}, {"ADL", "AXA"}, {"EZE", "AXA"}, {"AXA", "JFK"}, {"JFK", "AUA"}, {"ANU", "ADL"}, {"AXA", "TIA"}, {"ANU", "AUA"}, {"JFK", "EZE"}, {"AXA", "ADL"}, {"TIA", "EZE"}, {"JFK", "AXA"}, {"AXA", "ADL"}, {"EZE", "AUA"}, {"AXA", "ANU"}, {"ADL", "EZE"}, {"AUA", "EZE"}}))
}
func TestCompare(t *testing.T) {
	fmt.Println(strings.Compare("A", "B"))
	fmt.Println(strings.Compare("JFK", "SFO"))
	arr := []string{"JFK", "SFO"}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
	arr = []string{"SFO", "JFK"}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
}
