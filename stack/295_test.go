package stack

/*
	从数据流中找到中位数。中位数是有序整型列表的中间数值，如果列表元素个数是偶数，则没有中间数值，此时中位数是两个中间数值的平均值
	需要实现类的方法
	假设列表长度是 n，下标从 1 开始，则如果 n 是奇数，中位数是第 (n + 1) / 2 个，如果 n 是偶数，中位数是 (n + 1) / 2 以及 (n + 1) / 2 + 1 两个元素平均值

	使用一个大顶堆保存前 k 个最小的元素
	使用一个小顶堆保存前 k 个最大的元素
	最终的结果就在这两个堆的堆顶中计算。同时，因为都是找的中间的数值，因此这两个堆的元素数量相差值不能大于 1，当这个差值过大，需要将元素在堆之间移动
*/

type MedianFinder struct {
}

func Constructor() MedianFinder {

}

func (this *MedianFinder) AddNum(num int) {

}

func (this *MedianFinder) FindMedian() float64 {

}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
