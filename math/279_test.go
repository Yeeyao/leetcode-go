package math

import (
	"testing"
)

/*
279. Perfect Squares
Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...)
which sum to n.
给定正整数 n，找到合适的最少数量的平方数
使用回溯法 找到小于 n 的最大平方数
*/
func TestPro(t *testing.T) {
	t.Run("279. Perfect Squares", func(t *testing.T) {
		n := 12
		want := 3
		got := numSquares(n)
		if want != got {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("279. Perfect Squares2", func(t *testing.T) {
		n := 13
		want := 2
		got := numSquares(n)
		if want != got {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("279. Perfect Squares3", func(t *testing.T) {
		n := 1
		want := 1
		got := numSquares(n)
		if want != got {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("279. Perfect Squares4", func(t *testing.T) {
		n := 0
		want := 0
		got := numSquares(n)
		if want != got {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("279. Perfect Squares5", func(t *testing.T) {
		n := 9
		want := 1
		got := numSquares(n)
		if want != got {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	numSquares(n) = numSquares(n-k) + 1 k 属于 1, 4, 9, ...
	先构建小于等于 n 的平方数组 数组最大数值等于 n 或者 n = 1 就返回 1
	dp[0] 初始化为 0
	循环 i 从 1 到 n
		初始化 min = n
		循环 j 从 0 开始表示 平方数组的元素索引，从小于等于 i 所有满足 i - k
		的 dp 数中找到最小值，然后将它 + 1 作为当前的 dp 值
	最后返回 dp[n]
*/
func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	maxLess := 1
	var factor []int
	for ; maxLess*maxLess <= n; maxLess++ {
		factor = append(factor, maxLess*maxLess)
	}
	if n == 1 || maxLess*maxLess == n {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	factorLen := len(factor)
	// 对每个 dp 需要判断之前的元素 dp[i-1] dp[i-4] dp[i-9] ...
	for i := 1; i <= n; i++ {
		min := n
		for j := 0; j < factorLen && factor[j] <= i; j++ {
			if dp[i-factor[j]] < min {
				min = dp[i-factor[j]]
			}
		}
		dp[i] = min + 1
	}
	return dp[n]
}

/*
	贪心枚举
*/

/*
	TLE
*/
func numSquares10(n int) int {
	if n == 0 {
		return 0
	}
	maxLess := 1
	var factor []int
	for ; maxLess*maxLess <= n; maxLess++ {
		factor = append(factor, maxLess*maxLess)
	}
	if n == 1 || maxLess*maxLess == n {
		return 1
	}
	res := n
	helper(factor, len(factor)-1, n, 0, &res)
	return res
}

func helper(factor []int, start, n, temp int, count *int) {
	if n == 0 {
		if temp < *count {
			*count = temp
		}
		return
	}
	for i := start; i >= 0; i-- {
		if n-factor[i] >= 0 {
			n -= factor[i]
			temp++
			helper(factor, i, n, temp, count)
			n += factor[i]
			temp--
		} else {
			continue
		}
	}
}
