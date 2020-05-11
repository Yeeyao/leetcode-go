package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("40. Combination Sum II", func(t *testing.T) {
		input := []int{10, 1, 2, 7, 6, 1, 5}
		target := 8
		want := [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}
		got := solution(input, target)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("40. Combination Sum II2", func(t *testing.T) {
		input := []int{2, 5, 2, 1, 2}
		target := 5
		want := [][]int{{1, 2, 2}, {5}}
		got := solution(input, target)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定的数组来求总和得到给定的和
	类似 39 但是每个元素只能使用一次
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
		// 这里需要加上去重的判断
		// 重复的只能统计第一个出现的
		if i == start || candidates[i] != candidates[i-1] {
			solArr = append(solArr, candidates[i])
			// 仅仅 + 1 就会把重复的统计进去
			solution2(candidates, cLen, target-candidates[i], solArr, retArr, i+1)
			// 把最近加入的数丢掉构造剩余的不同的组合
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
