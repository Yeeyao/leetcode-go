package jianzhi_offer

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("25  复杂链表的复制", func(t *testing.T) {
		A := &Node{3, &Node{9, nil, nil}, &Node{20, &Node{15, nil, nil}, &Node{7, nil, nil}}}
		get := solution(A)
		want := []int{3, 9, 20, 15, 7}
		if !reflect.DeepEqual(get, want) {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
	直接遍历 BFS 或者 DFS，用额外 map 记录每个 Node 是否已经出现
	如果 head 是 nil，直接返回
	如果当前节点已经创建过，直接返回
	复制当前节点，然后递归调用 next 以及 random
*/
func solution(head *Node) *Node {
	// 已经创建了的节点信息
	var visited map[*Node]bool
	// nil 或者已经创建了
	if head == nil {
		return nil
	}
	if _, ok := visited[head]; ok {
		return nil
	} else {
		visited[head] = true
	}
	copyNode := Node{Val: head.Val, Next: head.Next, Random: head.Random}
	copyNode.Next = solution(copyNode.Next)
	copyNode.Random = solution(copyNode.Random)
	return &copyNode
}

/*
	迭代 ABC 变成 AA'BB'CC' 然后将它们分开
*/
func solution2(head *Node) *Node {
	if head == nil {
		return head
	}
	cur := head
	for cur != nil {
		newNode := new(Node)
		newNode.Val = cur.Val
		newNode.Next = cur.Next
		cur.Next = newNode
		cur = newNode.Next
	}
	// 更新第一步复制节点的 Random
	// 由于复制的节点紧跟在原始节点的后面，所以复制节点的 Random 就相当于
	// 原始节点 Random 的后方（就是复制出的 Random）
	cur = head
	for cur != nil {
		// 当前节点如果是原来的节点
		if cur.Random != nil {
			// 复制的节点的 Random 是当前节点的 Random.Next
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// AA' BB' CC' 重新连接
	curOldList := head
	curNewList := head.Next
	newHead := head.Next
	for curOldList != nil {
		// 旧的重新连接
		curOldList.Next = curOldList.Next.Next
		// 新的重新连接
		if curNewList.Next != nil {
			curNewList.Next = curNewList.Next.Next
		}
		curOldList = curOldList.Next
		curNewList = curNewList.Next
	}
	return newHead
}

// 同上
func solution3(head *Node) *Node {
	// 先将原链表的节点都拷贝并在原节点后面连接
	cloneNode(head)
	connectSiblingNode(head)
	return reconnectNode(head)
}

// 原节点复制并连接到节点的后面
// AA' BB' CC'
func cloneNode(head *Node) {
	pNode := head
	for pNode != nil {
		newNode := new(Node)
		newNode.Val = pNode.Val
		newNode.Next = pNode.Next
		pNode.Next = newNode
		pNode = newNode.Next
	}
}

// 这里更新第一步中复制节点的 Random
func connectSiblingNode(head *Node) {
	pNode := head
	for pNode != nil {
		// 获取复制的节点
		pCloned := pNode.Next
		// 当前节点是复制的节点，
		if pNode.Random != nil {
			pCloned.Random = pNode.Random.Next
		}
		pNode = pCloned.Next
	}
}

// 返回复制链表的头节点
func reconnectNode(head *Node) *Node {
	pNode := head
	var pCloneHead, pCloned *Node = nil, nil

	if pNode != nil {
		pCloneHead = pNode.Next
	}
	for pNode != nil {
		pCloned = pNode.Next
		pNode.Next = pCloned.Next
		if pNode.Next != nil {
			pCloned.Next = pNode.Next.Next
			pCloned = pCloned.Next
		}
		pNode = pNode.Next
	}
	return pCloneHead
}
