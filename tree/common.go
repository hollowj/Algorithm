package tree

import "math"

const null = math.MinInt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) FloorArray() []int {
	result := []int{}
	floor := []*TreeNode{t}
	for len(floor) > 0 {
		nextFloor := make([]*TreeNode, 0)
		hasNextFloor := false
		for _, node := range floor {
			if node != nil {
				result = append(result, node.Val)
				if node.Left != nil {
					hasNextFloor = true
				}
				if node.Right != nil {
					hasNextFloor = true
				}
				nextFloor = append(nextFloor, node.Left)
				nextFloor = append(nextFloor, node.Right)
			} else {
				result = append(result, null)
				nextFloor = append(nextFloor, nil)
				nextFloor = append(nextFloor, nil)
			}
		}
		if hasNextFloor {
			floor = nextFloor
		} else {
			floor = nil
		}
	}
	return result
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func buildTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	root := NewTreeNode(arr[0])
	treeNodes := []*TreeNode{root}
	index := 1
	for index < len(arr) {
		node := treeNodes[0]
		if index < len(arr) {
			if arr[index] != null {
				node.Left = NewTreeNode(arr[index])
				treeNodes = append(treeNodes, node.Left)
			}
			index++
		}
		if index < len(arr) && arr[index] != null {
			node.Right = NewTreeNode(arr[index])
			treeNodes = append(treeNodes, node.Right)
		}
		index++
		treeNodes = treeNodes[1:]
	}
	return root
}
