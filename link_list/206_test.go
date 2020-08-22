package link_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	在 head  前面加上一个 pre，然后使用一个 cur
	循环终止条件是 cur 等于 nil
		pre, cur 一直向后迭代，最后返回 pre
*/
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
