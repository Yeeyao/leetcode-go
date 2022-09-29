package link_list

import "fmt"

/*
给定单向链表头 head 以及 left, right 位置 其中 left <= right 将 left 到 right 之间的部分节点反转然后返回反转后的链表头
*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right || head == nil {
		return head
	}
	// 返回的虚拟节点，Next 指向返回的链表的 head
	retHead := &ListNode{}
	retHead.Next = head
	// 反转部分开头的虚拟节点，一开始是指向反转部分的第一个节点，最后需要指向反转部分的最后节点
	reversePre := &ListNode{}
	reversePre.Next = head
	tempLeft := left
	for head != nil && tempLeft-1 > 0 {
		tempLeft--
		reversePre = head
		head = head.Next
	}
	// 反转部分的开头节点，最后需要指向反转部分结束的下一个节点
	reverseHead := head
	// 中间部分反转，反转完之后 midPrev 的位置是反转部分的最后的节点，midCur 是下一个节点
	midPrev := &ListNode{}
	midCur := head
	for midCur != nil && right > left-1 {
		right--
		temp := midCur.Next
		midCur.Next = midPrev
		midPrev = midCur
		midCur = temp
	}
	// 如果 left 是 1 则 retHead.Next 需要指向的位置需要变化
	if left == 1 {
		retHead.Next = midPrev
	} else {
		reversePre.Next = midPrev
	}
	reverseHead.Next = midCur
	return retHead.Next
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}
	// 需要找到四个位置
	// 先找到反转部分的开始的位置和前一个位置
	retHead := &ListNode{}
	retHead.Next = head
	reversePre := &ListNode{}
	reversePre.Next = head
	tempLeft := left
	for head != nil && tempLeft > 0 {
		if tempLeft > 1 {
			reversePre = head
		}
		head = head.Next
		tempLeft--
	}
	fmt.Println("rp: ", reversePre.Val)
	fmt.Println("head: ", head.Val)
	// 这里从反转部分的第二个节点开始反转
	midPrev := reversePre.Next
	midCur := head
	fmt.Println("midPrev: ", midPrev.Val)
	fmt.Println("midCur: ", midCur.Val)
	for midCur != nil && right > left {
		temp := midCur.Next
		midCur.Next = midPrev
		midPrev = midCur
		midCur = temp
		right--
	}
	fmt.Println("midPrev: ", midPrev.Val)
	fmt.Println("midCur: ", midCur.Val)
	// 重新连接
	if left == 1 {
		retHead.Next = midPrev
		fmt.Println("retHead.Next: ", retHead.Next.Val)
		reversePre.Next = midCur
		fmt.Println("reversePre.Next: ", reversePre.Next.Val)
	} else {
		reversePre.Next.Next = midCur
		reversePre.Next = midPrev
	}
	return retHead.Next
}
