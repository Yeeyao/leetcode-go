package array

import (
	"testing"
)

/*
200. Number of Islands
Given a 2d grid map of '1's (land) and '0's (water), count the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.
给定二维数组，其中 1 表示陆地，0 表示水，计算陆地的数量
陆地被水包围然后相连的陆地还是算同一块
直接遍历每个元素，然后将当前元素是 1 就递增然后直接递归调用其周围的元素，每次遍历将元素置 0
最后返回总数
*/

func TestPro(t *testing.T) {
	t.Run("200. nNumber of Islands    ", func(t *testing.T) {
		input := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
		want := 2
		got := numIslands(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	同样遍历每个元素，辅助函数先判断索引合法性，如果非法就直接返回
	如果当前索引元素为 1 就置 0 然后直接递归调用
*/
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				helper(grid, i, j)
			}
		}
	}
	return res
}

func helper(grid [][]byte, i, j int) {
	if i < 0 || i > len(grid)-1 {
		return
	}
	if j < 0 || j > len(grid[0])-1 {
		return
	}
	if grid[i][j] == '1' {
		grid[i][j] = '0'
		helper(grid, i, j+1)
		helper(grid, i, j-1)
		helper(grid, i+1, j)
		helper(grid, i-1, j)
	}
}

func numIslands2(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	// row 4 col 5
	row := len(grid)
	col := len(grid[0])
	res := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				res++
				helper2(grid, i, j, row, col)
			}
		}
	}
	return res
}

/*
	当前元素以及周围四个元素处理
*/
func helper2(grid [][]byte, i, j, row, col int) {
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	if i > 0 {
		helper2(grid, i-1, j, row, col)
	}
	if i < row-1 {
		helper2(grid, i+1, j, row, col)
	}
	if j > 0 {
		helper2(grid, i, j-1, row, col)
	}
	if j < col-1 {
		helper2(grid, i, j+1, row, col)
	}
}
