package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("746. Min Cost Climbing Stairs", func(t *testing.T) {
		input := []int{10, 15, 20}
		want := 15
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("746. Min Cost Climbing Stairs2", func(t *testing.T) {
		input := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
		want := 6
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("746. Min Cost Climbing Stairs3", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 7, 6, 7, 4, 5, 6, 7}
		want := 36
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}

func solution(input []int) int {
	inputLen := len(input)
	for i := 2; i < inputLen; i++ {
		if input[i-1] > input[i-2] {
			input[i] += input[i-2]
		} else {
			input[i] += input[i-1]
		}
	}
	if input[inputLen-1] > input[inputLen-2] {
		return input[inputLen-2]
	} else {
		return input[inputLen-1]
	}
}

/*
	递归，超时的情况
*/

//func solution(input []int) int {
//	inputLen := len(input)
//	sumArr := make([]int, inputLen)
//	calcSum(input, inputLen-1, sumArr)
//	if sumArr[inputLen-1] > sumArr[inputLen-2 ] {
//		return sumArr[inputLen-2]
//	} else {
//		return sumArr[inputLen-1]
//	}
//}
//
//// 将小的加上去
//func calcSum(input []int, i int, sumArr []int) int {
//	// 开始返回情况
//	if i < 2 {
//		sumArr[i] = input[i]
//		return sumArr[i]
//	}
//	if calcSum(input, i-2, sumArr) > calcSum(input, i-1, sumArr) {
//		sumArr[i] = input[i] + calcSum(input, i-1, sumArr)
//	} else {
//		sumArr[i] = input[i] + calcSum(input, i-2, sumArr)
//	}
//	return sumArr[i]
//}
