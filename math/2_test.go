package math

/*
2. Add Two Numbers
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
[ref](https://leetcode-cn.com/problems/add-two-numbers/solution/liang-shu-xiang-jia-by-leetcode/)
给定两个非空的链表表示两个非负的整型，每个节点是数字的二进制表示，其中表头是最低位，没有前导 0
返回两个数字的链表和

newHead 初始化为空 ListNode
p,q 变量初始化为 l1, l2 初始化为新 ListNode, cur 初始化为  newHead carry 进位初始化为 0
循环判断 p 或者 q 非 nil
	x, y 分别从 p，q 当前节点取值，若为 nil 就取 0
	sum 计算总和+进位
	cur.Next 先用总和%10初始化，然后用它更新 cur
	p,q 非 nil 就向前
循环结束后，如果 carry 还是 1 就需要添加最后一个节点
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{}
	p, q, cur, carry := l1, l2, newHead, 0
	for p != nil || q != nil {
		x, y := 0, 0
		if p != nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}
		sum := carry + x + y
		carry = sum / 10
		cur.Next = &ListNode{sum % 10, nil}
		cur = cur.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		cur.Next = &ListNode{carry, nil}
	}
	return newHead.Next
}

// 代码太啰嗦了
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	res := &ListNode{}
	newHead := res
	addOn := 0
	for l1 != nil && l2 != nil {
		add := l1.Val + l2.Val + addOn
		if add >= 10 {
			res.Next = &ListNode{Val: add - 10}
			addOn = 1
		} else {
			res.Next = &ListNode{Val: add}
		}
		l1, l2, res = l1.Next, l2.Next, res.Next
	}
	if l1 != nil {
		if addOn == 0 {
			res.Next = l1
		} else {
			for l1 != nil {
				add := l1.Val + addOn
				if add >= 10 {
					res.Next = &ListNode{Val: add - 10}
					addOn = 1
				} else {
					res.Next = l1
					break
				}
				l1 = l1.Next
			}
		}
	}
	if l2 != nil {
		if addOn == 0 {
			res.Next = l2
		} else {
			for l2 != nil {
				add := l2.Val + addOn
				if add >= 10 {
					res.Next = &ListNode{Val: add - 10}
					addOn = 1
				} else {
					res.Next = l2
					break
				}
				l2 = l2.Next
			}
		}
	}
	if addOn > 0 {
		res.Next = &ListNode{addOn, nil}
	}
	return newHead.Next
}
