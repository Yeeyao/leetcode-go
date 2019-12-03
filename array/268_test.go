package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("leetcode 238  Product of Array Except Self", func(t *testing.T) {
		input := []int{1, 0, 3, 4}
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

//func solution(input []int) int {
//	inputLen := len(input)
//	totalSum := 0
//	inputSum := 0
//	for i := 0; i < inputLen+1; i++ {
//		totalSum += i
//	}
//	for _, v := range input {
//		inputSum += v
//	}
//	return totalSum - inputSum
//}

/*
	使用异或运算符 两个元素，如果都出现了，则对自己异或运算会得到0
	因此，最后的结果，将会得到缺失的那个元素，一个数和 0 异或将得到本身
*/
func solution(input []int) int {
	result := len(input)
	i := 0
	for _, v := range input {
		result ^= v
		result ^= i
		i++
	}
	return result
}
