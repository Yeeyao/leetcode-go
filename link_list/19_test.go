package link_list

/*
19. Remove Nth Node From End of List
Given a linked list, remove the n-th node from the end of list and return its head.
给定一个链表，移除从尾部到前面的第 n 个节点并返回链表头
使用两个指针 第一个指针先跑 然后第二个指针一起跑到链表尾部
 1->2->3->4->5, and n = 2. 1->2->3->5.
 1->2->3->4->5, and n = 5. 2->3->4->5.

创建一个新的 head 指向原有的 head，第一个指针向后遍历 n 次
第二个指针从新的 head 开始和第一个指针一同向后遍历

创建新的 head 指向原有的 head
第一个循环 n 次 将第一个指针向后移动
第二个循环条件是第一个指针的 next 非 nil
	两个指针一起移动
最后将第二个指针指向后一个 返回新 head 的 next
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{}
	newHead.Next = head
	first, second := newHead, newHead
	for ; n > 0; n-- {
		first = first.Next
	}
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return newHead.Next
}
