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
	二分的变种，如果找到位置就直接返回，找不到需要找可以插入的位置
	这里合适的插入位置怎么查找呢？如果找不到，最终 left right 会重合，这里的二分查找，
		最终 left 是否是大于的位置呢？不一定，left 的数值可能大于也可能小于 mid 的数值
	还可以提前和边界数值判断
*/
func solution(nums []int, target int) int {
	numsLen := len(nums)
	left, right := 0, numsLen-1
	// 比最小值小
	if target < nums[left] {
		return 0
	}
	// 比最大值大
	if target > nums[right] {
		return right + 1
	}
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// 只需要返回 left
	return left
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
