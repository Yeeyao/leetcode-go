package link_list

/*
	给定一个链表，交换每两个相邻的节点然后返回链表头
	两个相邻的节点才需要交换，类似反转链表的部分节点

*/
// 最快，使用递归 这个可以学
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ptr := head.Next
	ptr_Next := ptr.Next
	ptr.Next = head
	head.Next = swapPairs(ptr_Next)
	return ptr
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 单数的节点需要指向下一对的节点
	// 双数的节点需要反转
	// 最终的结果节点的前一个，只需要设置一次
	retHead := &ListNode{}
	retHead.Next = head
	if head.Next != nil {
		retHead.Next = head.Next
	}
	// 交换的两个节点
	prev := &ListNode{}
	cur := head
	// 当前交换对的前一对节点的第一个节点，因为不确定后面的节点是一个还是两个，所以后面一对节点的两个节点都需要被这个节点指向一次
	// 在偶数的时候更新，因为偶数更新的时候，确定这个节点是当前偶数的前一个，然后遍历到下一对的奇数时可以直接更新这个节点的下一个
	var prevSingle *ListNode
	posCount := 0
	for cur != nil {
		posCount++
		if posCount%2 != 0 {
			// 奇数需要将上一对的奇数节点 Next 指向当前节点
			if prevSingle != nil {
				prevSingle.Next = cur
			}
			temp := cur.Next
			prev = cur
			cur = temp
			// 只有偶数的时候需要反转
		} else {
			// 这里需要先上一对的奇数节点 Next 指向当前节点，然后才更新
			if prevSingle != nil {
				prevSingle.Next = cur
			}
			prevSingle = prev
			// 最后一对元素的奇数位置需要断开和下一个的连接
			prev.Next = nil
			temp := cur.Next
			cur.Next = prev
			prev = cur
			cur = temp
		}
	}
	return retHead.Next
}
