package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("53. Maximum Subarray", func(t *testing.T) {
		input := []int{1, 3, 5, 4, 7}
		want := 20
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("53. Maximum Subarray2", func(t *testing.T) {
		input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		want := 6
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("53. Maximum Subarray2", func(t *testing.T) {
		input := []int{-2, -3, -1, 2, -1, 3, -2}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	如果找到了 O(n) 的解决方法，尝试使用分治算法处理
	只要加上下一个数负数后变成负数，则直接重新开始计算
	相加 置 0 比较 的顺序
*/
func solution(input []int) int {
	inputLen := len(input)
	// 因为不清楚最小值是多少，直接使用第一个元素来作为初始的和
	sum := input[0]
	tempSum := 0
	// 第一个元素被先使用了，因此从第二个元素开始遍历
	for i := 0; i < inputLen; i++ {
		// 因为每次负数都置 0 先加再比较
		tempSum += input[i]
		if tempSum > sum {
			sum = tempSum
		}
		// 一旦小于 0 就直接重置为 0
		if tempSum < 0 {
			tempSum = 0
		}
	}
	return sum
}

// 细微的初始化以及顺序交换
/*
	如果使用第一个元素初始化，那一开始就要比较，然后再检查负数
	因为一开始就比较，所以结尾就需要加上当前元素
 */
func solution(input []int) int {
	inputLen := len(input)
	// 因为不清楚最小值是多少，直接使用第一个元素来作为初始的和
	sum := input[0]
	tempSum := input[0]
	// 第一个元素被先使用了，因此从第二个元素开始遍历
	for i := 0; i < inputLen; i++ {
		// 因为每次负数都置 0 先加再比较
		if tempSum > sum {
			sum = tempSum
		}
		// 一旦小于 0 就直接重置为 0
		if tempSum < 0 {
			tempSum = 0
		}
		tempSum += input[i]
	}
	return sum
}
