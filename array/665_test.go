package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("665. Non-decreasing Array", func(t *testing.T) {
		input := []int{4, 2, 3}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("665. Non-decreasing Array2", func(t *testing.T) {
		input := []int{4, 2, 1}
		want := false
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("665. Non-decreasing Array3", func(t *testing.T) {
		input := []int{4}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("665. Non-decreasing Array4", func(t *testing.T) {
		input := []int{4}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	非下降 因为是需要相邻的比较。那之前的较大的，无法传递到后面
	所以，每次遇到较大的，都需要将其向后赋值以传递下去
*/

func solution(input []int) bool {
	count := 0
	inputLen := len(input)
	/*
		因为只要改了，那后面的可能就不能计数了
		遇到非递增的 这里要么改前面的等于后面的，要么改后面的等于前面的；即变大或者变小
			1. 使得上一个数值等于或者小于当前的数值
			2. 使得当前的数值等于上一个数值
		当 i - 2 的位置元素等于或者小于当前的元素时，执行 1 因为这样不会使得 i - 2, i - 1 降序
		应该把 1 放在最优先级的情况。因为将前面的数值下降不会对结果有影响
		当 1 无法满足时，我们要使用 2。
		当将元素递增的时候，需要检查是否对后面的判断有影响
	*/
	for i := 1; i < inputLen && count < 2; i++ {
		if input[i-1] > input[i] {
			count++
			// 前面两个 或者是当前的上一个元素可以替换成当前元素
			if i < 2 || input[i-2] <= input[i] {
				// 1
				input[i-1] = input[i]
			} else {
				// 2 递增情况
				input[i] = input[i-1]
			}
		}
	}
	return count < 2
}

/*
	这种思路是找最大非降序子数组
	这里 count 的意义是非降序子数组的数量
	当 count > 2 则表示不满足了
	当 count == 2 则需要分析
		假如不满足非降序的位置索引是 1 或者 最后一个，则可以通过变换来处理，因此也是满足的
		还需要检查改变了位置的相关数值后也满足的情况
*/
func solution(input []int) bool {
	count := 1
	index := 0
	inputLen := len(input)
	if inputLen < 2 {
		return true
	}
	for i := 1; i < inputLen; i++ {
		if input[i] < input[i-1] {
			index = i
			count++
		}
	}
	if count == 1 {
		return true
	}
	if count == 2 {
		if index == 1 || index == inputLen-1 {
			return true
		}
		// i + 1 >= i - 1 i 的元素可以变成 i + 1 的元素
		// 或者 i - 2 <= i i - 1 对应元素可以变成 i 对应元素
		if input[index+1] >= input[index-1] || input[index] >= input[index-2] {
			return true
		}
		return false
	}
	return false
}
