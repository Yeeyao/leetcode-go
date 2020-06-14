package jianzhi_offer

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("3 从尾到头打印链表", func(t *testing.T) {
		s := Node{1, nil}
		 solution(s)
	})
}

/*
	可以直接用递归的方法处理
	也可以用栈的方法处理
 */

type Node struct  {
	val int
	next *Node
}

func solution(node Node) {
	if node.next != nil {
		solution(*node.next)
	}
	println(node.val)
}

//func solution(node Node) {
//	var st []int
//	for node != nil {
//		st.Push(node.val)
//		node = *node.next
//	}
//	for !st.IsEmpty {
//		n := st.Peek()
//		fmt.Println(n)
//		st.Pop()
//	}
//}
