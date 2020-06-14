package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("5 用两个栈实现队列", func(t *testing.T) {
		nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		value := 4
		want := true
		got := solution(nums, value)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	两个栈实现队列 队列的特点是先进先出 FIFO 栈的特点是先进后出 FILO
	两个栈轮流使用，一开始用一个栈 a 来保存 push 的元素，然后当需要 pop 的时候
	直接将 a 的元素全部 push 到 b 中，留下最后一个元素直接删除，下次再 push 的时候，
	将 b 的元素全部 push 到 a 中，然后再将新加入的元素 push 到 a

	比如一开始 push 1， 2， 3， 4 a: 1 2 3 4 b :()
	pop a:() b: 4 3 2 1(被 pop) 再 push a:2 3 4 b:() push 5 a: 2 3 4 5 b:()
*/
type Queue []int

func (q *Queue) Push(a int, s1, s2 Stack) {
	// 之前元素都在 s2 中
	if len(s1) == 0 {
		// 之前已经保存到 s2
		for len(s2) > 0 {
			s1.Push(s2.Pop())
		}
		s1.Push(a)
	}
}

func (q *Queue) Pop(s1, s2 Stack) int {
	// 之前已经从 s1 pop 到了 s2 直接从 s2 pop
	if len(s1) == 0{
		if len(s2) > 0 {
			return s2.Pop()
		} else {
			return 0
		}
		// 还在 s1 中
	} else {
		for len(s1) > 1 {
			s1.Push(s2.Pop())
		}
		return s1.Pop()
	}
}

type Stack []int

func (s *Stack) Push(a int) {
	*s = append(*s, a)
}
func (s *Stack) Pop() int {
	sLen := len(*s)
	res := (*s)[sLen-1]
	*s = (*s)[0 : sLen-2]
	return res
}
func (s Stack) Peek() int {
	return s[len(s)-1]
}
func (s Stack) Len() int {
	return len(s)
}
