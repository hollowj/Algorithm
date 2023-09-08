package backtracking

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BackTracking91 struct {
	results []string
	result  []string
}

func NewBackTracking91() *BackTracking91 {
	return &BackTracking91{results: make([]string, 0), result: make([]string, 0)}
}
func (b *BackTracking91) dfs(s string, begin int) {
	if len(b.result) == 4 {
		if begin != len(s) {
			return
		}
		b.results = append(b.results, strings.Join(b.result, "."))
		return
	}
	for i := begin; i < len(s); i++ {
		if b.isAvaliableIP(s[begin : i+1]) {
			b.result = append(b.result, s[begin:i+1])
			b.dfs(s, i+1)
			b.result = b.result[:len(b.result)-1]
		}
	}
}
func (b *BackTracking91) isAvaliableIP(s string) bool {
	if len(s) > 1 && strings.HasPrefix(s, "0") {
		return false
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return num <= 255
}
func restoreIpAddresses(s string) []string {
	tracking91 := NewBackTracking91()
	tracking91.dfs(s, 0)
	return tracking91.results
}
func TestT93(t *testing.T) {
	assert.Equal(t, []string{"255.255.11.135", "255.255.111.35"}, restoreIpAddresses("25525511135"))
	assert.Equal(t, []string{"0.0.0.0"}, restoreIpAddresses("0000"))
	assert.Equal(t, []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}, restoreIpAddresses("101023"))
}
