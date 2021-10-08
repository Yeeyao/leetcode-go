package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("377. Combination Sum IV", func(t *testing.T) {
		input := []int{10, 1, 2, 7, 6, 1, 5}
		target := 8
		want := 9
		got := solution(input, target)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
Given an array of distinct integers nums and a target integer target, return the number of possible combinations that
add up to target.
The answer is guaranteed to fit in a 32-bit integer.

给定一个含有唯一整型的 nums 数组以及一个 target，返回子数组和是 target 的数量，不同的序列看作是不同的组合
相比 39 这里允许结果数组中存在相同的元素，但是排列不同，在 39 的结果上，对每个数组构造排列？
问题变成了构造排列 permutation，但是题目描述是组合 combination。。。

[ref](https://leetcode-cn.com/problems/combination-sum-iv/solution/zu-he-zong-he-iv-by-leetcode-solution-q8zv/)
直接 dp 了

定义 dp[i] 为 target 为 i 时的排列数量 dp[0] = 1 表示不选取任何元素的时候，和为 0 只有一种方案
1 <= i <= target 如果存在一种排列，其中元素和为 i，则排列的最后一个元素一定是数组 nums 中的一个元素，设该元素为 num，则 num <= i
对元素之和 i - num 的每种排列，在最后添加 num 之后就可以得到元素和为 i 的排列，因此计算 dp[i] 的时候，应该计算所有 dp[i-num]
因此得到动态规划做法
	初始化 dp[0] = 1
	遍历 i 从 1 到 target，对每个 i
		- 遍历数组 nums 中每个元素 num，当 num <= i 时，将 dp[i - num] 的值加到 dp[i]
	最终 dp[target] 就是答案

上面是否考虑了选取元素的顺序？是的。外层循环遍历 1 到 target 的值，内层循环遍历数组 nums 的值，计算 dp[i] 的值时，nums 中每个小于等于 i
的元素都可能作为元素之和等于 i 的排列的最后一个元素。例如，1，3 都在 nums 中，计算 dp[4] 的时候 dp[1] 和 dp[3] 都被考虑到了
*/
func solution(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
