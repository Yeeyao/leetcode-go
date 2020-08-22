package jianzhi_offer

import "fmt"

/*
实现 FreqStack，模拟类似栈的数据结构的操作的一个类。FreqStack 有两个函数：
push(int x)，将整数 x 推入栈中。pop()，它移除并返回栈中出现最频繁的元素。如果最频繁的元素不只一个，则移除并返回最接近栈顶的元素。
◼ 示例： push [5,7,5,7,4,5] pop() -> 返回 5，因为 5 是出现频率最高的。栈变成 [5,7,5,7,4]。
pop() -> 返回 7，因为 5 和 7 都是频率最高的，但 7 最接近栈 顶。栈变成 [5,7,5,4]。
pop() -> 返回 5 。栈变成 [5,7,4]。 pop() -> 返回 4 。栈变成 [5,7]。
*/

/*
	[leetcode](https://leetcode-cn.com/problems/maximum-frequency-stack/)
实现 FreqStack，模拟类似栈的数据结构的操作的一个类。
FreqStack 有两个函数：
    push(int x)，将整数 x 推入栈中。
    pop()，它移除并返回栈中出现最频繁的元素。
        如果最频繁的元素不只一个，则移除并返回最接近栈顶的元素。
*/

/*
	精妙的是使用 grup map 相同的频率的元素按照入栈的顺序保存

	使用 map 保存每个元素的出现频率，maxFeq 保存最高频率
	freq 是 map[int]int key 是 x val 是频率
	group 是 map[int][]int key 是频率，val 是 stack 保存的是对应的当前频率的元素
		group 将相同频率的元素保存到 stack 中，然后最近的元素保存在栈顶
	初始化则初始化两个 map 以及 maxFreq 是 0

	push
		从 freq 中获取 x 的频率并 + 1 和保存
		如果它的频率大于 maxFreq 则更新 maxFreq
		判断 group 中 key 为 f 的 pair 是否存在，是则将 x 添加到 stack 中
		不存在则直接新建 stack 并保存
	pop
		获取最大频率 f 的 stack 并将栈顶元素弹出 note: 这里的元素按照入栈的顺序保存的，所以这里一定是保存最近的元素
		更新其最大频率 freq 为 f - 1
		如果对应的 group 的 maxFreq 的值为 0 则将 maxFreq--
*/
type freqStack struct {
	// key x val frequency
	freq map[int]int
	// key frequency val value stack
	group map[int][]int
	// max frequency
	maxFrequency int
}

func Constructor() freqStack {
	return freqStack{maxFrequency: 0}
}

func (this *freqStack) Push(x int) {
	// 处理 freq 以及 maxFreq
	var newFreq int
	if oldFreq, ok := this.freq[x]; ok {
		newFreq = oldFreq + 1
	} else {
		newFreq = 1
	}
	this.freq[x] = newFreq
	if newFreq > this.maxFrequency {
		this.maxFrequency = newFreq
	}
	// 处理 group
	if stack, ok := this.group[newFreq]; ok {
		stack = append(stack, x)
		this.group[newFreq] = stack
	} else {
		stack = []int{x}
		this.group[newFreq] = stack
	}
}

func (this *freqStack) Pop() int {
	// group 更新
	if stack, ok := this.group[this.maxFrequency]; ok {
		// 将栈顶元素出栈
		stackTopVal := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 栈顶元素的频率减小
		this.freq[stackTopVal]--
		// 栈空，表示当前没有了最大出现频率的数字
		if len(stack) == 0 {
			this.maxFrequency = this.maxFrequency - 1
		}
		return stackTopVal
	} else {
		fmt.Errorf("fatal error!\n")
		return -1
	}
}
