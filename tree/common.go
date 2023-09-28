package tree

import "math"

const null = math.MinInt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
