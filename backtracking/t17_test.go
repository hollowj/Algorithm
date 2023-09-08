package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var numMap = map[string][]byte{
	"1": {},
	"2": {'a', 'b', 'c'},
	"3": {'d', 'e', 'f'},
	"4": {'g', 'h', 'i'},
	"5": {'j', 'k', 'l'},
	"6": {'m', 'n', 'o'},
	"7": {'p', 'q', 'r', 's'},
	"8": {'t', 'u', 'v'},
	"9": {'w', 'x', 'y', 'z'},
}

func GetChars(numStr string) []byte {
	return numMap[numStr]
}

var results17 []string
var result17 []byte

func letterCombinations(digits string) []string {
	results17 = make([]string, 0)
	dfs17(digits, 0)
	return results17
}
func dfs17(digits string, begin int) {
	if len(result17) == len(digits) {
		if len(result17) > 0 {
			results17 = append(results17, string(result17))
		}
		return
	}
	for _, ch := range GetChars(string(digits[begin])) {
		result17 = append(result17, ch)
		dfs17(digits, begin+1)
		result17 = result17[:len(result17)-1]
	}

}
func TestT17(t *testing.T) {
	assert.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
	assert.Equal(t, []string{}, letterCombinations(""))
	assert.Equal(t, []string{"a", "b", "c"}, letterCombinations("2"))
	assert.Equal(t, []string{"p", "q", "r", "s"}, letterCombinations("7"))
}
