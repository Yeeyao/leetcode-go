package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("713. Subarray Product Less Than K", func(t *testing.T) {
		input := 8
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	使用滑动窗口
	先判断新遍历的元素乘积是否大于 k 
	如果是，就需要从前面开始的位置将前面遍历过的元素
	从乘积里面除去，直到乘积小于 k
*/
func solution(nums []int, k int) int {
	if k == 0 {
		return 0
	}
	// 这里初始值
	res, tempProduct := 0, 1
	numsLen := len(nums)
	for i, j := 0, 0; j < numsLen; j++ {
		tempProduct *= nums[j]
		// 注意这里的条件 当前的积大于等于 k 就将最前面的除去
		for i <= j && tempProduct >= k {
			tempProduct /= nums[i]
			i++
		}
		res += j - i + 1
	}
	return res
}

/*
	需要找到连续的子数组，使得子数组的元素乘积小于 k，返回子数组的个数
	类似 907 ?
	一直向后遍历，当前开始元素计 1 个，加入 1 个元素后，如果还是满足，则满足的子数组增加的数量，
	就是以加入的元素为结尾的子数组数量，那就是前面元素数量 + 1
*/
// func solution2(nums []int, k int) int {
// 	numsLen := len(nums)
// 	tempProduct, tempNum := 1, 0
// 	res := 0
// 	for i := 0; i < numsLen; i++ {
// 		// 当前遍历的元素满足条件
// 		tempProductN := tempProduct * nums[i]
// 		if  tempProductN < k {
// 			tempNum++
// 			tempProduct = tempProductN
// 			res += tempNum

// 		} else if nums[i] < k {
// 			// 需要向前遍历
// 			tempProduct, tempNum = 1, 1
// 			res += tempNum
// 		} else {
// 			tempProduct, tempNum = 1, 0
// 		}
// 	}
// 	return res
// }