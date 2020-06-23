package jianzhi_offer

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("20. 包含min函数的栈 ", func(t *testing.T) {
		nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		want := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
		get := solution(nums)
		if !reflect.DeepEqual(get, want) {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	如果只使用一个元素来保存当前的最小元素，则每次 pop 都需要更新该元素
	如果每次插入都要求元素有序，那就需要维护一个单调递减栈
	上面的都太麻烦了

	直接使用一个辅助的栈来保存当前的最小元素值
	还有一种方法是只使用一个栈， 然后栈中的每个元素都是一个节点，节点里面保存当前的元素以及当前元素为止的最小值
*/
type MinStack struct {
	stack1 []int // 保存实际的元素
	stack2 []int // 保存当前的最小值
}

/** initialize your data structure here. */
func Constructor() MinStack {
	stack1 := make([]int, 0)
	stack2 := make([]int, 0)
	return MinStack{stack1: stack1, stack2: stack2}
}

// 加入新的元素，如果 stack2 为空或者当前插入的元素小于等于 stack2 栈顶元素，则需要将当前元素也保存到 stack2
func (this *MinStack) Push(x int) {
	stack2Len := len(this.stack2)
	if stack2Len == 0 || x <= this.stack2[stack2Len-1] {
		this.stack2 = append(this.stack2, x)
	}
	this.stack1 = append(this.stack1, x)
}

// 如果当前出栈的元素等于最小的元素，则两个栈都需要将元素出栈
func (this *MinStack) Pop() {
	stack1Len := len(this.stack1)
	stack1Top := this.stack1[stack1Len-1]
	stack2Len := len(this.stack2)
	stack2Top := this.stack2[stack2Len-1]
	if stack1Top == stack2Top {
		this.stack2 = this.stack2[0 : stack2Len-1]
	}
	this.stack1 = this.stack1[0 : stack1Len-1]
}

func (this *MinStack) Top() int {
	return this.stack1[len(this.stack1)-1]
}

func (this *MinStack) Min() int {
	return this.stack2[len(this.stack2)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

/*

	还有一种方法是只使用一个栈， 然后栈中的每个元素都是一个节点，节点里面保存当前的元素以及当前元素为止的最小值
*/
type Node struct {
	v    int
	minV int
}

type MinStack2 struct {
	stack []Node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	stack := make([]Node, 0)
	return MinStack2{stack: stack}
}

// 如果当前的元素小于等于上一个元素的最小值，则需要将当前元素作为当前节点的最小值，否则使用上一个元素的最小值
func (this *MinStack2) Push(x int) {
	stackLen := len(this.stack)
	lastMinV := this.stack[stackLen-1].minV
	if x <= lastMinV {
		this.stack = append(this.stack, Node{x, x})
	} else {
		this.stack = append(this.stack, Node{x, lastMinV})
	}
}

func (this *MinStack2) Pop() {
	stackLen := len(this.stack)
	this.stack = this.stack[0 : stackLen-1]
}

func (this *MinStack2) Top() int {
	return this.stack[len(this.stack)-1].v
}

func (this *MinStack2) Min() int {
	return this.stack[len(this.stack)-1].minV
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
