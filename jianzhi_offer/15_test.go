package jianzhi_offer

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("15 反转链表", func(t *testing.T) {
		head := ListNode{1, &ListNode{2, &ListNode{3, nil}}}
		solution3(&head)
		want := ListNode{3, &ListNode{2, &ListNode{1, nil}}}
		if !reflect.DeepEqual(head, want) {
			t.Errorf("got: %v, want: %v", head, want)
		}
	})
}

/*
	输入一个链表，反转链表后，输出新链表的表头
	同样可以使用栈来保存，然后一个个出栈并指向上一个栈顶元素
	直接反转也可以
		先判断下一个是否 nil 如果不是，就先保存下一个
		然后将下一个的下一个指向当前，之后继续遍历到倒数第二个，修改完最后一个指向倒数第二个
		将原来的第一个的下一个指向 nil 返回倒数第一个
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func solution(head *ListNode) *ListNode {
	cur, temp, pre := head, &ListNode{}, &ListNode{}
	for cur != nil {
		// 当前的下一个节点
		temp = cur.Next
		// 当前节点指向上一个
		cur.Next = pre
		// 当前和上一个都继续遍历
		pre = cur
		cur = temp
	}
	return pre
}

/*
	递归版本 需要将最后的返回值向上传递
*/
func solution3(head *ListNode) *ListNode {
	// base
	if head == nil || head.Next == nil {
		return head
	} else {
		// 这里把最后的返回值向上传递
		cur := solution3(head.Next)
		head.Next.Next = head
		// 这里是第一个元素的 next 设置为 nil
		head.Next = nil
		return cur
	}
}
