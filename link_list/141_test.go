package link_list

/*
Given a linked list, determine if it has a cycle in it.
To represent a cycle in the given linked list, we use an integer pos which represents
the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.
给定一个链表，判断它是否包含一个环 直接两个指针，一个跑一步一次，一个两步一次，如果有环就会相交
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	if fast != nil && slow != nil {
		// 先向前
		slow = slow.Next
		// 如果快的已经不能前进了
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		// 然后判断
		if fast == slow {
			return true
		}
	}
	return false
}

/*
	先判断当前 head 或者 head.Next 是否 nil
	循环判断条件是两个节点都非 nil
		两个节点先向后遍历，然后判断，遍历前需要判断 nil
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil && slow != nil {
		slow = slow.Next
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		// 两个节点相遇，直接返回 true
		if fast == slow {
			return true
		}
	}
	return false
}

/*
	差不多
	先判断当前 head 或者 head.Next 是否 nil
	初始化 slow = head, fast = head.Next
	循环判断条件是 fast 非 nil 同时 slow != fast
		两个节点先向后遍历，然后判断，遍历前需要判断 nil
	最后返回 fast == slow
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast != nil && slow != fast {
		slow = slow.Next
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
	}
	return fast == slow
}
