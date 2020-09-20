package math

/*
445. Add Two Numbers II 类似 leetcode 2
You are given two non-empty linked lists representing two non-negative integers.
The most significant digit comes first and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

What if you cannot modify the input lists? In other words, reversing the lists is not allowed.
给定两个数字，求它们的和，数字由链表组成，最高有效位在链表头部
直接递归处理可以，或者使用额外的数组保存元素，但是递归怎么处理
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 两个临时数组保存链表的数据
	s1, s2 := make([]int, 0), make([]int, 0)
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	s1Len, s2Len := len(s1), len(s2)
	longer := 0
	if s1Len > s2Len {
		longer = s1Len
	} else {
		longer = s2Len
	}
	// 中间结果数组，这里从前面两个数组的尾部向前面计算
	res := make([]int, 0)
	carry := 0
	for i := 0; i < longer; i++ {
		if i < s1Len {
			carry += s1[s1Len-i-1]
		}
		if i < s2Len {
			carry += s2[s2Len-i-1]
		}
		res = append(res, carry%10)
		carry /= 10
	}
	if carry > 0 {
		res = append(res, carry)
	}
	// 这里避免判断 nil
	head := &ListNode{}
	cur := head
	for i := len(res) - 1; i >= 0; i-- {
		cur.Next = &ListNode{res[i], nil}
		cur = cur.Next
	}
	return head.Next
}
