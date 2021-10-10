package array

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("77. Combinations", func(t *testing.T) {
		n, k := 4, 2
		want := [][]int{{2, 2, 3}, {7}}
		got := solution(n, k)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定两个整数 n 和 k，返回在 1 到 n 返回的 k 个数字的所有组合
	每个都需要判断能否作为开始，开始则要求该元素以及后面的数量 >= k
	对每个开始的数字，下一个数字需要是后面的所有数字（这里需要判断后面的数字是否足够组成 k 个）
*/
func solution(n, k int) [][]int {
	retArr := make([][]int, 0)
	solution2(n, k, []int{}, &retArr, 1)
	return retArr
}

// 后面是中间结果数组，最终结果数组，起始数组下标
func solution2(n, k int, solArr []int, retArr *[][]int, start int) {
	// 满足数量
	if k == 0 {
		*retArr = append(*retArr, append([]int{}, solArr...))
		return
	}
	saLen := len(solArr)
	for i := start; i <= n && n-i+1 >= k-saLen; i++ {
		solArr = append(solArr, i)
		solution2(n, k-1, solArr, retArr, i+1)
		// 把最近加入的数丢掉构造剩余的不同的组合
		solArr = solArr[:saLen]
	}
}
