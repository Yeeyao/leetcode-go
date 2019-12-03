package array

import "testing"

func TestPro(t *testing.T) {

	t.Run("766. Toeplitz Matrix", func(t *testing.T) {
		input := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("766. Toeplitz Matrix2", func(t *testing.T) {
		input := [][]int{{1, 2}, {2, 2}}
		want := false
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("766. Toeplitz Matrix3", func(t *testing.T) {
		input := [][]int{{1}, {2}}
		want := false
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(matrix [][]int) bool {
	row := len(matrix)
	col := len(matrix[0])
	// 最后一行不需要和下一行比较
	for i := 0; i < row-1; i++ {
		// 最后一列也不需要和下一行比较
		for j := 0; j < col-1; j++ {
			if matrix[i][j] == matrix[i+1][j+1] {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func solution(matrix [][]int) bool {
	row := len(matrix)
	col := len(matrix[0])
	// 最后一行不需要和下一行比较
	for i := 0; i < row-1; i++ {
		// 最后一列也不需要和下一行比较
		for j := 0; j < col-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}
