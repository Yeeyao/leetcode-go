package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("1277. Count Square Submatrices with All Ones", func(t *testing.T) {
		input := [][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}
		want := 15
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("1277. Count Square Submatrices with All Ones2", func(t *testing.T) {
		input := [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}
		want := 7
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	化简一下方法
	先计算行列数量，然后直接将输入 matrix 复制到 dp 中
	i, j 两个循环遍历 matrix 每个元素，
		如果 i 或者 j 等于 0 dp[i][j] = matrix[i][j]
		如果 matrix[i][j] == 0 dp[i][j] = 0
		如果 matrix[i][j] == 1 则套用公式，左边，上面，左上角三个的最小值 + 1
		计算一个 dp[i][j] 就累加到正方形数量中
*/
func solution(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = matrix[i][j]
			} else if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
			}
			res += dp[i][j]
		}
	}
	return res
}

/*
	直接遍历，统计所有的正方形
	之后，针对每个正方形，计算其包含的正方形数量
	每个小正方形的计算是
	n > 1 1 + 4 * ( n - 2 ) + n * n
	n = 1 1

	brute force 是，统计每个大小的正方形的数量
	但是，思路是从大的往小的方向统计，然后每个大的就会包含小的，计算出该正方形的数量

	复制一个 copy 二维数组 先从 matrix 中复制最左边一列和最上面一行
	从第二列第二行开始遍历 matrix 每个元素，当前元素数值是 1，copy 对应的当前元素数值是上，左，左上三个元素的 copy 最小值 + 1
	最后遍历一次 copy 数组，将所有元素数值累计作为结果
*/
func solution2(matrix [][]int) int {
	var num int

	m := len(matrix)
	n := len(matrix[0])

	// allocation
	copy := make([][]int, m)
	for i := range copy {
		copy[i] = make([]int, n)
	}

	// copy the left col and top row elements
	for i := 0; i < m; i++ {
		copy[i][0] = matrix[i][0]
	}

	for j := 0; j < n; j++ {
		copy[0][j] = matrix[0][j]
	}

	// 这里的原理是从第二行和第二列开始遍历
	// 如果这个位置是 1 则需要判断周围三个的值，如果三个是 1
	// 则这个位置的值为 2，否则为 1
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 1 {
				// self and the other three
				copy[i][j] = min(copy[i-1][j-1], min(copy[i][j-1], copy[i-1][j])) + 1
			}
		}
	}

	// 遍历并直接将所有非 0 相加
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			num += copy[i][j]
		}
	}

	return num
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
