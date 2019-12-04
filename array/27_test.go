package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("27. Remove Element", func(t *testing.T) {
		input := []int{1, 3, 3, 1}
		val := 3
		want := 2
		got := solution(input, val)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("27. Remove Element2", func(t *testing.T) {
		input := []int{3, 2, 2, 3}
		val := 3
		want := 2
		got := solution(input, val)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

//func solution(input []int, val int) int {
//	inputLen := len(input)
//	count := 0
//	for i := 0; i < inputLen; i++ {
//		if input[i] == val {
//			count++
//		} else {
//			input[i-count] = input[i]
//		}
//	}
//	return inputLen - count
//}

// 只计算单独的元素
func solution(input []int, val int) int {
	inputLen := len(input)
	count := 0
	for i := 0; i < inputLen; i++ {
		if input[i] != val {
			input[count] = input[i]
			count++
		}
	}
	return count
}
