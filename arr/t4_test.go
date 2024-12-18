package arr

//
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	total := len(nums1) + len(nums2)
//	var left, right, leftValue, rightValue int
//	if total%2 == 0 {
//		left = total / 2
//		right = left
//	} else {
//		left = total/2 + 1
//		right = total/2 + 1
//	}
//	index1, index2 := 0, 0
//	for index1 < len(nums1) && index2 < len(nums2) {
//		var v int
//		if nums1[index1] <= nums2[index2] {
//			v = nums1[index1]
//			index1++
//		} else {
//			v = nums2[index2]
//			index2++
//		}
//		if index1+index2 == left {
//			leftValue = v
//		}
//
//	}
//	return float64(leftValue+rightValue) / 2
//
//}
//func TestT4(t *testing.T) {
//	assert.Equal(t, findMedianSortedArrays([]int{1, 3}, []int{2}), 2)
//	assert.Equal(t, findMedianSortedArrays([]int{1, 2}, []int{3, 4}), 2.5)
//}
