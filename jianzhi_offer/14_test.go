package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("14 链表中倒数第k个结点", func(t *testing.T) {
		n := 3
		want := 4
		got := solution(n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	输入一个链表，输出该链表中倒数第k个结点
	将链表入栈，然后 pop k 次元素后得到的就是所求
	直接遍历，然后倒数第 k 个节点就是正数的第 n - k 个节点
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	输入一个链表，输出该链表中倒数第k个结点
	将链表入栈，然后 pop k 次元素后得到的就是所求
	直接遍历，然后倒数第 k 个节点就是正数的第 n - k 个节点
	使用双指针就不需要先进行遍历来找长度
		 初始化 former latter 指针，开始都指向 head
		 former 先向前遍历 k 个节点，此时两个指针相距 k 个节点
		 两个指针一起遍历直到 former 指向 null 即尾部的下一个
		 latter 与尾指针距离为 k - 1，此时 latter 就是所求
*/

func solution(head *ListNode, k int) *ListNode {
	former, latter := head, head
	for i := 0; i < k; i++ {
		if former.Next != nil {
			former = former.Next
		} else {
			return nil
		}
	}
	for former != nil {
		former = former.Next
		latter = latter.Next
	}
	return latter
}
