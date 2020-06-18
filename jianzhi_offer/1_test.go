package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1 二维数组查找", func(t *testing.T) {
		nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		value := 4
		want := true
		got := solution(nums, value)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	二维数组查找

在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

注意到这里的二维数组性质，直接将开始的元素 a[i][j] 来和所查找的元素 value 进行比较，
边界条件是 行数小于最大行数，列数大于等于 0
这里的思路是，先将行数和列数初始为 0，最后一列
然后循环判断，如果当前元素小于目标值，直接将行数 + 1 遍历下一行，最终会找到一行的最后一个元素大于等于目标值。或者找不到表示元素太大了
直接针对目标行，如果当前元素太大了，则将列数 - 1
*/
func solution(nums [][]int, value int) bool {
	// 行数
	row := len(nums)
	// 空的情况
	if row == 0 {
		return false
	}
	// 列数
	col := len(nums[0])
	// 注意这里的初始值 第一行和最后一列
	i, j := 0, col-1
	for i < row && j >= 0 {
		if nums[i][j] == value {
			return true
			// 太大则需要向前列遍历
		} else if nums[i][j] > value {
			j--
			// 太小则直接下一行
		} else {
			i++
		}
	}
	return false
}

/*
	变种，查找特定的数 LeetCode
*/
