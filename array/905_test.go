package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("leetcode 832 Sort Array By Parity", func(t *testing.T) {
		input := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
		want := [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(A []int) []int {
	var odd, even []int
	for i := 0; i < len(A); i++ {
		// 这里不要自作聪明直接位与了，让编译器优化
		if A[i]%2 == 0 {
			even = append(even, A[i])
		} else {
			odd = append(odd, A[i])
		}
	}
	return append(even, odd...)
}

func IntSliceEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
