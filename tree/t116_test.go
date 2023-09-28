package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
//type Node struct {
//	Val   int
//	Left  *Node
//	Right *Node
//	Next  *Node
//}
//
//func connect(root *Node) *Node {
//	nextFloor := []*Node{}
//	if root != nil {
//		nextFloor = append(nextFloor, root)
//	}
//	for len(nextFloor) != 0 {
//		tmp := make([]*Node, 0)
//		var pre *Node
//		for _, node := range nextFloor {
//			if pre != nil {
//				pre.Next = node
//			}
//			pre = node
//			if node.Left != nil {
//				tmp = append(tmp, node.Left)
//			}
//			if node.Right != nil {
//				tmp = append(tmp, node.Right)
//			}
//		}
//		nextFloor = tmp
//
//	}
//
//	return root
//}

func TestT116(t *testing.T) {
	assert.Equal(t, []float64{3.00000, 14.50000, 11.00000}, averageOfLevels(buildTree([]int{3, 9, 20, null, null, 15, 7})))
	assert.Equal(t, []float64{3.00000, 14.50000, 11.00000}, averageOfLevels(buildTree([]int{3, 9, 20, 15, 7})))
}
