package array

import (
	"testing"
)

func TestPro(t *testing.T) {

	t.Run("1260. Shift 2D Grid", func(t *testing.T) {
		grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		k := 1
		want := [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}}
		got := solution(grid, k)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1260. Shift 2D Grid2", func(t *testing.T) {
		grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		k := 1
		want := [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}}
		got := solution(grid, k)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(grid [][]int, k int) [][]int {
	row := len(grid)
	col := len(grid[0])
	total := row * col
	if k == total {
		return grid
	}
	indent := k % total
	retArr := make([][]int, row)
	for i := 0; i < row; i++ {
		retArr[i] = make([]int, col)
	}
	for i := 0; i < total; i++ {
		var indent2 int
		if i-indent < 0 {
			indent2 = total + i - indent
		} else {
			indent2 = i - indent
		}
		retArr[i/col][i%col] = grid[indent2/col][indent2%col]
	}
	return retArr
}

// 看看有什么不同
func shiftGrid(grid [][]int, k int) [][]int {
	totalRow := len(grid)
	totalCol := len(grid[0])
	res := make([][]int, totalRow)
	for i, _ := range(res) {
		res[i] = make([]int, totalCol)
	}

	for r, row := range(grid) {
		for c, col := range(row) {
			indx := r*totalCol + c + k
			newRow := (indx / totalCol) % totalRow
			newCol := indx % totalCol
			res[newRow][newCol] = col
		}
	}
	return res
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
