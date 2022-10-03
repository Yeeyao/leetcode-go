package link_list

/*
138. Copy List with Random Pointer
A linked list is given such that each node contains an additional random pointer
which could point to any node in the list or null.
Return a deep copy of the list.

复制含有随机节点的链表 同 剑指 offer 25

如果想要一开始先忽略随机的节点的复制，先按照顺序复制，然后考虑处理随机节点。如果这样，需要记录旧节点到复制节点的映射关系。第二次遍历的时候。新节点
查找映射关系找到新的随机节点

这里将全部节点复制一次，然后原来的节点的下一个节点指向复制的节点。通过指向记录两者的映射关系，方便处理随机节点
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
先将 ABC 构造成 AA'BB'CC' 然后分离开
*/
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	// 先复制节点，先复制 Next 以及 Val
	cur := head
	for cur != nil {
		copyCur := Node{Val: cur.Val, Next: cur.Next}
		cur.Next = &copyCur
		cur = copyCur.Next
	}
	// 再复制 random 节点
	cur = head
	for cur != nil {
		// 需要判断 nil
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 再分离
	oldCur := head
	newCur := head.Next
	retHead := head.Next
	for oldCur != nil {
		// 旧链表最后位置的 Next 需要指向 nil
		oldCur.Next = oldCur.Next.Next
		// 新链表最后位置需要判断
		if newCur.Next != nil {
			newCur.Next = newCur.Next.Next
		}
		oldCur = oldCur.Next
		newCur = newCur.Next
	}
	return retHead
}
