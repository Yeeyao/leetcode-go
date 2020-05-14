package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("64. Minimum Path Sum", func(t *testing.T) {
		input := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
		want := 7
		got := solution2(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定二维数组，找从左上角到右下角的路径，只能向下或者向右移动
	使得路径上的总和最小
	一点优化，得到当前的最小路径后，如果一次移动中，路径大于当前最小路径，那就应该直接停止了
	DP?递归?回溯
	从(0,0)到(m-1,n-1)
	对于点(i, j) 有两条路可以到这里 (i-1,j), (i, j-1)
	i,j 可以走的点是 i + 1, j 或者 i, j + 1
	难道用回溯法来穷举？回溯法实锤了 但是这里 tle 了
	DP[i][j] = min(DP[i-1][j], DP[i][j-1]) + grid[i][j]
	DP 了
*/
func solution(grid [][]int) int {
	const intMax = int(^uint(0) >> 1)
	minSum := intMax
	m := len(grid)
	n := len(grid[0])
	solutionHelper(grid, &minSum, m, n, grid[0][0], 0, 0)
	return minSum
}

func solutionHelper(grid [][]int, minSum *int, m, n, tempSum, x, y int) {
	// 终止条件是到达右下角或者当前的路径总和大于目前的最小和
	if tempSum > *minSum {
		return
	}
	if x == m-1 && y == n-1 {
		if tempSum < *minSum {
			*minSum = tempSum
		}
	}
	// 每个点有两条路径，需要判断边界
	if x < m-1 {
		tempSum += grid[x+1][y]
		solutionHelper(grid, minSum, m, n, tempSum, x+1, y)
		// 需要减去，避免下面的加多了
		tempSum -= grid[x+1][y]
	}
	if y < n-1 {
		tempSum += grid[x][y+1]
		solutionHelper(grid, minSum, m, n, tempSum, x, y+1)
	}
}

/*
	DP[i][j] = min(DP[i-1][j], DP[i][j-1]) + grid[i][j]
	DP 了
*/
func solution2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 边界处理
			if i == 0 && j != 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			}
			if i != 0 && j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			}
			if i != 0 && j != 0 {
				if grid[i-1][j] > grid[i][j-1] {
					grid[i][j] = grid[i][j-1] + grid[i][j]
				} else {
					grid[i][j] = grid[i-1][j] + grid[i][j]
				}
			}
		}
	}
	return grid[m-1][n-1]
}
