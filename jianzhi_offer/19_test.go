package jianzhi_offer

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("19 顺时针打印矩阵", func(t *testing.T) {
		nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		want := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
		get := solution(nums)
		if !reflect.DeepEqual(get, want) {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	同 LeetCode 54
	先是第一行，然后最后一列，然后第一列，然后第二行
	每次遍历一行，需要将当前行数递增，然后遍历列数，一开始是最大列数 - 当前列数，然后遍历完最左边的列数就递增

	1 2 3  1  2  3  4
	4 5 6  5  6  7  8
	7 8 9  9 10 11 12
          13 14 15 16

	按层模拟
	左上角开始从左到右遍历(top, left)到(top, right)
	从上到下遍历(top+1, right) 到 (bottom, right)
	若 left < right 且 top < bottom 从右到左遍历 (bottom, right - 1) 到 (bottom, left + 1)，从下到上遍历 （bottom, left) 到 (top + 1, left)
	遍历完一层，需要将 top, left 都 + 1 bottom, right 都 - 1

	先判断，如果行或者列为 0 则直接返回空数组
	left, right, top, bottom 分别初始化
	循环判断条件是 left <= right && top <= bottom
	按照上面的步骤遍历

*/
func solution(nums [][]int) []int {
	// 当前的最大行数和列数
	row := len(nums)
	col := len(nums[0])
	if row == 0 || col == 0 {
		return []int{}
	}
	res := make([]int, 0)
	left, right, top, bottom := 0, col-1, 0, row-1
	for left <= right && top <= bottom {
		// 从左到右遍历
		for i := left; i <= right; i++ {
			res = append(res, nums[top][i])
		}
		// 从上到下遍历
		for i := top; i <= bottom; i++ {
			res = append(res, nums[i][right])
		}
		// 注意这里的边界情况需要处理好
		if left < right && top < bottom {
			// 从右到左遍历
			for i := right - 1; i > left; i-- {
				res = append(res, nums[bottom][i])
			}
			// 从下到上遍历
			for i := bottom - 1; i > top; i-- {
				res = append(res, nums[i][left])
			}
		}
		left++
		top++
		right--
		bottom--
		// 索引处理
	}
	return res
}
