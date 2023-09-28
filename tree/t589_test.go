package tree

//
//type Node struct {
//	Val      int
//	Children []*Node
//}
//
//var results []int
//
//func preorder(root *Node) []int {
//	results = make([]int, 0)
//	preorder1(root)
//	return results
//}
//func preorder1(root *Node) {
//	if root == nil {
//		return
//	}
//	results = append(results, root.Val)
//	for _, node := range root.Children {
//		preorder1(node)
//	}
//
//}
//func TestT589(t *testing.T) {
//	assert.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, levelOrder(buildTree([]int{3, 9, 20, math.MinInt, math.MinInt, 15, 7})))
//	assert.Equal(t, [][]int{{1}}, levelOrder(buildTree([]int{1})))
//	assert.Equal(t, [][]int{}, levelOrder(buildTree([]int{})))
//}
