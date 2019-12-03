package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("122. Best Time to Buy and Sell Stock II", func(t *testing.T) {
		input := []int{7, 1, 5, 3, 6, 4}
		want := 7
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input []int) int {
	if len(input) == 0 {
		return 0
	}
	maxBenefit := 0
	buy := input[0]
	for _, v := range input {
		if v > buy {
			maxBenefit += v - buy
		}
		buy = v
	}
	return maxBenefit
}
