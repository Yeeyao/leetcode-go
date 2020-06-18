package jianzhi_offer

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("16 合并两个排序的链表", func(t *testing.T) {
		head1 := ListNode{1, &ListNode{2, &ListNode{3, nil}}}
		head2 := ListNode{1, &ListNode{2, &ListNode{3, nil}}}
		head := solution(&head1, &head2)
		want := ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2,
			&ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}}}
		if !reflect.DeepEqual(head, want) {
			t.Errorf("got: %v, want: %v", head, want)
		}
	})
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
	初始化两个指针，如果第一个的数值大于第二个，那需要将第一个指向第二个，否则还是同一个
*/
func solution(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 这里需要另外一个保存初始的头部
	res, temp := &ListNode{}, &ListNode{}
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			res.Next, l2 = l2, l2.Next
		} else {
			res.Next, l1 = l1, l1.Next
		}
	}
	if l1 == nil {
		res.Next = l2
	}
	if l2 == nil {
		res.Next = l1
	}
	return temp.Next
}
