package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("84. Largest Rectangle in Histogram", func(t *testing.T) {
		nums := []int{2, 1, 5, 6, 2, 3}
		want := 10
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("84. Largest Rectangle in Histogram2", func(t *testing.T) {
		nums := []int{2, 1, 5, 6, 2, 3}
		want := 10
		got := solution2(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("84. Largest Rectangle in Histogram3", func(t *testing.T) {
		nums := []int{1}
		want := 1
		got := solution2(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("84. Largest Rectangle in Histogram4", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		want := 30
		got := solution2(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("84. Largest Rectangle in Histogram2", func(t *testing.T) {
		nums := []int{2, 1, 5, 6, 2, 3}
		want := 10
		got := solution3(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("84. Largest Rectangle in Histogram3", func(t *testing.T) {
		nums := []int{1}
		want := 1
		got := solution3(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("84. Largest Rectangle in Histogram4", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		want := 30
		got := solution3(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	找出给定的高度 类似 11 42 但是这里没有说高度不能是 0 的
	需要当前元素大于下一个元素，以当前元素为局部的最高值，然后向前遍历到开头计算每个的面积
	因为这里选取局部最高值，非最高值的面积也将会被计算进去的
	从左到右遍历，非最后一个元素且下一个元素大于等于当前元素，下一次循环
	相当于在 brute forces 基础上，过滤掉一些元素节省重复判断，剪枝
*/
func solution(heights []int) int {
	res := 0
	hLen := len(heights)
	for i := 0; i < hLen; i++ {
		// 过滤掉非局部最高值
		if i+1 < hLen && heights[i] <= heights[i+1] {
			continue
		}
		// 当前的最低高度
		minH := heights[i]
		// 向前计算每个面积
		for j := i; j >= 0; j-- {
			if heights[j] < minH {
				minH = heights[j]
			}
			area := minH * (i - j + 1)
			if area > res {
				res = area
			}
		}
	}
	return res
}

/*
	stack 做法 一样当前元素小于下一个元素就跳过当前元素
	使用单调递增栈
*/
func solution2(heights []int) int {
	res := 0
	// 先存放一个 0 元素方便处理最后一个元素
	stTop := 0
	heights = append(heights, 0)
	hLen := len(heights)
	st := make([]int, hLen)
	for i := 0; i < hLen; i++ {
		// 空或者是当前元素大于栈顶就入栈
		if stTop == 0 || heights[i] > heights[st[stTop-1]] {
			st[stTop] = i
			stTop++
		} else {
			// 这里计算，高度是栈顶，长度空的话是当前元素到栈顶的索引差
			// 因为这里是当前元素小于栈顶了，所以使用当前元素作为最低高度
			stTop--
			width := i
			if stTop > 0 {
				width = i - st[stTop-1] - 1
			}
			temp := heights[st[stTop]] * width
			if temp > res {
				res = temp
			}
			// 注意这里的递减，只要当前元素还不能入栈就一直判断
			// 跟上面的思路是一样的，只不过用单调递增栈保存，然后反过来遍历
			// 这里递减也实现维护单调递增栈的功能
			i--
		}
	}
	return res
}

/*
	上面的优化 这个更加常规一点
*/
func solution3(heights []int) int {
	res := 0
	// 先存放一个 0 元素方便处理最后一个元素
	stTop := 0
	heights = append(heights, 0)
	hLen := len(heights)
	st := make([]int, hLen)
	for i := 0; i < hLen; i++ {
		// 这里只要当前元素小于栈顶就不断出栈
		for stTop != 0 && heights[st[stTop-1]] >= heights[i] {
			stTop--
			width := i
			if stTop > 0 {
				width = i - st[stTop-1] - 1
			}
			temp := heights[st[stTop]] * width
			if temp > res {
				res = temp
			}
		}
		// 当前元素入栈
		st[stTop] = i
		stTop++
	}
	return res
}

/*
	自定义 stack
*/
type stack []int

func (s *stack) Push(val int) { *s = append(*s, val) }
func (s stack) Peek() int     { return s[len(s)-1] }
func (s stack) Len() int {
	return len(s)
}
func (s *stack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return x
}
