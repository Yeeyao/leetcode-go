package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("747. Largest Number At Least Twice of Others", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		want := -1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("747. Largest Number At Least Twice of Others2", func(t *testing.T) {
		input := []int{3, 6, 1, 0}
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("747. Largest Number At Least Twice of Others3", func(t *testing.T) {
		input := []int{0, 0, 0, 1}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input []int) int {
	inputLen := len(input)
	var max1, max2 int
	if inputLen < 2 {
		return 0
	}
	if input[0] > input[1] {
		max1, max2 = 0, 1
	} else {
		max1, max2 = 1, 0
	}
	for i := 2; i < inputLen; i++ {
		if input[i] > input[max1] {
			max2 = max1
			max1 = i
		} else if input[i] > input[max2] {
			max2 = i
		}
	}
	if input[max1] >= 2*input[max2] {
		return max1
	} else {
		return -1
	}
}
