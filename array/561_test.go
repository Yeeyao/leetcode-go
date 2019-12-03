package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("561. Array Partition I", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(nums []int) int {
	resultLen := len(nums)
	sort.Ints(nums)
	sum := 0
	for i := 0; i < resultLen; i += 2 {
		sum += nums[i]
	}
	return sum
}
