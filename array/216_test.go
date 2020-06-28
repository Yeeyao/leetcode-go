package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("216. Combination Sum III", func(t *testing.T) {
		k := 3
		n := 7
		want := [][]int{{1, 2, 4}}
		got := solution(k, n)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("216. Combination Sum III2", func(t *testing.T) {
		k := 3
		n := 7
		want := [][]int{{1, 2, 4}}
		got := solution2(k, n)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定 k 和 n，找到所有的数的组合
	其中 k 是每组的数字数量，n 是改组的数字的总和
	其中每组以内的数字只能用一次

	需要使用回溯算法
	可以加入的条件是总和满足以及元素数量满足了
	初始化一个结果的 slice 然后保存结果的 slice
	使用一个保存当前元素的 slice solArr 从第一个元素开始，递归调用来计算所有的组合
	一开始是 1,2,3... 保存在 solArr 中，递归调用该函数，每次调用完，需要将 solArr 的最后一个元素删除
	因为循环中再次调用使用的是新的元素

*/
func solution(k, n int) [][]int {
	retArr := make([][]int, 0)
	solution2(&retArr, []int{}, k, 1, n)
	return retArr
}

func solution2(retArr *[][]int, solArr []int, k, start, n int) {
	// 找到解了，需要保存下来
	lenSa := len(solArr)
	// 超过指定数量了，直接返回
	if lenSa > k {
		return
	}
	// 满足数量和总和
	if lenSa == k && n == 0 {
		*retArr = append(*retArr, append([]int{}, solArr...))
		return
	}
	// 解的元素数量还不足
	// 循环将每个元素加入中间结果数组中
	for i := start; n-i >= 0 && i <= 9; i++ {
		solArr = append(solArr, i)
		solution2(retArr, solArr, k, i+1, n-i)
		// 把最近加入的数丢掉构造剩余的不同的组合
		solArr = solArr[:lenSa]
	}
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

//func solution2(nums []int) [][]int {
//	retArr := make([][]int, 0)
//	numsLen := len(nums)
//	solution2_helper(nums, retArr, []int{}, 0, numsLen)
//	return retArr
//}
//
//func solution2_helper(nums []int, retArr [][]int, solArr []int, start, numsLen int) {
//	// 什么时候保存结果 这里每次都直接保存
//	*retArr = append(*retArr, append([]int{}, solArr...))
//	saLen := len(solArr)
//	// 循环递归处理
//	for i := start; i < numsLen; i++ {
//		solArr = append(solArr, nums[i])
//		solution2_helper(nums, retArr, solArr, start+1, numsLen)
//		solArr = solArr[:saLen]
//	}
//}
