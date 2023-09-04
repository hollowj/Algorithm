package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func strStr(haystack string, needle string) int {
//		slow := 0
//		iLen := 0
//		for slow < len(haystack) {
//			fast := slow
//			for fast < len(haystack) && haystack[fast] == needle[iLen] {
//				iLen++
//				fast++
//				if iLen == len(needle) {
//					return slow
//				}
//			}
//			iLen = 0
//			slow++
//		}
//
//		return -1
//	}

func getNext(pattern string) []int {
	next := make([]int, len(pattern))
	j := 0
	next[0] = 0
	for i := 1; i < len(pattern); i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = next[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		next[i] = j
	}
	return next
}
func strStr(haystack string, needle string) int {
	next := getNext(needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
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
}
