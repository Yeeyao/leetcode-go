package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run(" 90. Subsets2 ", func(t *testing.T) {
		input := []int{1, 2, 2}
		want := [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}
		got := solution(input)
		if IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	求排列
	1			2			3
2		3	1		3	1		2
3		2	3		1	2		1
*/
func solution(nums []int) [][]int {
	retArr := make([][]int, 0)
	// 注意这里的排序处理
	sort.Ints(nums)
	numsLen := len(nums)
	solutionHelper(nums, &retArr, []int{}, numsLen)
	return retArr
}

func solutionHelper(nums []int, retArr *[][]int, solArr []int, start, numsLen int) {
	*retArr = append(*retArr, append([]int{}, solArr...))
	saLen := len(solArr)
	// 循环递归处理
	for i := start; i < numsLen; i++ {
		// 去重处理 相邻元素相等，只能处理一个
		if i == start || nums[i] != nums[i-1] {
			solArr = append(solArr, nums[i])
			solutionHelper(nums, retArr, solArr, i+1, numsLen)
			solArr = solArr[:saLen]
		}
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
