package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("121. Best Time to Buy and Sell Stock", func(t *testing.T) {
		input := []int{7, 1, 5, 3, 6, 4}
		want := 5
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("121. Best Time to Buy and Sell Stock2", func(t *testing.T) {
		input := []int{7, 6, 4, 3, 1}
		want := 0
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	可以怎么知道当前是最低价位？
*/
func solution(input []int) int {
	if len(input) == 0 {
		return 0
	}
	maxBenefit := 0
	minPrice := input[0]
	for _, v := range input {
		benefit := v - minPrice
		if benefit > maxBenefit {
			maxBenefit = benefit
		}
		if v < minPrice {
			minPrice = v
		}
	}
	return maxBenefit
}
