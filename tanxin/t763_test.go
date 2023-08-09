package tanxin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func partitionLabels(s string) []int {
	count := make([]int, 27)
	for i, ch := range s {
		count[ch-'a'] = i
	}
	result := make([]int, 0)
	right := 0
	lastIndex := 0

	for i, ch := range s {
		right = max(right, int(count[ch-'a']))
		if i == right {
			result = append(result, i-lastIndex+1)
			lastIndex = i + 1
		}
	}
	return result
}

func TestT763(t *testing.T) {
	assert.Equal(t, []int{9, 7, 8}, partitionLabels("ababcbacadefegdehijhklij"))
	assert.Equal(t, []int{10}, partitionLabels("eccbbbbdec"))
	assert.Equal(t, []int{246, 254}, partitionLabels("ntswuqqbidunnixxpoxxuuupotaatwdainsotwvpxpsdvdbwvbtdiptwtxnnbtqbdvnbowqitudutpsxsbbsvtipibqpvpnivottsxvoqqaqdxiviidivndvdtbvadnxboiqivpusuxaaqnqaobutdbpiosuitdnopoboivopaapadvqwwnnwvxndpxbapixaspwxxxvppoptqxitsvaaawxwaxtbxuixsoxoqdtopqqivaitnpvutzchkygjjgjkcfzjzrkmyerhgkglcyffezmehjcllmlrjghhfkfylkgyhyjfmljkzglkklykrjgrmzjyeyzrrkymccefggczrjflykclfhrjjckjlmglrmgfzlkkhffkjrkyfhegyykrzgjzcgjhkzzmzyejycfrkkekmhzjgggrmchkeclljlyhjkchmhjlehhejjyccyegzrcrerfzczfelzrlfylzleefgefgmzzlggmejjjygehmrczmkrc"))
}
