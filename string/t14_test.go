package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestCommonPrefix(strs []string) string {
	for index := 0; index < len(strs[0]); index++ {
		for i := 1; i < len(strs); i++ {
			if index >= len(strs[i]) || strs[0][index] != strs[i][index] {
				return strs[0][:index]
			}
		}

	}
	return strs[0]
}
func TestT14(t *testing.T) {
	assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	assert.Equal(t, "", longestCommonPrefix([]string{""}))
}
