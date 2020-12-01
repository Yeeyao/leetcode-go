package array

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	//t.Run(" 90. Subsets2 ", func(t *testing.T) {
	//	input := []int{1, 2, 2}
	//	want := [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}
	//	got := solution2(input)
	//	if !reflect.DeepEqual(got, want) {
	//		t.Errorf("got: %v, want: %v", got, want)
	//	}
	//})

	t.Run(" 90. Subsets22 ", func(t *testing.T) {
		input := []int{2, 1, 2, 1, 3}
		want := [][]int{{}, {1}, {1, 1}, {1, 1, 2}, {1, 1, 2, 2}, {1, 1, 2, 2, 3}, {1, 1, 2, 3}, {1, 1, 3}, {1, 2}, {1, 2, 2}, {1, 2, 2, 3}, {1, 2, 3}, {1, 3}, {2}, {2, 2}, {2, 2, 3}, {2, 3}, {3}}
		got := solution(input)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run(" 90. Subsets23 ", func(t *testing.T) {
		input := []int{2, 1, 2, 1, 3}
		want := [][]int{{}, {1}, {1, 1}, {1, 1, 2}, {1, 1, 2, 2}, {1, 1, 2, 2, 3}, {1, 1, 2, 3}, {1, 1, 3}, {1, 2}, {1, 2, 2}, {1, 2, 2, 3}, {1, 2, 3}, {1, 3}, {2}, {2, 2}, {2, 2, 3}, {2, 3}, {3}}
		got := solution2(input)
		if !reflect.DeepEqual(got, want) {
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
	solutionHelper(nums, &retArr, []int{}, 0, numsLen)
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

/*
	slice 的错误问题
*/
func solution2(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	solution2Helper(&res, nums, []int{}, 0, len(nums))
	return res
}

func solution2Helper(res *[][]int, nums, temp []int, begin, end int) {
	temp2 := append([]int{}, temp...)
	fmt.Println(&temp2)
	*res = append(*res, temp2)
	//*res = append(*res, append([]int{}, temp...))
	tempOld := temp
	for i := begin; i < end; i++ {
		if i == begin || nums[i] != nums[i-1] {
			temp = append(temp, nums[i])
			solution2Helper(res, nums, temp, i+1, end)
			temp = tempOld
		}
	}
}
