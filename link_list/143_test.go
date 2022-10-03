package link_list

/*
重新排列链表

L0 L1 ... Ln-1 Ln 变成 L0 Ln L1 Ln-1

这里需要找到后面的节点 看起来可以先遍历一次，然后将后半部分的节点都保存起来

-  因为没有办法知道链表的长度，因此需要全部都先保存一次

保存可以使用数组，然后重新遍历链表进行连接
*/
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	var nodeArray []*ListNode
	// 第一次先保存
	cur := head
	for cur != nil {
		nodeArray = append(nodeArray, cur)
		cur = cur.Next
	}
	// 重新组装
	begin, end := 0, len(nodeArray)-1
	for begin < end {
		temp := nodeArray[begin].Next
		nodeArray[begin].Next = nodeArray[end]
		nodeArray[end].Next = temp
		begin++
		end--
	}
	// 断开中间或者中间后一个到下一个的连接
	nodeArray[begin].Next = nil
}
