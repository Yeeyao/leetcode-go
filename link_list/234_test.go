package link_list

/*
Given a singly linked list, determine if it is a palindrome.
O(n) 时间 O(1) 空间
给定单向链表，判断它是否是回文
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用快慢指针找到中间的位置，然后将前面的部分进行反转，最后比较前后两部分
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	// 节点数量是奇数，slow 刚好在中间，是偶数，则 slow 在中间的下一个位置
	reverseHead := &ListNode{}
	for slow != nil && fast != nil && fast.Next != nil {
		reverseHead.Next = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 前面的部分进行反转
	var prev *ListNode
	cur := head
	for cur != slow {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	// 节点数量是奇数，则从反转的开头和 slow 节点的下一个节点开始比较
	// 节点数量是偶数，则从反转的开头和 slow 节点开始比较
	// fast 非 nil 表示节点数量是奇数
	if fast != nil {
		front := reverseHead.Next
		for front != nil {
			if front.Val != slow.Next.Val {
				return false
			}
			front = front.Next
			slow = slow.Next
		}
	} else {
		front := reverseHead.Next
		for front != nil {
			if front.Val != slow.Val {
				return false
			}
			front = front.Next
			slow = slow.Next
		}
	}
	return true
}

/*
	直接将前面的部分翻转，然后从中间向后和从开头到中间逐个比较
	在找到中间部分的同时将前面部分反转，比较前后两个部分
	rev, fast, slow
	rev 是一个反转链表，之前前面，然后 fast 一次前进两步，slow 一次前进一步
	需要注意，如果链表元素数量为奇数，则 slow 需要向前一步
	最后比较 rev 和 slow，然后判断 rev 是否为 nil

	head 或者 head.Next 为 nil 就返回 true判断
	head, slow, fast 分别初始化为 nil， head，head
	第一个循环条件是 fast, fast.Next 非 nil
		head = slow, slow.Next = head, slow = slow.Next fast = fast.Next.Next
	中间判断 fast 非 nil 表示链表元素个数为奇数，slow 向后一步
	第二个循环条件是 head 非 nil 且 head 与 slow 的数值需要相同
		head 和 slow 都向后一步
	最后返回 head == nil
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	head, slow, fast := nil, head, head
	for fast != nil && fast.Next != nil {
		head, slow.Next, slow, fast = slow, head, slow.Next, fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	for head != nil && head.Val == slow.Val {
		head, slow = head.Next, slow.Next
	}
	return head == nil
}
