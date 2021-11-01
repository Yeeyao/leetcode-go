package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("491. Increasing Subsequences", func(t *testing.T) {
		input := []int{10, 9, 2, 5, 3, 7, 101, 18}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}

/*
给定一个整型数组 nums，返回所有的至少含有两个元素的不同的上升子序列
nums 可能含有重复的元素，两个相同的整型可以认为是特殊的上升子序列

类似排列的题目，可以直接遍历处理，
子函数
	从第一个开始到最后一个
	先将当前的加到临时结果中，递归处理，然后删除，继续遍历
*/
func solution(nums []int) [][]int {
	nLen := len(nums)
	res := make([][]int, 0)
	genSub(nums, 0, nLen, []int{}, &res)
	return res
}

func genSub(nums []int, begin, end int, temp []int, res *[][]int) {
	// 只要临时数组元素数量大于 1 就可以保存到结果数组
	if len(temp) > 1 {
		*res = append(*res, temp)
	}
	tempLen := len(temp)
	tempLast := 0
	if tempLen > 0 {
		tempLast = temp[tempLen-1]
	}
	for i := begin; i < end; i++ {
		// 临时数组为空或者当前元素大于临时数组的最后一个元素才能添加进去
		if tempLen == 0 || nums[i] >= tempLast {
			temp = append(temp, nums[i])
			genSub(nums, i+1, end, temp, res)
			temp = temp[tempLen:]
		}
	}
}
