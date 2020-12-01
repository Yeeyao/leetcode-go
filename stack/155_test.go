package stack

/*
	两种思路
		1.栈保存{ele, min}
		2.使用额外的栈保存 min，入栈时更新 min 栈（如果当前元素小于或者等于 min 都需要入栈）
		出栈时判断是否 min 并对应 min 处理
*/
type MinStack struct {
	eleStack []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	eleStack := make([]int, 0)
	minStack := make([]int, 0)
	return MinStack{eleStack: eleStack, minStack: minStack}
}

func (this *MinStack) Push(x int) {
	lenMinS := len(this.minStack)
	if lenMinS == 0 {
		this.minStack = append(this.minStack, x)
	} else {
		minTop := this.minStack[lenMinS-1]
		if x <= minTop {
			this.minStack = append(this.minStack, x)
		}
	}
	this.eleStack = append(this.eleStack, x)
}

func (this *MinStack) Pop() {
	minTop := this.minStack[len(this.minStack)-1]
	eleTop := this.eleStack[len(this.eleStack)-1]
	if minTop == eleTop {
		this.minStack = this.minStack[0 : len(this.minStack)-1]
	}
	this.eleStack = this.eleStack[0 : len(this.eleStack)-1]
}

func (this *MinStack) Top() int {
	return this.eleStack[len(this.eleStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
