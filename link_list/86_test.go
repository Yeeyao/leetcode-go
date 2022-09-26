package link_list

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	划分链表：给定一个链表头和数值 x，将链表划分，其中节点数值小于 x 的节点需要在大于等于 x 的节点前面
	需要保留节点原来的相对数值
*/
func TestPro(t *testing.T) {
	t.Run("86. test", func(t *testing.T) {
		head := &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		}
		retNode := partition(head, 2)
		fmt.Println(retNode.Val)
	})
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	lessHead := &ListNode{}
	lessCur := lessHead
	greaterHead := &ListNode{}
	greaterCur := greaterHead
	for head != nil {
		if head.Val < x {
			lessCur.Next = head
			lessCur = lessCur.Next
		} else {
			greaterCur.Next = head
			greaterCur = greaterCur.Next
		}
		head = head.Next
	}
	lessCur.Next = greaterHead.Next
	// 断开避免环
	greaterCur.Next = nil
	return lessHead.Next
}
