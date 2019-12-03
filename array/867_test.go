package array

import (
	"testing"
)

func TestPro(t *testing.T) {

	t.Run("867. Transpose Matrix", func(t *testing.T) {
		input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		want := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("867. Transpose Matrix2", func(t *testing.T) {
		input := [][]int{{1, 2, 3}, {4, 5, 6}}
		want := [][]int{{1, 4}, {2, 5}, {3, 6}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("867. Transpose Matrix3", func(t *testing.T) {
		input := [][]int{{1, 2, 3}}
		want := [][]int{{1}, {2}, {3}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input [][]int) [][]int {
	col := len(input[0])
	row := len(input)
	retArr := make([][]int, col)
	for i := 0; i < col; i++ {
		retArr[i] = make([]int, row)
	}
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			retArr[j][i] = input[i][j]
		}
	}
	return retArr
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
