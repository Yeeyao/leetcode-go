package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("56 删除链表中重复的结点 ", func(t *testing.T) {
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
	删除链表中的重复节点
	应该需要记录当前节点和下一个节点
	当前节点，判断下一个节点是否为空，如果是就直接返回根节点了
	当前节点和下一个节点非空，判断下一个节点和当前节点是否数值相同
		是，需要将节点继续移动
		否，继续遍历

	1 2 2 2 3 3 3 4 4 5
*/
func solution(head *Node) *Node {
	newHead := head
	for head != nil && head.Next != nil {
		temp := head
		// 过滤相同的节点
		for temp.Val == temp.Next.Val {
			temp = temp.Next
		}
		head.Next = temp.Next
		head = head.Next
	}
	return newHead
}
