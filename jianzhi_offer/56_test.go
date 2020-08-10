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
	添加一个头节点处理第一个和第二个节点相同的情况
	使用 pre 指向当前确定不重复的节点，last 指向当前处理的节点

	1 2 2 2 3 3 3 4 4 5
	只有一个节点就直接返回
	开头节点前添加一个节点
	初始化 pre, last 两个节点分别指向添加的头以及原来的头，这里逐个判断的是 last
	循环终止条件是 last == nil
		判断 last.Next != nil 以及 last.Val == last.Next.Val 同时，一直过滤
			pre.Next = last.Next last = last.Next
		不满足上述判断
			pre = pre.Next last = last.Next
*/
func solution(phead *Node) *Node {
	// 只有一个节点，直接返回
	if phead == nil || phead.Next == nil {
		return phead
	}
	// 在开头节点添加一个节点
	head := &Node{0, nil}
	head.Next = phead
	pre, last := head, head.Next
	for last != nil {
		// 非最后节点然后和后面的值相同
		if last.Next != nil && last.Val == last.Next.Val {
			// 继续向后找到最后一个相同的节点
			for last.Next != nil && last.Val == last.Next.Val {
				last = last.Next
			}
			// 注意这里是 pre.Next，所以 pre 还是在 last 前面
			pre.Next = last.Next
			last = last.Next
		} else {
			// 没有相等的，两个指针正常向后移动
			pre = pre.Next
			last = last.Next
		}
		// 这里起那面加了一个
		return head.Next
	}
}
