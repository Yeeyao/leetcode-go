package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("1013. Partition Array Into Three Parts With Equal Sum", func(t *testing.T) {
		input := []int{1, 1, 0, 1, 1, 1}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input []int) int {
	maxOneNum := 0
	oneNum := 0
	for _, v := range input {
		if v == 1 {
			oneNum++
			if oneNum > maxOneNum {
				maxOneNum = oneNum
			}
		} else {
			oneNum = 0
		}
	}
	return maxOneNum
}
