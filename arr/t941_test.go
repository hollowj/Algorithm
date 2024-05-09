package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	up := true
	for i := 0; i < len(arr); i++ {
		if up {
			if i+1 < len(arr)-1 && arr[i] > arr[i+1] {
				up = false

			} else if i+1 < len(arr)-1 && arr[i] == arr[i+1] {
				return false
			}
		} else {
			if arr[i] >= arr[i-1] {
				return false
			}
		}

	}
	return true
}
func TestT941(t *testing.T) {
	assert.Equal(t, false, validMountainArray([]int{2, 1}))
	assert.Equal(t, false, validMountainArray([]int{3, 5, 5}))
	assert.Equal(t, true, validMountainArray([]int{0, 3, 2, 1}))
}
