package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("977. Squares of a Sorted Array", func(t *testing.T) {
		input := []int{-4, -1, 0, 3, 10, 20}
		want := 6
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("977. Squares of a Sorted Array2", func(t *testing.T) {
		input := []int{0, 1, 2, 3}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("977. Squares of a Sorted Array3", func(t *testing.T) {
		input := []int{1, 2, 3, 3}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("977. Squares of a Sorted Array4", func(t *testing.T) {
		input := []int{-3, -2, -2, -2, -1, -1, 0, 0}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("977. Squares of a Sorted Array5", func(t *testing.T) {
		input := []int{-3, -2, -1, -1, 0, 0, 1, 1, 2, 3, 4, 5}
		want := 6
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("977. Squares of a Sorted Array6", func(t *testing.T) {
		input := []int{0, 0}
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("977. Squares of a Sorted Array7", func(t *testing.T) {
		input := []int{3, 3, 2, 1, 0, -1, -2, -3, -3}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*

- 循环终止条件是两个指针相交

- 双指针，从两边向中间遍历，判断两个指针的绝对值大小

    - 如果相等，将其中一个进行平方计算然后计数器 + 1，两个指针都向中间移动一次

    - 如果左边大于右边 将右边的进行平方计算然后计数器 + 1，右边指针向中间移动一次

    - 如果左边小于右边 将左边的进行平方计算然后计数器 + 1，左边指针向中间移动一次

- 没有说明元素都是唯一的，所以需要过滤重复元素

*/
func solution(nums []int) int {
	numsLen := len(nums)
	res := 0
	i, j := 0, numsLen-1
	for i <= j {
		left, right := nums[i]*nums[i], nums[j]*nums[j]
		if left < right {
			res++
			for begin := j; j >= 0 && nums[begin] == nums[j]; j-- {

			}
		}
		if left > right {
			res++
			for begin := i; i < numsLen && nums[begin] == nums[i]; i++ {

			}
		}
		if left == right {
			res++
			for begin := j; j >= 0 && nums[begin] == nums[j]; j-- {

			}
			for begin := i; i < numsLen && nums[begin] == nums[i]; i++ {

			}
		}
	}
	return res
}
