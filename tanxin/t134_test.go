package tanxin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func canCompleteCircuit(gas []int, cost []int) int {
//	iRemain := 0
//	iMin := math.MaxInt
//	for i := 0; i < len(gas); i++ {
//		rest := gas[i] - cost[i]
//		iRemain += rest
//		if iRemain < iMin {
//			iMin = iRemain
//		}
//	}
//	if iRemain < 0 {
//		return -1
//	}
//	if iMin >= 0 {
//		return 0
//	}
//	for i := len(gas) - 1; i > 0; i-- {
//		rest := gas[i] - cost[i]
//		iMin += rest
//		if iMin >= 0 {
//			return i
//		}
//	}
//	return -1
//}

func canCompleteCircuit(gas []int, cost []int) int {
	iRemain := 0
	iTotalSum := 0
	beginIndex := 0
	for i := 0; i < len(gas); i++ {
		rest := gas[i] - cost[i]
		iRemain += rest
		iTotalSum += rest
		if iRemain < 0 {
			iRemain = 0
			beginIndex = i + 1
		}
	}
	if iTotalSum < 0 {
		return -1
	}
	return beginIndex
}

func TestT134(t *testing.T) {
	assert.Equal(t, 3, canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	assert.Equal(t, -1, canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
