package link_list

/*
Write a program to find the node at which the intersection of two singly linked lists begins.
找到两个链表的相交节点
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	这里直接将两个链表都遍历，然后遍历到其中一个的尾部就从头部重新开始遍历
	如果两个开头有一个是 nil 则直接返回 nil
	但是这里怎么判断是否相交 使用两个迭代来处理
		1.当其中一个指针到达结尾的时候，将它重置为**另一个链表的开头**
		2.持续移动两个指针知道它们指向相同的节点
	当两个链表相交，第二个迭代将指向交点，不相交，第二个迭代中的节点将都是 null

	解释
		a, b 为链表 A，B 相交节点前的长度，c 是相交节点后的长度
		List A 长度是 a + c List B 长度是 b + c A 或者 B 遍历完了，跑到另一个链表遍历
			有相交节点则最后 A 遍历 a + c + b + c B 遍历 b + c + a + c，
			没有相交节点则最后 A 遍历 a + b B 遍历 b + a
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
