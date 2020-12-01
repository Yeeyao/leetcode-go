package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("leetcode 832 Flipping an Image", func(t *testing.T) {
		input := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
		want := [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("leetcode 832 Flipping an Image2", func(t *testing.T) {
		input := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
		want := [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		lenAi := len(A[i])
		for j := 0; j < lenAi/2; j++ {
			rindex := lenAi - j - 1
			temp := A[i][j]
			A[i][j] = A[i][rindex]
			A[i][rindex] = temp
		}
		for j := 0; j < lenAi; j++ {
			if A[i][j] == 1 {
				A[i][j] = 0
			} else {
				A[i][j] = 1
			}
		}
	}
	return A
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
