package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("62. Unique Paths", func(t *testing.T) {
		m, n := 3, 2
		want := 3
		got := solution(m, n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("62. Unique Paths2", func(t *testing.T) {
		m, n := 7, 3
		want := 28
		got := solution(m, n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	这个类似 64？使用 DP
	DP[i][j] = DP[i-1][j] + 1
	DP[i][j] = DP[i][j-1] + 1
同样 DP DP[i][j] = DP[i-1][j] + DP[i][j-1]
初始化 dp slice 遍历每个元素，判断当前元素坐标
如果 i 或者 j 等于 0 就直接等于非 0 的路径数量
*/
func solution(m, n int) int {
	// 保存当前的路径数量
	arr := make([][]int, m)
	for i, _ := range arr {
		arr[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				arr[i][j] = 1
			}
			// 边界上不会增加可行走路径
			if i == 0 && j != 0 {
				arr[i][j] = arr[i][j-1]
			}
			// 边界上不会增加可行走路径
			if i != 0 && j == 0 {
				arr[i][j] = arr[i-1][j]
			}
			if i != 0 && j != 0 {
				arr[i][j] = arr[i][j-1] + arr[i-1][j]
			}
		}
	}
	return arr[m-1][n-1]
}

func uniquePaths2(m int, n int) int {
	dp := make([][]int, m)
	for _, v := range dp {
		v = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[m-1][n-1]
}
