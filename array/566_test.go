package array

import (
	"testing"
)

func TestPro(t *testing.T) {

	t.Run("566. Reshape the Matrix", func(t *testing.T) {
		input := [][]int{{1, 2}, {3, 4}}
		r, c := 1, 4
		want := [][]int{{1, 2, 3, 4}}
		got := solution(input, r, c)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("566. Reshape the Matrix2", func(t *testing.T) {
		input := [][]int{{1, 2, 3}, {4, 5, 6}}
		r, c := 3, 2
		want := [][]int{{1, 2}, {3, 4}, {5, 6}}
		got := solution(input, r, c)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	首先检查数量是否满足要求
*/
func solution(input [][]int, r, c int) [][]int {
	inputRow := len(input)
	inputCol := len(input[0])
	TotalNum := inputRow * inputCol
	if TotalNum != r*c {
		return input
	}
	retArr := make([][]int, r)
	for i := 0; i < r; i++ {
		retArr[i] = make([]int, c)
	}
	for i := 0; i < inputRow*inputCol; i++ {
		retArr[i/c][i%c] = input[i/inputCol][i%inputCol]
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
