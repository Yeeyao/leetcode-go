package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("643. Maximum Average Subarray I", func(t *testing.T) {
		input := []int{1, 12, -5, -6, 50, 3}
		k := 4
		want := 12.75
		got := solution(input, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("643. Maximum Average Subarray I 2", func(t *testing.T) {
		input := []int{0, 1, 1, 3, 3}
		k := 4
		want := 2.0
		got := solution(input, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

// 这里是需要固定 k 个元素
func solution(input []int, k int) float64 {
	inputLen := len(input)
	var sum float64
	if inputLen <= k {
		for i := 0; i < inputLen; i++ {
			sum += float64(input[i])
		}
		return sum / float64(inputLen)
	}
	var tempSum float64
	for i := 0; i < k; i++ {
		tempSum += float64(input[i])
	}
	sum = tempSum
	for i := k; i < inputLen; i++ {
		tempSum = tempSum - float64(input[i-k]) + float64(input[i])
		if tempSum/float64(k) > sum/float64(k) {
			sum = tempSum
		}
	}
	return sum / float64(k)
}

