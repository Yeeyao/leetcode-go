package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("717. 1-bit and 2-bit Characters", func(t *testing.T) {
		input := []int{1, 0, 0}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("717. 1-bit and 2-bit Characters2", func(t *testing.T) {
		input := []int{1, 1, 1, 0}
		want := false
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("717. 1-bit and 2-bit Characters3", func(t *testing.T) {
		input := []int{0}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

//func solution(input []int) bool {
//	i := 0
//	inputLen := len(input)
//	for i < inputLen-1 {
//		if input[i] == 1 {
//			i += 2
//		} else {
//			i++
//		}
//	}
//	if i == inputLen {
//		return false
//	} else {
//		return true
//	}
//}

func solution(input []int) bool {
	var result bool
	i := 0
	inputLen := len(input)
	for i < inputLen {
		if input[i] == 1 {
			result = false
			i += 2
		} else {
			result = true
			i++
		}
	}
	return result
}
