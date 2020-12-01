package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("724. Find Pivot Index", func(t *testing.T) {
		input := []int{1, 7, 3, 6, 5, 6}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("724. Find Pivot Index2", func(t *testing.T) {
		input := []int{1, 2, 3}
		want := -1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("724. Find Pivot Index3", func(t *testing.T) {
		input := []int{-1, -1, -1, -1, -1, 0}
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("724. Find Pivot Index4", func(t *testing.T) {
		input := []int{-1, -1, 0, -1, -1, 0}
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

// 只利用两个和
func solution(input []int) int {
	inputLen := len(input)
	if inputLen == 0 {
		return -1
	}
	leftSum, rightSum := 0, 0
	for _, v := range input {
		rightSum += v
	}
	for i, v := range input {
		rightSum -= v
		if leftSum == rightSum {
			return i
		}
		leftSum += v
	}
	return -1
}

// 这里是空间换取时间
func solution(input []int) int {
	inputLen := len(input)
	if inputLen == 0 {
		return -1
	}
	leftSum := 0
	rightSum := make([]int, inputLen)
	tempSum := 0
	for i := inputLen - 1; i >= 0; i-- {
		rightSum[i] = tempSum
		tempSum += input[i]
	}
	for i := 0; i < inputLen; i++ {
		if leftSum == rightSum[i] {
			return i
		}
		leftSum += input[i]
	}
	return -1
}
