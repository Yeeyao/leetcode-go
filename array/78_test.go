package array

import (
	"testing"
	"sort"
)

func TestPro(t *testing.T) {
	t.Run(" 78. Subsets ", func(t *testing.T) {
		input := []int{1, 2, 3}
		want := [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}
		got := solution(input)
		if IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run(" 78. Subsets2 ", func(t *testing.T) {
		input := []int{1, 2, 3}
		want := [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}
		got := solution2(input)
		if IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	找当前集合的所有子集
	计数器的原理就行了，读取当前的子集的元素数量，然后直接转换为2进制
	然后根据 0 表示元素不存在，1 表示存在来获得所有的子集

	首先计算所有的组合数量即 2 * len 个
	遍历每个组合，针对每个组合，检验其所有的二进制位是否为 1
	是，则表示需要将输入数组的对应位保存到该组合的数组中
*/
func solution(nums []int) [][]int {
	numsLen := len(nums)
	totalNum := PowTwo(numsLen)
	if numsLen == 0 {
		return [][]int{}
	}
	retArr := make([][]int, totalNum)
	for i := 0; i < totalNum; i++ {
		count := 0
		for j := 0; j < numsLen; j++ {
			// 当前判断的二进制位
			be := PowTwo(j)
			// 与二进制数值做位与
			if i&be != 0 {
				retArr[i] = append(retArr[i], nums[count])
			}
			count++
		}
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

func PowTwo(num int) int {
	Res := 1
	for i := 0; i < num; i++ {
		Res *= 2
	}
	return Res
}

func solution2(nums []int) [][]int {
	retArr := make([][]int, 0)
	sort.Ints(nums)
	numsLen := len(nums)
	solution2_helper(nums, &retArr, []int{}, 0, numsLen)
	return retArr
}

func solution2_helper(nums []int, retArr *[][]int, solArr []int, start, numsLen int) {
	// 什么时候保存结果 这里每次都直接保存
	*retArr = append(*retArr, append([]int{}, solArr...))
	saLen := len(solArr)
	// 循环递归处理
	for i := start; i < numsLen; i++ {
		solArr = append(solArr, nums[i])
		solution2_helper(nums, retArr, solArr, i+1, numsLen)
		solArr = solArr[:saLen]
	}
}
