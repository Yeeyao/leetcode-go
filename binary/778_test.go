package binary

import "fmt"

/*
N * N grid 水池方格中 grid[i][j] 表示坐标 (i,j) 的平台高度。现在下雨了，时间为 t 时水池的水位是 t。
两个方格之间水位同时联通的时候可以游过去，从坐标的左上平台 (0,0) 出发，需要最少耗时多久可以游到坐标 (N-1,N-1)
[题目](https://leetcode.com/problems/swim-in-rising-water/description/)

解空间是 [0,max(grid)]，简单的思路是一个个试，试试是否可行就是能力检测，实际上如果 x 不行则小于 x 的都不行
又变成 寻找最左满足 >= target，使用最左二分模板解决，这里解决的是每次选择多少时间的问题。

至于能否可达就是能力检测函数的计算了
*/

func swimInWater(grid [][]int) int {
	var maxHeight int
	// 找到最大的高度
	for _, v := range grid {
		for _, b := range v {
			if b > maxHeight {
				maxHeight = b
			}
		}
	}
	left, right := 0, maxHeight
	for left <= right {
		mid := left + (right-left)/2
		if testLink(grid, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// 能否到达右下角的判断二维网格 DFS，递归处理，对每个当前的坐标需要四个方向扩展遍历判断
// 当前坐标本身判断是否成立，比如坐标越界，同时需要记录下是否已经访问过
// 这里需要额外判断水位大小
func testLink(grid [][]int, height int) bool {
	return visit(0, 0, height, grid, map[string]interface{}{})
}

func visit(x, y, height int, grid [][]int, visited map[string]interface{}) bool {
	if x < 0 || y < 0 || x > len(grid)-1 || y > len(grid[0])-1 {
		return false
	}
	if grid[x][y] > height {
		return false
	}
	// 满足到达右下角，直接返回
	if x == len(grid)-1 && y == len(grid[0])-1 {
		return true
	}
	visitedStr := fmt.Sprintf("%d.%d", x, y)
	if _, ok := visited[visitedStr]; ok {
		return false
	}
	visited[visitedStr] = struct{}{}
	return visit(x-1, y, height, grid, visited) || visit(x+1, y, height, grid, visited) ||
		visit(x, y-1, height, grid, visited) || visit(x, y+1, height, grid, visited)
}
