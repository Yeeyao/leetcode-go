package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("769. Max Chunks To Make Sorted", func(t *testing.T) {
		input := []int{4, 3, 2, 1, 0}
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("769. Max Chunks To Make Sorted", func(t *testing.T) {
		input := []int{1, 0, 2, 3, 4}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	将给定的数组进行划分，然后对每个划分后的数组进行排序，
	之后连接它们，得到的数组结果等于已排序的数组
	求最大的划分数组数量
	能不能划分的依据是，划分后的部分之间需要是有序的

	需要观察划分数组可以满足题目条件所包含的特征
	遍历数组，对元素 A[i] 如果 max(A[0]...A[i]) = i，我们就可以在这个位置上将是数组划分
	因为 A 本来就是数组元素唯一的排列，对于当前的元素来说，到当前位置的子数组最大值
	等于当前的元素值，那它一定只包含 A[0]...A[i] 的这些数，所以就可以划分
*/
func solution(arr []int) int {
	// 第一个是累计的最大值
	curMax, num := -1, 0
	for i, v := range arr {
		if v > curMax {
			curMax = v
		}
		if curMax == i {
			num++
		}
	}
	return num
}
