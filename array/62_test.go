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
