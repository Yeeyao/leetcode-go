package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("39. Combination Sum", func(t *testing.T) {
		input := []int{2, 3, 6, 7}
		target := 7
		want := [][]int{{2, 2, 3}, {7}}
		got := solution(input, target)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("39. Combination Sum", func(t *testing.T) {
		input := []int{2, 3}
		target := 5
		want := [][]int{{2, 3}}
		got := solution(input, target)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定的数组来求总和得到给定的和
	其中，结果的数组不能有相同的数组。给定的数组的元素可以重复使用
	同样的回溯算法
	这里可以全部都进行组合，然后对结果的数组进行去重，但是效率很低
	假设 candidates [1,2,3,4] 拿就会组合 1 1,2,3,4  2 1,2,3,4 依此类推

	先进行排序，初始化最后结果的数据
*/
func solution(candidates []int, target int) [][]int {
	retArr := make([][]int, 0)
	sort.Ints(candidates)
	cLen := len(candidates)
	solution2(candidates, cLen, target, []int{}, &retArr, 0)
	return retArr
}

// 后面是中间结果数组，最终结果数组，起始数组下标
func solution2(candidates []int, cLen, target int, solArr []int, retArr *[][]int, start int) {
	// 满足总和
	if target == 0 {
		*retArr = append(*retArr, append([]int{}, solArr...))
		return
	}

	saLen := len(solArr)
	// 没有超过总和且不超过数组长度访问
	for i := start; i < cLen && target-candidates[i] >= 0; i++ {
		solArr = append(solArr, candidates[i])
		solution2(candidates, cLen, target-candidates[i], solArr, retArr, i)
		// 把最近加入的数丢掉构造剩余的不同的组合
		solArr = solArr[:saLen]
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
