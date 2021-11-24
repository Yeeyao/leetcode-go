package array

import (
	"math"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("491. Increasing Subsequences", func(t *testing.T) {
		input := []int{10, 9, 2, 5, 3, 7, 101, 18}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}

/*
给定一个整型数组 nums，返回所有的至少含有两个元素的不同的上升子序列
nums 可能含有重复的元素，两个相同的整型可以认为是特殊的上升子序列

类似排列的题目，可以直接遍历处理，
子函数
	从第一个开始到最后一个
	先将当前的加到临时结果中，递归处理，然后删除，继续遍历
[ref](https://leetcode-cn.com/problems/increasing-subsequences/solution/di-zeng-zi-xu-lie-by-leetcode-solution)

方法 1 是直接二进制枚举 + 哈希 对序列中每个数字，有两种状态，被选中或者不被选中。因此，长度为 n 有 2^n 种序列，然后使用哈希算法去重
方法 2 递归枚举 + 剪枝，递归方法代替二进制枚举。
	通用模板
		使用临时数组 temp 保存当前已经选择的子序列，cur 表示当前选择的位置的下标
		在 dfs(cur, nums) 开始之前，[0,cur−1] 这个区间内的所有元素都已经被考虑过，[cur, n] 这个区间的元素还没有被考虑
		执行 dfs(cur, nums) 时，考虑 cur 选不选，
			如果选，就加入到 temp，然后递归下一个位置，递归前需要将当前的元素从 temp 删除
			如果不选，就直接递归下一个位置
	额外条件：需要合法以及不重复
		合法：要求当前元素大于上一个元素才能选择
		没有重复：当前元素等于上一个选择元素，不能考虑当前元素，需要直接递归后面的元素，出现重复，有四种情况
			前者被选择，后者被选择
			前者被选择，后者不被选择
			前者不被选择，后者被选择
			前者不被选择，后者不被选择
		中间两种情况是等价的，限制之后，保留了 3 舍弃了 2 可以去重
		时间复杂度：O(2^n * n)
*/
var (
	temp []int
	ans  [][]int
)

func solution(nums []int) [][]int {
	ans = [][]int{}
	dfs(0, math.MinInt32, nums)
	return ans
}

// 这里同时记录临时结果的最后一个元素的数值
func dfs(cur, last int, nums []int) {
	// 到达结尾了，可以直接返回了
	if cur == len(nums) {
		if len(temp) >= 2 {
			t := make([]int, len(temp))
			copy(t, temp)
			ans = append(ans, t)
		}
		return
	}
	// a
	// 这里本身就已经对 cur 进行选择与不选择的判断了（需要满足大于的条件下）, b 就是针对重复元素进行的处理
	// 满足合法的要求，这里是一定选择的情况
	// 没有重复的情况 1，因为 b 的条件一定不满足，
	// 同时，在 b 满足的情况下，下一次递归在，这里 a 就满足了，因此这里也是没有重复的情况 3
	if nums[cur] >= last {
		temp = append(temp, nums[cur])
		dfs(cur+1, nums[cur], nums)
		temp = temp[:len(temp)-1]
	}
	// b
	// 当前元素不等于上一个元素，下一个相同的元素也会满足，因此这里是没有重复的情况 4
	// 本次不选择之后，因为上面的条件，下次会选择
	if nums[cur] != last {
		dfs(cur+1, last, nums)
	}
}
