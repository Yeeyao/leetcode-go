package array

/*
221. Maximal Square 类似 200 1277
Given a 2D binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.
但是这里要求的是正方形面积
*/

/*
	[ref](https://leetcode-cn.com/problems/maximal-square/solution/zui-da-zheng-fang-xing-by-leetcode-solution/)
[ref](https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones/solution/tong-ji-quan-wei-1-de-zheng-fang-xing-zi-ju-zhen-2/)
	先计算行列数量，然后直接将输入 matrix 复制到 dp 中
	第二个循环从第二列，第二行遍历 dp 每个元素
		如果 dp[i][j] == 1 就用公式更新
		这里每次更新最长边长
	这里前面是一样的，只需要每次更新 copy 元素的时候更新一下最大边长，最后计算面积

	1 1 1
	1 2 2
    1 2 3
*/
func maximalSquare(matrix [][]byte) int {
	// m, n 行 列
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	res := 0
	dp := make([][]int, m)
	// 这里全都复制到 dp
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				res = 1
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}

	return res * res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}
