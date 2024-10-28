// func pushBack(list *ListNode, value int) {
// 	cur := list
// 	for cur.Next != nil {
// 		cur = cur.Next
// 	}
// 	cur.Next = &ListNode{Val: value}
// }

// func genSlice(list *ListNode) []int {
// 	var slice []int
// 	cur := list
// 	for cur != nil {
// 		slice = append(slice, cur.Val)
// 		cur = cur.Next
// 	}
// 	return slice
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	sliceFirst := genSlice(l1)
// 	sliceSecond := genSlice(l2)
// 	var resultSlice []int
// 	var ost = 0

// 	for i:=0;i<len(sliceFirst) || i<len(sliceSecond) || ost!=0;i++{

// 		if i<len(sliceFirst) && i<len(sliceSecond){
// 			resultSlice = append(resultSlice,(sliceFirst[i] + sliceSecond[i] + ost)%10)
// 			ost = (sliceFirst[i] + sliceSecond[i] + ost) / 10
// 		} else if i<len(sliceFirst){
// 			resultSlice = append(resultSlice,(sliceFirst[i] + ost)%10)
// 			ost = (sliceFirst[i] + ost) / 10
// 		} else if i<len(sliceSecond){
// 			resultSlice = append(resultSlice,(sliceSecond[i] + ost)%10)
// 			ost = (sliceSecond[i] + ost) / 10
// 		} else if ost!=0{
// 			resultSlice = append(resultSlice, ost)
// 			ost = 0
// 		}
// 	}

// 	var result *ListNode
// 	if len(resultSlice) > 0 {
// 		result = &ListNode{Val: resultSlice[0]}
// 		for i := 1; i < len(resultSlice); i++ {
// 			pushBack(result, resultSlice[i])
// 		}
// 	}
// 	return result
// }
