package link_list

/*
	给定一个已排序的链表，重复的元素只能留下一个。返回的链表需要也是已排序的
	合并的类似思路，只不过是判断下一个元素是否和当前元素数值相同
*/

// 最快 前后指针的思想，类似数组，数组是当前索引和后面的索引进行追加。 tmp 是需要判断的节点
// 看起来如果一定需要设置第一个元素为开头就可能是这种方法
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prev := head
	cur := head.Next
	for cur != nil {
		if cur.Val != prev.Val {
			prev.Next = cur
			prev = prev.Next
		}
		cur = cur.Next
	}
	// 断开最后的元素都是相同的情况下和后面的重复元素的连接
	prev.Next = nil
	return head
}

// 自己
func deleteDuplicates(head *ListNode) *ListNode {
	var fakeHead = &ListNode{}
	cur := fakeHead
	// 第一个元素需要直接处理，这里如果尾部全是相同的，需要断开和这些相同的元素的连接，因此需要将 cur.Next 设置为 nil
	if head == nil {
		return head
	}
	cur.Next = head
	cur = cur.Next
	head = head.Next
	cur.Next = nil
	for head != nil {
		if cur.Val != head.Val {
			cur.Next = head
			cur = cur.Next
			head = head.Next
			cur.Next = nil
		} else {
			head = head.Next
		}
	}
	return fakeHead.Next
}
