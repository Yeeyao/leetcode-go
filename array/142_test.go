package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("142. Linked List Cycle II", func(t *testing.T) {
		input := []*ListNode{1, 3, 4, 2, 2}
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	类似 287 题目没有说明一定有循环，所以需要判断
*/
func solution(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	isCycle := false

	for slow != nil && fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			isCycle = true
			break
		}
	}
	// 没有循环直接返回 nil
	if !isCycle {
		return nil
	}
	// 有循环，找开始的节点
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
