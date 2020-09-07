package link_list

/*
138. Copy List with Random Pointer
A linked list is given such that each node contains an additional random pointer
which could point to any node in the list or null.
Return a deep copy of the list.

复制含有随机节点的链表 同 剑指 offer 25
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
	先将 ABC 构造成 AA'BB'CC'
	然后分离开 这里使用新的变脸保存 head
*/

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 先处理复制和 next
	cur := head
	for cur != nil {
		copyNode := Node{Val: cur.Val, Next: cur.Next}
		cur.Next = &copyNode
		cur = copyNode.Next
	}
	// 处理 random
	cur = head
	for cur != nil {
		// 这里也需要判断，因为可能原来的 random 是 nil
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 分离 这里需要保存新的头部先
	newHead := head.Next
	oldList := head
	newList := head.Next
	for oldList != nil {
		oldList.Next = oldList.Next.Next
		// 需要处理只有一个元素的情况
		if newList.Next != nil {
			newList.Next = newList.Next.Next
		}
		oldList = oldList.Next
		newList = newList.Next
	}
	return newHead
}
