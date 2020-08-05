package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("55 链表中环的入口结点 ", func(t *testing.T) {
		root := Node{1, nil}
		get := solution(&root)
		want := nil
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

type Node struct {
	Val  int
	Next *Node
}

/*
	类似龟兔赛跑，一个一次一步，一个一次两步相遇就有环
	也可以用 hash map 来存储每个遍历过的节点并每次都判断是否遍历过
*/
func solution(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head
	// 未相遇
	for fast != slow {
		if fast == nil || slow == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 相遇之后找入口
	first, second := fast, head
	for first != second {
		first = first.Next
		second = second.Next
	}
	return first
}
