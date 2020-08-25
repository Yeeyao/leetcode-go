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
	for fast != nil  && fast.Next != nil {
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

/*
	这个不行是因为
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	rev := &ListNode{}
	slow, fast := head, head
	for fast != nil  && fast.Next != nil {
		fast = fast.Next.Next
		rev, rev.Next, slow = slow, rev, slow.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	for rev != nil && rev.Val == slow.Val {
		rev = rev.Next
		slow = slow.Next
	}
	return rev == nil 
}
