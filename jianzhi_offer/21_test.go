package jianzhi_offer

import (
	"fmt"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("21. 栈的压入、弹出序列 ", func(t *testing.T) {
		pushed := []int{1, 2, 3, 4, 5}
		poped := []int{4, 5, 3, 2, 1}
		want := true
		get := solution(pushed, poped)
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	给定一个栈以及弹出元素的栈，判断第二个能否作为第一个的弹出的顺序 元素不会相等
	使用一个额外数组保存每个元素值的相对位置，遍历出栈元素的栈元素。
	然后保存一个当前出栈中的最小元素值索引，每次遍历到元素都需要比较是否小于最小的元素索引，
	更新最小的索引并继续遍历
	上面的思路只适用于入栈的元素数组已经固定的情况

	如果前面的元素已经入栈了，想要 pop 后面的元素出来
	这里题目的意思只是第一个数组是入栈的，第二个是出栈的顺序而已，具体出栈入栈操作未知
	类似拓扑排序？

	使用一个辅助栈来模拟入栈，出栈操作
		入栈操作按照压栈序列顺序执行压入辅助栈
		出栈操作 如果辅助栈栈顶的元素等于弹出栈序列的栈顶，则不断执行出栈
		当出栈的元素栈为空表示全部都可以出栈了
*/
func solution(pushed []int, popped []int) bool {
	stackLen := len(pushed)
	tempStack := make([]int, stackLen)
	stIndex, popCount := 0, 0
	// 遍历的是 pushed
	for i := 0; i < stackLen; i++ {
		// 先从 pushed 中将元素放入辅助栈， 然后比较
		tempStack[stIndex] = pushed[i]
		stIndex++
		for stIndex > 0 && tempStack[stIndex-1] == popped[popCount] {
			popCount++
			stIndex--
		}
	}
	return stIndex == 0
}

/*
	stack implementation
*/
type Stack []int

func (s *Stack) Push(val int) { *s = append(*s, val) }
func (s Stack) Peek() int     { return s[len(s)-1] }
func (s Stack) Len() int      { return len(s) }
func (s *Stack) Pop() int {
	temp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return temp
}
func (s Stack) Empty() bool {
	return len(s) == 0
}
