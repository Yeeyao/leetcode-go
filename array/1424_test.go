package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1424. Diagonal Traverse II", func(t *testing.T) {
		input := 8
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*

*/
func solution(nums [][]int) []int {
	var res []int
	var flags []int
	for i := 0; ; i++ {
		// i 小于 nums 长度，就存放序号
		if i < len(nums) {
			flags = append(flags, i)
		}
		// 从后向前遍历
		for j := len(flags) - 1; j >= 0; j-- {
			// v 是当前序号到遍历行的长度
			v := i - flags[j]
			// 如果长度大于等于遍历行号的列数
			if v >= len(nums[flags[j]]) {
				copy(flags[j:], flags[j+1:])
				flags = flags[:len(flags)-1]
			} else {
				res = append(res, nums[flags[j]][v])
			}
		}
		if len(flags) == 0 {
			break
		}
	}
	return res
}

/*
	利用观察到的，目标对角线的所有点，其行和列的坐标和是相等的
	所以可以遍历所有元素，然后将根据坐标和将元素存放 最后将存放的从小到大遍历
	内部遍历是反向保存
*/
func solution2(nums [][]int) []int {
	var res []int
	sumMap := make(map[int][]int)
	// 这里直接反序存放，后面就不需要反序遍历
	for i := len(nums) - 1; i >= 0; i-- {
		for j := range nums[i] {
			if _, ok := sumMap[i + j]; ok {
				sumMap[i + j] = append(sumMap[i + j], nums[i][j])
			} else {
				sumMap[i+j] = []int{nums[i][j]}
			}
		}
	}
	sumNum := len(sumMap)
	for i := 0; i < sumNum; i++{
		for _, v := range sumMap[i] {
			res = append(res, v)
		}
	}
	return res
}

// /*
// 	斜向上顺序遍历元素
// 	首先从第一行向最后一行遍历 
// 		当前行往上面的行遍历直到第一行，每次列号 + 1 需要判断结尾
// 	然后最后一行遍历所有列，同样向上遍历
// 	不完整版本 同时会 TLE
// */
// func solution(nums [][]int) []int {
// 	var res []int
// 	// 行数
// 	row := len(nums)
// 	if row == 1{
// 		return nums[0]
// 	}
// 	maxCol := 0
// 	// 遍历行 第一行到最后一行
// 	for i := 0; i < row; i++ {
// 		// 遍历的列数递增 每次行数增加，需要向上遍历，列数增加
// 		// 遍历的行数递减
// 		rCol := i
// 		for j := 0; j < i + 1; j++ {
// 			// 该行该位置有元素
// 			rowCol := len(nums[rCol])
// 			if rowCol > maxCol {
// 				maxCol = rowCol
// 			}
// 			if j < rowCol {
// 				res = append(res, nums[rCol][j])
// 			}
// 			rCol--
// 		}
// 	}
// 	// 只有一行
// 	if row == 1 {
// 		return res
// 	}
// 	// 遍历最长的列 这里最后一行的列数不一定是最大的，因此需要获得最大列数来处理
// 	// 注意这里从第一列开始遍历
// 	for i := 1; i < maxCol; i++ {
// 		// 遍历的行数递减
// 		rCol := maxCol - 1
// 		// 遍历的列数递增
// 		for j := i; j < maxCol; j++ {
// 			// 该行该位置有元素
// 			rowCol := len(nums[rCol])
// 			if j < rowCol {
// 				res = append(res, nums[rCol][j])
// 			}
// 			rCol--
// 		}
// 	}
// 	return res
// }