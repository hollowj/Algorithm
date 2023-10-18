package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func ladderLength(beginWord string, endWord string, wordList []string) int {
//	set := make(map[string]struct{})
//	for _, word := range wordList {
//		set[word] = struct{}{}
//	}
//	pathMap := make(map[string]int)
//	queue := make([]string, 0)
//	queue = append(queue, beginWord)
//	pathMap[beginWord] = 1
//	for len(queue) != 0 {
//		str := queue[0]
//		queue = queue[1:]
//		iPath := pathMap[str]
//		for i := 0; i < len(str); i++ {
//			for j := 0; j < 26; j++ {
//				s := []rune(str)
//				s[i] = rune('a' + j)
//				newStr := string(s)
//				if _, ok := set[newStr]; ok {
//					if newStr == endWord {
//						return iPath + 1
//					}
//					if _, ok2 := pathMap[newStr]; !ok2 {
//						pathMap[newStr] = iPath + 1
//						queue = append(queue, newStr)
//					}
//				}
//			}
//		}
//	}
//	return 0
//}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	graph := make([][]int, 0)
	wordIdMap := make(map[string]int)
	addWord := func(str string) int {
		var wordId int
		if tmpId, ok := wordIdMap[str]; !ok {
			wordId = len(wordIdMap)
			wordIdMap[str] = wordId
			graph = append(graph, []int{})
		} else {
			wordId = tmpId
		}
		return wordId
	}
	addEdge := func(str string) {
		rawWordId := addWord(str)
		for i := 0; i < len(str); i++ {
			runes := []rune(str)
			runes[i] = '*'
			newStr := string(runes)
			wordId := addWord(newStr)
			graph[rawWordId] = append(graph[rawWordId], wordId)
			graph[wordId] = append(graph[wordId], rawWordId)
		}
	}
	addEdge(beginWord)
	for _, word := range wordList {
		addEdge(word)
	}
	endWordId, ok := wordIdMap[endWord]
	if !ok {
		return 0
	}
	pathMap := make(map[int]int)
	beginWordId := addWord(beginWord)
	pathMap[beginWordId] = 1
	queue := []int{beginWordId}
	for len(queue) != 0 {
		wordId := queue[0]
		queue = queue[1:]
		targets := graph[wordId]
		iPath := pathMap[wordId]
		for _, targetId := range targets {
			if targetId == endWordId {
				return iPath/2 + 1
			}
			if _, ok := pathMap[targetId]; !ok {
				pathMap[targetId] = iPath + 1
				queue = append(queue, targetId)
			}
		}
	}
	return 0
}
func TestT127(t *testing.T) {
	assert.Equal(t, 5, ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	assert.Equal(t, 0, ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
