package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("59. Spiral Matrix II", func(t *testing.T) {
		input := 3
		want := [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("59. Spiral Matrix II2", func(t *testing.T) {
		input := 3
		want := [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}
		got := solution2(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定 n 需要生成从 1 到 n^2 的螺旋矩阵
	难度本身不是生成，是每个元素存放的位置判定
	生成的矩阵是 n * n 的 目标行和列计算
	找出边界元素值对应的坐标关系，然后每次循环，相当于 n 减小 2

	从四周向中间处理
	k 为初始值，i 为第几次遍历完四周，j 控制坐标变化。
	四周的行或者列需要数值判断
	第一行的左到右 这里行不用变化，列变化，后面的同理
	同时注意是按照数值大小顺序来处理遍历
	[理解意思，但是代码很繁琐](https://leetcode.com/problems/spiral-matrix-ii/discuss/22309/Simple-C%2B%2B-solution(with-explaination))
*/
func solution(n int) [][]int {
	retArr := make([][]int, n)
	for i, _ := range retArr {
		retArr[i] = make([]int, n)
	}
	// 第几次对四周的处理，数组起始数值
	i, k := 0, 1
	for k <= n*n {
		j := i
		// 第一行 注意这里每次大循环，行号会增加，列数减少
		// 单次循环行号不变，列号增加
		for j < n-i {
			retArr[i][j] = k
			j++
			k++
		}
		// 最后一列 注意这里每次大循环，列号会减小，行数会减少
		// 上面 j 变成了 n - i
		// 单次循环列号不变，行号增加
		j = i + 1
		for j < n-i {
			retArr[j][n-i-1] = k
			j++
			k++
		}
		// 上面的 j 变成了 n - i
		// 最后一行 注意这里每次大循环，行号会减小，列数会减少
		// 单次循环行号不变，列号减小
		// 注意这里要减一次
		j = n - i - 2
		for j > i {
			retArr[n-i-1][j] = k
			j--
			k++
		}
		// 上面的 j 变成 0
		// 第一列 注意这里每次大循环，行号会增加，列数会减少
		// 单次循环列号不变，行号减小
		j = n - i - 1
		for j > i {
			retArr[j][i] = k
			j--
			k++
		}
		i++
	}
	return retArr
}

/*
	跟上面思路一样，但是循环简单一些
*/
func solution2(n int) [][]int {
	retArr := make([][]int, n)
	for i, _ := range retArr {
		retArr[i] = make([]int, n)
	}
	i, j, di, dj := 0, 0, 0, 1
	k := 1
	for k <= n*n {
		// 初始是 0 这里就 + 1
		retArr[i][j] = k
		// 注意，这里负数取余，需要得到正数，因为是作为数组的索引
		cx := (i + di) % n
		if cx < 0 {
			cx = -cx
		}
		cy := (j + dj) % n
		if cy < 0 {
			cy = -cy
		}
		// 遇到需要变换行或者列的情况
		if retArr[cx][cy] != 0 {
			di, dj = dj, -di
		}
		i += di
		j += dj
		k++
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
