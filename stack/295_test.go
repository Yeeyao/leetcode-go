package stack

import (
	"container/heap"
	"sort"
)

/*
	从数据流中找到中位数。中位数是有序整型列表的中间数值，如果列表元素个数是偶数，则没有中间数值，此时中位数是两个中间数值的平均值
	需要实现类的方法
	假设列表长度是 n，下标从 1 开始，则如果 n 是奇数，中位数是第 (n + 1) / 2 个，如果 n 是偶数，中位数是 (n + 1) / 2 以及 (n + 1) / 2 + 1 两个元素平均值

	使用一个大顶堆保存前 k 个最小的元素
	使用一个小顶堆保存前 k 个最大的元素
	最终的结果就在这两个堆的堆顶中计算。同时，因为都是找的中间的数值，因此这两个堆的元素数量相差值不能大于 1，当这个差值过大，需要将元素在堆之间移动，将较多的移动到较少的

	新增的数字需要怎么判断应该添加到哪个堆里面呢，两个都添加？以及堆之间的元素移动是在哪个时机进行？

	几个问题，首先是元素需要先放到哪个堆？这里哪个先放都可以
	如果放不下第一个堆，是否需要放到第二个堆？需要。两个都放不下，元素应该是需要被丢弃的（表示元素数值是在中间的），是被丢弃
	什么时候判断堆的大小，放入元素之后吧，这里可以是放入之后立即判断，也可以两个都判断放入之后才判断
	最终的结果需要根据先放入的堆来确定，因为两个堆都存放前 K 个大（小）元素，因此只需要判断两个堆的大小就可以知道返回哪个是结果

	这里给的例子是先放到大顶堆（如果可以放入的话），然后才放到小顶堆。
	之后比较两个堆的大小，先比较大顶堆元素是否比小顶堆元素数量超过 1，是就将大顶堆堆顶放到小顶堆
	如果小顶堆元素数量比大顶堆多，则将小顶堆堆顶元素放到大顶堆
	这里的选择关系到最后返回的结果，如果元素是奇数，则只需要返回大顶堆堆顶，如果偶数，则两个堆数量相等？返回两堆顶的均值

	如果是先放到小顶堆，然后才放到大顶堆，则类似。先放，然后比较堆大小，
	小顶堆元素是否比大顶堆元素数量超过 1，是就将小顶堆堆顶放到大顶堆
	这里的选择关系到最后返回的结果，如果元素是奇数，则只需要返回小顶堆堆顶，如果偶数，则两个堆数量相等？返回两堆顶的均值
	[ref](https://leetcode.cn/problems/find-median-from-data-stream/solution/shu-ju-liu-de-zhong-wei-shu-by-leetcode-ktkst/)

	times: 1

*/

// 因为两个都是使用 int slice(元素按照大小升序排列) 因此小顶堆需要将元素取相反数再入堆
type MedianFinder struct {
	MinHeap hp
	MaxHeap hp
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	minH, maxH := &this.MinHeap, &this.MaxHeap
	// 等于也要放入
	if minH.Len() == 0 || num <= -minH.IntSlice[0] {
		heap.Push(minH, -num)
	} else {
		heap.Push(maxH, num)
	}
	// 需要注意符号变化
	if minH.Len()-maxH.Len() > 1 {
		heap.Push(maxH, -heap.Pop(minH).(int))
	} else if maxH.Len() > minH.Len() {
		heap.Push(minH, -heap.Pop(maxH).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	minH, maxH := this.MinHeap, this.MaxHeap
	if minH.Len() > maxH.Len() {
		return float64(-minH.IntSlice[0])
	}
	return float64(-minH.IntSlice[0]+maxH.IntSlice[0]) / 2
}

// 利用 heap.Push heap.Pop 所需要实现的 interface
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

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
