package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("712. Minimum ASCII Delete Sum for Two Strings", func(t *testing.T) {
		input := []int{10, 9, 2, 5, 3, 7, 101, 18}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}

/*
	给定两个字符串 s1, s2 返回最小的 ASCII 和，使得将两个字符串通过删除字母变成相同。仅包含小写字母
	[ref](https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/discuss/642422/for-those-who-have-no-clue-%3A-step-by-step)
	动态规划？
	1. 明白问题: 要求只能删除字符串的字母让两个字符串相等，并且删除的字母的 ASCII 总和最小
	2. 逻辑构建以及递归关系
		2.1 分析不同的情况并逻辑构建
			存在两种情况
			1. 一些 ith jth 的字母在两个字符串中匹配 2. 它们都不匹配
			分析：
			1.如果一些 ith 和 jth 字母匹配，则我们可以在序列中包含它们，因此不需要将它们添加到 ASCII 总和中
			2.如果 ith 和 jth 不匹配，我们有两个选择
				1. 跳过 ith 字母，假定 jth 字母在后面有用，因此我们将 ith 的 ASCII 加到总和并对剩下的递归处理
				2. 跳过 jth 字母，假定 ith 字母在后面有用，因此我们将 jth 的 ASCII 加到总和并对剩下的递归处理
		2.2 获得基本情况并写递归关系
			基本：如果 A 或者 B 是空的，则总和是非空的那个的所有字母 如果相同，则总和是 0
			递归关系：
				如果 string_a[ith] 等于 string_b[jth] 可以让字符串相同，因此跳过 sum = get_sum_for(string_a ith+1, string_b, jth+1)
				如果不想等，则
					1. sum = string_a[ith] 的 ASCII + get_sum_for(string_a ith+1, string_b, jth) // 如果 jth 有用
					2. sum = string_b[jth] 的 ASCII + get_sum_for(string_a ith, string_b, jth+1) // 如果 ith 有用
				ans = min(option1, option2)
*/

/*
	bottom-up dp
	dp[i][j] 表示遍历到 s1 的第 i 个字母，s2 的第 j 个字母的时候，当前删除的 ASCII 总和。最终结果是 dp[len(a)[len(b)]
	dp[0][0] = 0
	dp[i][j] = min(dp[i][j-1]+a[j-1], dp[i-1][j]+a[i-1])
*/
func solution(s1, s2 string) int {
	n, m := len(s1), len(s2)
	dp := make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, m+1)
	}
	// 基本情况的初始化，这里计算其中一个字符串为空的时候，另一个字符串的总和
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for j := 1; j <= m; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// 如果当前的字母相等，就不需要加上和
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}
	return dp[n][m]
}

/*
	下面的方法，计算 sub 的时候，存在重复的计算，因此使用一个数据结构保存已经计算的结果
	top-down dynamic
*/
func solution2(s1, s2 string) int {
	n, m := len(s1), len(s2)
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, m)
	}
	return sub2(s1, s2, 0, 0, dp)
}

func sub2(a, b string, i, j int, dp [][]int) int {
	n, m := len(a), len(b)
	sum := 0
	// 其中一个字符串遍历完了
	if i == n || j == m {
		// 两个都遍历完，需要删除的和就是 0
		if i == n && j == m {
			return 0
		}
		if i == n {
			return dead_end_sum(b, j)
		} else {
			return dead_end_sum(a, i)
		}
	}
	// 这里使用 dp 数组来判断是否已经计算过这个组合 a,b,i,j
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	// 递归的第一种情况，元素相等，直接计算后面的
	if a[i] == b[j] {
		sum = sub2(a, b, i+1, j+1, dp)
		// 不相等，计算两种情况 使用 i 或者使用 j，取较小者
	} else {
		sum = min(sub2(a, b, i+1, j, dp)+int(a[i]), sub2(a, b, i, j+1, dp)+int(b[j]))
	}
	dp[i][j] = sum
	return sum
}

/*
	TLE
*/
func solution3(s1, s2 string) int {
	return sub3(s1, s2, 0, 0)
}

// 如果另一个字符串为空，则计算剩余的字符串的字母总和
func dead_end_sum(s string, i int) int {
	sum := 0
	for i < len(s) {
		sum += int(s[i])
		i++
	}
	return sum
}

func sub3(a, b string, i, j int) int {
	n, m := len(a), len(b)
	sum := 0
	// 其中一个字符串遍历完了
	if i == n || j == m {
		// 两个都遍历完，需要删除的和就是 0
		if i == n && j == m {
			return 0
		}
		if i == n {
			return dead_end_sum(b, j)
		} else {
			return dead_end_sum(a, i)
		}
	}
	// 递归的第一种情况，元素相等，直接计算后面的
	if a[i] == b[j] {
		sum = sub3(a, b, i+1, j+1)
		// 不相等，计算两种情况 使用 i 或者使用 j，取较小者
	} else {
		sum = min(sub3(a, b, i+1, j)+int(a[i]), sub3(a, b, i, j+1)+int(b[j]))
	}
	return sum
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
