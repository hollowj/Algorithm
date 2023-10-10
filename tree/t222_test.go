package tree

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func countNodes(root *TreeNode) int {
//		lastFloorNodes := []*TreeNode{}
//		if root != nil {
//			lastFloorNodes = append(lastFloorNodes, root)
//		}
//		iCount := 0
//		for len(lastFloorNodes) != 0 {
//			tmp := make([]*TreeNode, 0)
//			for _, node := range lastFloorNodes {
//				iCount++
//				if node.Left != nil {
//					tmp = append(tmp, node.Left)
//				}
//				if node.Right != nil {
//					tmp = append(tmp, node.Right)
//				}
//			}
//			lastFloorNodes = tmp
//		}
//		return iCount
//	}
//
//	func countNodes(root *TreeNode) int {
//		if root == nil {
//			return 0
//		}
//		return countNodes(root.Left) + countNodes(root.Right) + 1
//	}
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ok, length := isFullTree(root)
	if ok {
		return int(math.Pow(2, float64(length))) - 1
	} else {
		return countNodes(root.Left) + countNodes(root.Right) + 1
	}
}
func isFullTree(root *TreeNode) (bool, int) {
	cur := root
	iLeftLength := 0
	for cur != nil {
		iLeftLength++
		cur = cur.Left
	}
	iRightLength := 0
	cur = root
	for cur != nil {
		iRightLength++
		cur = cur.Right
	}
	if iLeftLength == iRightLength {
		return true, iLeftLength
	}
	return false, 0
}
func TestT222(t *testing.T) {
	assert.Equal(t, 6, countNodes(buildTree([]int{1, 2, 3, 4, 5, 6})))
	assert.Equal(t, 0, countNodes(buildTree([]int{})))
	assert.Equal(t, 1, countNodes(buildTree([]int{1})))
}

func TestPow(t *testing.T) {
	fmt.Println(2 << 1) //==2*2
	fmt.Println(2 << 2) //==2*2*2
	fmt.Println(3 << 1) //3*2
	fmt.Println(3 << 2) //3*2*2
	fmt.Println(2 >> 1) //2/2
	fmt.Println(2 >> 2) //2/2/2
}
