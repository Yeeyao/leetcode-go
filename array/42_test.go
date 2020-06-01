package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("42. Trapping Rain Water", func(t *testing.T) {
		nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		want := 6
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("42. Trapping Rain Water2", func(t *testing.T) {
		nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		want := 6
		got := solution2(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("42. Trapping Rain Water3", func(t *testing.T) {
		nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		want := 6
		got := solution3(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("42. Trapping Rain Water4", func(t *testing.T) {
		nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		want := 6
		got := solution4(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	雨水填充 计算填充了多少水 这题好有创意啊
	使用双指针处理 从开头遍历，先找第一个高点，然后找第二个高点，之后中间的空位
	这里的高点是找上一个比当前的点高的点，考虑使用单调递增栈
	就是填充的水，第二个高点赋值给第一个，第二个继续遍历下去
	这里的高点，应该是单调递增的最高点
	水量的计算是两个高点的较小者的高度 * 两个点的距离 - 中间的非空

	先判断哪个高点需要移动

	dp 首先从开头遍历一次，保存到遍历位置前的最高高度
	之后从尾部向前遍历，保存遍历位置后面的最大值，
	取两者的最小值跟当前位置的高度比较
	水量就是 dp[i] - height[i] 的总和
	这里的思路就是将当前位置的左右两边的较小值来比较
*/
func solution(height []int) int {
	hLen := len(height)
	dp := make([]int, hLen)
	res, mx := 0, 0
	// 先保存遍历位置前的最高高度 先更新 dp，再更新最高高度
	for i := 0; i < hLen; i++ {
		dp[i] = mx
		// 更新当前最高高度
		if height[i] > mx {
			mx = height[i]
		}
	}
	// 将最高高度重置 这里直接使用这个来判断
	mx = 0
	for i := hLen - 1; i >= 0; i-- {
		// 保存的高度与当前的最高高度取较小值并比较当前的高度
		if mx < dp[i] {
			dp[i] = mx
		}
		// 更新当前最高高度
		if height[i] > mx {
			mx = height[i]
		}
		// 更新当前最高高度 比较当前高度判断是否需要加
		if dp[i] > height[i] {
			res += dp[i] - height[i]
		}
	}
	return res
}

/*
	双指针扫描 如果左边的较小，则继续向右遍历一位，然后从当前位置开始小于最小值则
	不断将结果累加，右边较小则同样处理
	因为这里的逻辑很清楚，只需要知道较小值然后不断处理，不关心较大值
*/
func solution2(height []int) int {
	hLen := len(height)
	res, l, r := 0, 0, hLen-1
	temp := 0
	for l < r {
		// 左边较小，左移一位后不断判断并左移
		if height[l] < height[r] {
			temp = height[l]
			l++
			for l < r && height[l] < temp {
				res += temp - height[l]
				l++
			}
			// 右边较小
		} else {
			temp = height[r]
			r--
			for l < r && height[r] < temp {
				res += temp - height[r]
				r--
			}
		}
	}
	return res
}

/*
	solution2 [优化代码](https://leetcode.com/problems/trapping-rain-water/discuss/17364/7-lines-C-C%2B%2B)
	不断更新较小值 lower 以及当前的较小值中的最大值 lv
*/
func solution3(height []int) int {
	hLen := len(height)
	res, l, r := 0, 0, hLen-1
	lower, lv := 0, 0
	for l < r {
		if height[l] < height[r] {
			lower = height[l]
			l++
		} else {
			lower = height[r]
			r--
		}
		if lower > lv {
			lv = lower
		}
		res += lv - lower
	}
	return res
}

/*
	stack solution 这里使用 单调递减栈
	一开始将所有小于当前栈顶的元素索引值入栈
	如果当前的元素值大于栈顶，则需要将栈顶元素出栈并计算水位
	比较出栈后的栈顶以及当前元素，取两者的较小值作为装水的最低水位再计算
*/
func solution4(height []int) int {
	hLen := len(height)
	if hLen == 0 {
		return 0
	}
	st := make([]int, hLen)
	stTop, i, res := 0, 0, 0
	for i < hLen {
		// 空或者当前元素小于栈顶元素，元素索引值入栈 注意这里栈顶的值
		if stTop == 0 || height[i] <= height[st[stTop-1]] {
			st[stTop] = i
			i++
			stTop++
		} else {
			maxBw := 0
			if stTop-1 == 0 {
				maxBw = 0
			} else {
				// 比较当前的高度和出栈后的元素的高度，取较小值
				if height[i] < height[st[stTop-2]] {
					maxBw = (height[i] - height[st[stTop-1]]) * (i - st[stTop-2] - 1)
				} else {
					maxBw = (height[st[stTop-2]] - height[st[stTop-1]]) * (i - st[stTop-2] - 1)
				}
			}
			// 出栈
			stTop--
			res += maxBw
		}
	}
	return res
}
