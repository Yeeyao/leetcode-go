package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("414. Third Maximum Number", func(t *testing.T) {
		input := []int{2, 6, 4, 8, 10, 9, 15}
		want := 9
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("414. Third Maximum Number2", func(t *testing.T) {
		input := []int{3, 2, 1}
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("414. Third Maximum Number3", func(t *testing.T) {
		input := []int{2, 1}
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("414. Third Maximum Number4", func(t *testing.T) {
		input := []int{2, 2, 3, 1}
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input []int) int {
	const intMax = int(^uint(0) >> 1)
	const intMin = ^intMax
	max1, max2, max3 := intMin, intMin, intMin
	for _, v := range input {
		if v != max1 && v != max2 && v != max3 {
			if v > max1 {
				max3 = max2
				max2 = max1
				max1 = v
			} else if v > max2 {
				max3 = max2
				max2 = v
			} else if v > max3 {
				max3 = v
			}
		}
	}
	if max3 == intMin {
		return max1
	}
	return max3
}
