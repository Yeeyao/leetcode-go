package link_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	其中一个为 nil 就返回另外一个
	循环条件是两个链表当前节点非 nil
	将两个数值较小的保存到新节点，然后较小的向后遍历
	最后，判断其中一个链表还有节点就追加到结果链表

	这里使用一个虚拟头方便返回。总体思路是遍历两个链表，较小的就放到当前的下一个直到某个链表为空
	边界情况是一开始某个链表为空就可以直接返回另外一个链表的头
*/
func mergeTwoLIst(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var cur = &ListNode{}
	var fakeHead = cur
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return fakeHead.Next
}

// 为啥不能 画图，这里 head = l1 然后 head = head.Next 错误
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head = &ListNode{}
	var head2 = head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head = l1
			l1 = l1.Next
		} else {
			head = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	if l1 == nil && l2 != nil {
		head = l2
	}
	if l2 == nil && l1 != nil {
		head = l1
	}
	return head2.Next
}
