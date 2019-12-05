package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("35. Search Insert Position", func(t *testing.T) {
		input := []int{1, 3, 5, 6}
		target := 2
		want := 1
		got := solution(input, target)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("35. Search Insert Position2", func(t *testing.T) {
		input := []int{1, 3, 5, 6}
		target := 5
		want := 2
		got := solution(input, target)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("35. Search Insert Position3", func(t *testing.T) {
		input := []int{1, 3, 5, 6}
		target := 7
		want := 4
		got := solution(input, target)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("35. Search Insert Position4", func(t *testing.T) {
		input := []int{1, 3, 5, 6}
		target := 0
		want := 0
		got := solution(input, target)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}

/*
	二分查找
	相加再求中间，配合赋中值
	中间元素位置计算
*/
func solution(input []int, target int) int {
	inputLen := len(input)
	left, right := 0, inputLen-1
	if target < input[left] {
		return 0
	}
	if target > input[right] {
		return right + 1
	}
	for left < right {
		mid := (right + left) / 2
		if input[mid] == target {
			return mid
			// 这里递增需要和循环退出条件一致
		} else if target > input[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func solution(input []int, target int) int {
	i := 0
	for ; i < len(input); i++ {
		if input[i] >= target {
			return i
		}
	}
	return i
}
