package kmp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getNext(str string) []int {
	next := make([]int, len(str))
	k := 0
	next[0] = k
	for i := 1; i < len(str); i++ {
		for k > 0 && str[i] != str[k] {
			k = next[k-1]
		}
		if str[i] == str[k] {
			k++
		}
		next[i] = k
	}
	return next

}
func strStr(haystack string, needle string) int {
	next := getNext(needle)
	k := 0
	for i := 0; i < len(haystack); i++ {
		for k > 0 && haystack[i] != needle[k] {
			k = next[k-1]
		}
		if haystack[i] == needle[k] {
			k++
			if k == len(needle) {
				return i - k + 1
			}
		}
	}
	return -1
}
func TestT28(t *testing.T) {
	str := "abeab"
	next := getNext(str)
	fmt.Println(str, next)
	assert.Equal(t, 0, strStr("sadbutsad", "sad"))
	assert.Equal(t, -1, strStr("leetcode", "leeto"))
	assert.Equal(t, 4, strStr("mississippi", "issip"))
	assert.Equal(t, -1, strStr("aaa", "aaaa"))
	assert.Equal(t, -1, strStr("babba", "bbb"))
}
