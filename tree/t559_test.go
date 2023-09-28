package tree

//	type Node struct {
//		Val      int
//		Children []*Node
//	}
//
//	func maxDepth(root *Node) int {
//		nextFloor := []*Node{}
//		if root != nil {
//			nextFloor = append(nextFloor, root)
//		}
//		iDepth := 0
//		for len(nextFloor) != 0 {
//			iDepth++
//			tmp := make([]*Node, 0)
//			for _, node := range nextFloor {
//				for _, child := range node.Children {
//					tmp = append(tmp, child)
//
//				}
//			}
//			nextFloor = tmp
//
//		}
//
//		return iDepth
//	}
//func maxDepth(root *Node) int {
//	if root == nil {
//		return 0
//	}
//	iMaxDepth := 1
//	for _, child := range root.Children {
//		iDepth := maxDepth(child) + 1
//		if iDepth > iMaxDepth {
//			iMaxDepth = iDepth
//		}
//
//	}
//
//	return iMaxDepth
//}
//func TestT559(t *testing.T) {
//	assert.Equal(t, 3, maxDepth(buildTree([]int{3, 9, 20, null, null, 15, 7})))
//	assert.Equal(t, 2, maxDepth(buildTree([]int{1, null, 2})))
//}
