package stack

import (
	"reflect"
	"testing"
)

/*
Given a list of daily temperatures T, return a list such that, for each day in the input,
tells you how many days you would have to wait until a warmer temperature.
If there is no future day for which this is possible, put 0 instead.
For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].
Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100].

给定一个每日温度列表，需要找到每个天数后面第几天比今天的温度高，如果没有就是 0
即找到每个温度后面第一个比它高的温度并计算天数差
从最后的元素向前遍历，最后一个结果是 0，然后倒数第二个判断最后一个是否大于它，
右边构建单调递增栈，元素是数值加上位置

[ref](https://leetcode-cn.com/problems/daily-temperatures/solution/mei-ri-wen-du-by-leetcode-solution/)

初始化栈，对每个元素，找到栈中第一个大于它的元素并将距离保存
	如果找不到就保存 0
之后元素需要判断入栈，栈顶元素如果大于当前元素就直接入栈，否则，将栈顶元素出栈最后入栈
*/

func TestPro(t *testing.T) {
	t.Run("739. Daily Temperatures", func(t *testing.T) {
		input := []int{73, 74, 75, 71, 69, 72, 76, 73}
		want := []int{1, 1, 4, 2, 1, 1, 0, 0}
		got := dailyTemperatures(input)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("739. Daily Temperatures2", func(t *testing.T) {
		input := []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}
		want := []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0}
		got := dailyTemperatures(input)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

type valPos struct {
	Val int
	Pos int
}

func dailyTemperatures(T []int) []int {
	stack := make([]valPos, 0)
	res := make([]int, 0)
	tLen := len(T)
	// 这里是先找然后才入栈
	for i := tLen - 1; i >= 0; i-- {
		stackLen := len(stack)
		// 栈为空
		if stackLen == 0 {
			res = append(res, 0)
			stack = []valPos{{T[i], i}}
		} else {
			stackHead := stackLen - 1
			// 需要注意等号处理 因为这里是从后面向前面遍历，所以相等需要将栈顶元素出栈，用较前面的更新后面的
			for stackHead >= 0 && stack[stackHead].Val <= T[i] {
				stackHead--
			}
			// 栈空了
			if stackHead < 0 {
				res = append(res, 0)
				stack = []valPos{{T[i], i}}
			} else {
				res = append(res, stack[stackHead].Pos-i)
				stack = stack[:stackHead+1]
				stack = append(stack, valPos{T[i], i})
			}
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

/*
从头开始遍历 stack 保存的是 T 元素索引
内部循环条件，栈非空同时栈顶索引对应的元素小于当前元素就保存结果并将栈顶元素出栈
最后将元素入栈

维护一个存储下标的单调栈，栈底到栈顶的元素温度依次递减
如果一个下标在单调栈里，表示还未找到下一个温度更高的下标
正向遍历温度列表，对每个元素 T[i]，如果栈空则直接将 i 入栈，非空，则比较栈顶元素
T[st] 与当前元素的温度 T[i]，如果 T[i] > T[st] 则知道 res[st] = i - st
重复直到栈顶元素大于 T[i] 或者栈空
这里的思路是将栈顶和当前元素的身份反过来比上面的解法

感觉我想复杂了
*/
func dailyTemperatures(T []int) []int {
	ret := make([]int, len(T))
	stack := []int{}
	for i := 0; i < len(T); i++ {
		// 注意这里是循环，同样是构建一个单调递增栈
		// 栈非空同时当前元素大于栈顶元素，栈顶元素出栈同时返回值需要计算
		// 更新栈顶元素的数值
		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			ret[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ret
}

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	stack := make([]int, 0)
	for i := 0; i < len(T); i++ {
		// 这里的意思是栈顶遇到后面的第一个比它大的元素 i
		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
