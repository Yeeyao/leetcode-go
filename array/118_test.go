package array

import (
	"testing"
)

func TestPro(t *testing.T) {

	t.Run("118. Pascal's Triangle", func(t *testing.T) {
		input := 5
		want := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("118. Pascal's Triangle2", func(t *testing.T) {
		input := 3
		want := [][]int{{1}, {1, 1}, {1, 2, 1}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}

func solution(numRows int) [][]int {
	retArr := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		retArr[i] = make([]int, i+1)
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				retArr[i][j] = 1
			} else if i > 1 {
				retArr[i][j] = retArr[i-1][j-1] + retArr[i-1][j]
			}
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
