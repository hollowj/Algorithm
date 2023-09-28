package tree

type Node struct {
	Val      int
	Children []*Node
}

//
//func levelOrder(root *Node) [][]int {
//	results := make([][]int, 0)
//	nextFloor := []*Node{}
//	if root != nil {
//		nextFloor = append(nextFloor, root)
//	}
//	for len(nextFloor) != 0 {
//		result := make([]int, 0)
//		tmp := make([]*Node, 0)
//		for _, node := range nextFloor {
//			result = append(result, node.Val)
//			for _, child := range node.Children {
//				tmp = append(tmp, child)
//			}
//		}
//		results = append(results, result)
//		nextFloor = tmp
//	}
//
//	return results
//}
//
//func TestT429(t *testing.T) {
//	assert.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, levelOrder(buildTree([]int{3, 9, 20, math.MinInt, math.MinInt, 15, 7})))
//	assert.Equal(t, [][]int{{1}}, levelOrder(buildTree([]int{1})))
//	assert.Equal(t, [][]int{}, levelOrder(buildTree([]int{})))
//}
