package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("48. Rotate Image", func(t *testing.T) {
		input := []int{1, 3, 5, 6}
		target := 2
		want := 1
		got := solution(input, target)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
DO NOT allocate another 2D matrix and do the rotation.

给定一个 n * n 二维矩阵， 直接将矩阵顺时针旋转 90 度，只能对输入的矩阵进行修改
两次反转 先水平 180 反转 再左上到右下反转
*/
func rotate(matrix [][]int) {
	row := len(matrix)
	// 180 反转
	for i := 0; i < row/2; i++ {
		matrix[i], matrix[row-i-1] = matrix[row-i-1], matrix[i]
	}
	// 左上右下对角线反转
	for i := 0; i < row; i++ {
		for j := i + 1; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
