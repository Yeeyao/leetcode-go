package link_list

/*
	给定一个已排序的链表，删除所有的重复元素，只要是重复的元素都删除。返回的链表需要也是已排序的
	类似 83 但是需要删除重复的元素
*/

/*
	这里可能需要删除全部的元素
	当前节点需要不断向后判断，如果下一个元素和该节点相等就需要不断删除该数值的节点
	这里 head 也可以被删除
	好多逻辑
*/

/*
	prev, cur 判断 cur 和 cur.Next 的数值
	- 不相等则 prev, cur 一起前进
	- 如果相等，则需要 cur 一直前进到和当前数值( prev.Next.Val )不同的位置，因为 cur 是一直变化的，不是当前的数值
	同时需要注意边界条件的判断
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fHead := &ListNode{}
	fHead.Next = head
	prev := fHead
	cur := fHead.Next
	for cur != nil && cur.Next != nil {
		// 这里不成立，则 cur 和下一个节点是一定数值相同
		if cur.Val != cur.Next.Val {
			prev = prev.Next
			cur = cur.Next
			continue
		}
		// 跳过重复数值的节点，注意当前的数值是 prev.Next.Val 的数值(这里记录相同数值节点的开头位置)
		for cur != nil && prev.Next.Val == cur.Val {
			cur = cur.Next
		}
		prev.Next = cur
	}
	return fHead.Next
}

// 自己的 需要时间，不够优雅
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 需要虚拟头
	fHead := &ListNode{}
	prev := fHead
	temp := fHead
	cur := head

	curValCount := 0
	curVal := cur.Val
	for cur != nil {
		if cur.Val == curVal {
			curValCount++
		} else {
			curVal = cur.Val
			if curValCount == 1 {
				temp.Next = prev
				temp = prev
				// 同样需要先断掉和后面的节点的连接
				temp.Next = nil
			} else {
				curValCount = 1
			}
		}
		prev = cur
		cur = cur.Next
	}
	// 最后一个节点需要另外处理
	if curValCount == 1 {
		temp.Next = prev
		temp = prev
		temp.Next = nil
	}
	return fHead.Next
}
