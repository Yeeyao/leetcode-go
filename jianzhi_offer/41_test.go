package jianzhi_offer

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("41  和为s的连续正数序列 ", func(t *testing.T) {
		target := 9
		get := solution(target)
		want := [][]int{{2, 3, 4}, {4, 5}}
		if !reflect.DeepEqual(get, want) {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	这里用简单的回溯加上剪枝
	需要注意元素数值连续
*/
func solution(target int) [][]int {
	var ret [][]int
	var temp []int
	// 要求元素数量至少有 2 个
	for i := 1; i < target-1; i++ {
		solutionHelper(&ret, temp, i, target, target)
	}
	return ret
}

func solutionHelper(ret *[][]int, temp []int, begin, sum, target int) {
	// 先将当前元素统计并判断
	sum -= begin
	temp = append(temp, begin)
	// 如果 sum 等于 0 直接保存并返回
	if sum == 0 {
		*ret = append(*ret, temp)
		return
	}
	// 剪枝
	if sum < 0 {
		return
	}
	// 递归下去下一个元素
	solutionHelper(ret, temp, begin+1, sum, target)
}
