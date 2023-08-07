package tanxin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lemonadeChange(bills []int) bool {
	count10 := 0
	count5 := 0
	for _, tmp := range bills {
		switch tmp {
		case 5:
			count5++
		case 10:
			count10++
			count5--
		case 20:
			if count10 > 0 {
				count10--
				count5--
			} else {
				count5 -= 3
			}
		}
		if count5 < 0 {
			return false
		}
	}
	return true
}

func TestT860(t *testing.T) {
	assert.Equal(t, true, lemonadeChange([]int{5, 5, 5, 10, 20}))
	assert.Equal(t, false, lemonadeChange([]int{5, 5, 10, 10, 20}))
	assert.Equal(t, true, lemonadeChange([]int{5, 5, 5, 20}))
}
