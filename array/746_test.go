package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("121. Best Time to Buy and Sell Stock", func(t *testing.T) {
		input := []int{10, 15, 20}
		want := 15
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("121. Best Time to Buy and Sell Stock2", func(t *testing.T) {
		input := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
		want := 6
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	贪心算法？
	优先选择第二个，当第二个比第一个大则选择第一个
*/

func solution(input []int) int {
	sumStep := 0
	inputLen := len(input)
	if inputLen == 1 {
		return 0
	}
	if inputLen == 2 {
		if input[0] > input[1] {
			return input[1]
		} else {
			return input[0]
		}
	}
	// 这里总结一下，只是比较三个元素中，中间元素和其他相邻两个元素之和的大小
	// 但这里有个问题是，类似 a b c d e f，取了 b 之后，下一组必须取 d 了
	for i := 0; i < inputLen-2; i += 2 {
		if input[i+2]+input[i] > input[i+1] {
			sumStep += input[i+1]
		} else {
			sumStep += input[i] + input[i+2]
		}
	}
	return sumStep
}
