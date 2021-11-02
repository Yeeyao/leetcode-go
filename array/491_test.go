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
	递归枚举子序列的通用模板，即用一个临时数组 temp 来保存当前选出的子序列，使用 cur 来表示当前位置的下标，在 dfs(cur, nums) 开始之前，[0,cur−1] 这个区间内的所有元素都已经被考虑过，
	而 [cur,n] 这个区间内的元素还未被考虑。在执行 dfs(cur, nums) 时，我们考虑 cur 这个位置选或者不选，如果选择当前元素，那么把当前元素加入到 temp 中，
	然后递归下一个位置，在递归结束后，应当把 temp 的最后一个元素删除进行回溯；如果不选当前的元素，直接递归下一个位置。
	限制条件来枚举出合法且不重复
		合法则在选择的时候，当前元素大于上一个元素才能选择这个元素
		没有重复，只有当前元素不等于上一个选择元素的时候，才考虑不选择当前元素，直接递归后面的元素。有两个相同的元素，有四种情况
			前者被选择，后者被选择
			前者被选择，后者不被选择
			前者不被选择，后者被选择
			前者不被选择，后者不被选择
	二三种情况是等价的，限制之后，保留了 3 舍弃了 2 可以去重
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

func dfs(cur, last int, nums []int) {
	if cur == len(nums) {
		if len(temp) >= 2 {
			t := make([]int, len(temp))
			copy(t, temp)
			ans = append(ans, t)
		}
		return
	}
	if nums[cur] >= last {
		temp = append(temp, nums[cur])
		dfs(cur+1, nums[cur], nums)
		temp = temp[:len(temp)-1]
	}
	if nums[cur] != last {
		dfs(cur+1, last, nums)
	}
}
