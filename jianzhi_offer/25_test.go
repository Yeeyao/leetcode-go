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
	var visited map[*Node]bool
	// nil 或者已经创建了
	if head == nil {
		return nil
	}
	if _, ok := visited[head]; ok {
		return nil
	}
	copyNode := Node{Val: head.Val, Next: head.Next, Random: head.Random}
	copyNode.Next = solution(copyNode.Next)
	copyNode.Random = solution(copyNode.Random)
	return &copyNode
}
