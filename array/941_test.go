package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("941. Valid Mountain Array", func(t *testing.T) {
		input := []int{1, 2}
		want := false
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("941. Valid Mountain Array2", func(t *testing.T) {
		input := []int{3, 5, 5}
		want := false
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("941. Valid Mountain Array3", func(t *testing.T) {
		input := []int{0, 3, 2, 1}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	这里需要检查数组是否先递增，再递减
*/
func solution(input []int) bool {
	inputLen := len(input)
	if inputLen < 3 {
		return false
	}
	i := 0
	desc, asc := false, false
	for ; i < inputLen-1 && input[i] < input[i+1]; i++ {
		asc = true
	}
	for ; i < inputLen-1 && input[i] > input[i+1]; i++ {
		desc = true
	}
	if desc && asc && i == inputLen-1 {
		return true
	}
	return false
}

func validMountainArray(A []int) bool {
	if len(A) < 3 || A[0] >= A[1] {
		return false
	}
	// 前面两个比较了，开始是升序
	prev := A[1]
	goingUp := true
	for _, cur := range A[2:] {
		// 升序情况下，如果下一个元素是较小的，变成降序
		if goingUp {
			if cur < prev {
				goingUp = false
			}
			// 降序情况下 等于或者大于则直接返回
		} else {
			if cur == prev {
				return false
			}
			if cur > prev {
				return false
			}
		}
		// 更新上一个元素值
		prev = cur
	}
	// 最后需要是降序
	return !goingUp
}
