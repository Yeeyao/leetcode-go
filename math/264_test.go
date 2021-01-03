package math

import "testing"

func TestPro(t *testing.T) {
	t.Run("264. Ugly Number II", func(t *testing.T) {
		num := 9
		want := true
		got := solution(num)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
264. Ugly Number II
Write a program to find the n-th ugly number.
Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.
找到第 n 个丑数
我记得做过吧 剑指 offer 33

动态规划
使用三个指针和一个数组，然后三个指针初始化位置都是 0，建立一个 dp 数组，长度 n, dp[0] = 1
每次将三个指针对应 * 2， * 3， * 5 并计算结果的最小值，然后更新 dp[i] 之后判断，如果当前指针的结果等于 dp[i]
就需要将指针移动到下一个位置

*/
func solution(n int) int {
	dp := make([]int, n)
	a, b, c := 0, 0, 0
	dp[0] = 1
	for i := 1; i < n; i++ {
		n1, n2, n3 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(n1, min(n2, n3))
		if dp[i] == n1 {
			a++
		}
		if dp[i] == n2 {
			b++
		}
		if dp[i] == n3 {
			c++
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
