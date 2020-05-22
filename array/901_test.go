package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("901. Online Stock Span", func(t *testing.T) {
		nums := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
		W := 3
		want := true
		got := solution(nums, W)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	the span of stock's prices 定义为从开始天数开始价格小于或者等于
	当天的连续天数 这里用单调递减栈 这里的栈需要保存 {value, num}
	数值即数值,num 是遍历到该位置小于等于当前元素的元素数量
	遍历数组，如果当前的元素大于栈元素，就一直将栈元素出栈，同时当前元素的 span
	需要加上栈元素的 span。最后将当前元素入栈。
	这里用两个数组来处理模拟
*/
// 这里两个数组需要对应
type StockSpanner struct {
	price []int
	span  []int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	priceLen := len(this.price)
	if priceLen == 0 {
		this.price = append(this.price, price)
		this.span = append(this.span, 1)
		return 1
	}
	// 默认 span
	span := 1
	// 栈顶
	stTop := priceLen - 1
	// 这里直接将栈顶向前移动
	for stTop >= 0 && price >= this.price[stTop] {
		// 当前元素的 span 加上栈顶元素的 span
		span += this.span[stTop]
		stTop -= this.span[stTop]
	}
	// 当前元素入栈
	this.price = append(this.price, price)
	this.span = append(this.span, span)
	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

// 另一种模拟方式 [price, span]
type StockSpanner struct {
	stack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{make([][2]int, 0)}
}

func (this *StockSpanner) Next(price int) int {
	span := 1
	for len(this.stack) > 0 && this.stack[len(this.stack)-1][0] <= price {
		// 加上 span
		span += this.stack[len(this.stack)-1][1]
		// 模拟出栈
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, [2]int{price, span})
	return span
}
