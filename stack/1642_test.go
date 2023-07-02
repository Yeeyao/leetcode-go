package stack

import (
	"container/heap"
	"sort"
)

/*
	给定建筑物高度数组以及 n 个 bricks 以及 n 个 ladders。需要从 0 号建筑物开始不断向后移动，从建筑物 i 到 i + 1 移动时：
	- 如果建筑物 i 高度大于等于 i + 1 则不需要梯子和砖块
	- 如果建筑物 i 高度小于 i + 1，可以使用一架梯子和 h[i+1]-h[i] 个砖块
	- 使用最佳的方式使用梯子和砖块，返回可以到达的最远的建筑物的下标

	这里的关键是怎么使用砖块，怎么感觉有点动态规划的味道

	一开始的条件是 0 号建筑，bricks 个砖头和 ladders 个梯子
	终止条件是 bricks 不够 hdiff=h(x+1)-h(x) 或者 ladders == 0，此时在 x 号建筑
	转换方程：
		如果 h(x) > h(x+1) x = x + 1
		否则 hdiff=h(x+1)-h(x)
			如果 ladders > 0 以及 bricks >= hdiff 则两个都可以选择
				选择其中一个并继续下去，多种情况下更新最大的建筑物序号
			否则到达终止条件

	题解更加简单。
	第一次遍历的时候都使用梯子，然后等到梯子不够用的时候，就使用砖头。
	用砖头的时机是当初用梯子的高度差比现在的高度差小的时候（将梯子留给高度差更大的？）。
	这里需要将使用梯子的建筑物高度保存起来，梯子用完了就使用砖头来兑换
	优先兑换的是高度差大的，每次从存储的高度差中选最大的。动态求极值就使用堆

	题解里面给的例子是一开始无脑使用砖头，然后选择高度差最大的来使用梯子进行替换以获得最多的砖头
	这里什么情况下梯子替换的时候会加砖头？就是选择替换高度差最大的时候得到的砖头-当前的高度差值大于 0

*/

func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := hp{IntSlice: []int{}}
	// 从第二个建筑物开始看
	for i := 1; i < len(heights); i++ {
		// 计算高度差
		diff := heights[i] - heights[i-1]
		// 不需要借助工具，就到下一个建筑
		if diff <= 0 {
			continue
		}
		// 砖块不够同时还有剩余的梯子
		if bricks < diff && ladders > 0 {
			// 使用梯子
			ladders -= 1
			// 替换堆最大的砖块，如果有多出来的砖块则需要添加到总砖块上
			// 因为这里只需要 diff 个砖块，如果替换之前的高度比 diff 大就有剩余的砖块可以使用
			if h.Len() > 0 && -h.IntSlice[0] > diff {
				bricks -= heap.Pop(&h).(int)
			} else {
				continue
			}
		}
		// 使用砖块
		bricks -= diff
		// 不够了，直接返回
		if bricks < 0 {
			return i - 1
		}
		// 每次都需要将高度差加入到堆
		heap.Push(&h, -diff)
	}
	// 全部够了，返回
	return len(heights) - 1
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	top := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return top
}
