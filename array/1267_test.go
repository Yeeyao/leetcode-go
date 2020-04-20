package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1267. Count Servers that Communicate", func(t *testing.T) {
		input := [][]int{{1, 0}, {0, 1}}
		want := 0
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1267. Count Servers that Communicate2", func(t *testing.T) {
		input := [][]int{{1, 0}, {1, 1}}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1267. Count Servers that Communicate3", func(t *testing.T) {
		input := [][]int{{1, 0, 0, 1, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1267. Count Servers that Communicate4", func(t *testing.T) {
		input := [][]int{{1, 0, 0, 1, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	不需要两个节点相邻，只需要在一条直线
	遍历一列，然后找非 0 节点数量，如果等于 1，将所有节点都置 2 并计数
	对于第一个节点，需要记录下来，如果该行的非 0 节点数量大于 1 则需要将剩余的节点都置 2 并将第一个节点置 2
	遍历一行，然后找非 0 节点数量，如果等于 1，将节点置 2，如果等于 2 跳过处理
*/
func solution(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	count := 0
	for i := 0; i < m; i++ {
		colCount := 0
		fc, fr := -1, -1
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if fc < 0 || fr < 0 {
					fc, fr = i, j
				}
				colCount++
				// 该行 1 节点数量大于 1 才需要将节点置 2
				if colCount > 1 {
					grid[i][j] = 2
				}
			}
		}
		// 第一个节点也需要处理，同时，数量大于 1 才能计数
		if colCount > 1 {
			grid[fc][fr] = 2
			count += colCount
		}
	}
	for i := 0; i < n; i++ {
		rowCount := 0
		dupCount := 0
		for j := 0; j < m; j++ {
			if grid[j][i] == 1 {
				rowCount++
			}
			if grid[j][i] == 2 {
				dupCount++
			}
		}
		if rowCount > 1 {
			count += rowCount
		}
		if rowCount == 1 && dupCount > 0 {
			count++
		}
	}
	return count
}

/*
	类似之前有一题从第二行和第二列开始，倒数第二的行和列为止，遍历所有 1 的点，
	检查其上下左右四个位置是否为 1，如果是需要将其计数
	并改成2，包括当前访问的点 下面的是理解错意思了，不需要两个节点相邻，只需要在一条直线
*/
func solution2(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	count := 0
	// 中间部分
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 1 {
				if grid[i-1][j] == 1 {
					count += 2
					grid[i-1][j] = 2
					grid[i][j] = 2
				}
				if i < m-1 && grid[i+1][j] == 1 {
					count += 2
					grid[i+1][j] = 2
					grid[i][j] = 2
				}
				if grid[i][j-1] == 1 {
					count += 2
					grid[i][j-1] = 2
					grid[i][j] = 2
				}
				if j < n-1 && grid[i][j+1] == 1 {
					count += 2
					grid[i][j+1] = 2
					grid[i][j] = 2
				}
			}
		}
	}
	// 第一行 看左右
	for i := 0; i < n-1; i++ {
		if grid[0][i] == 1 {
			if i > 0 && grid[0][i-1] > 0 {
				count++
				// 两个都是 1 的时候，需要加上两个
				if grid[0][i-1] == 1 {
					count++
				}
				grid[0][i-1] = 2
				grid[0][i] = 2
			}
			if grid[0][i+1] > 0 {
				count++
				if grid[0][i+1] == 1 {
					count++
				}
				grid[0][i+1] = 2
				grid[0][i] = 2
			}
		}

	}
	// 第一列 看上下
	for j := 0; j < m-1; j++ {
		if grid[j][0] == 1 {
			if j > 0 && grid[j-1][0] > 0 {
				count++
				if grid[j-1][0] == 1 {
					count++
				}
				grid[j-1][0] = 2
				grid[j][0] = 2
			}
			if grid[j+1][0] > 0 {
				count++
				if grid[j+1][0] == 1 {
					count++
				}
				grid[j+1][0] = 2
				grid[j][0] = 2
			}
		}
	}
	return count
}
