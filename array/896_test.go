package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("1013. Partition Array Into Three Parts With Equal Sum", func(t *testing.T) {
		input := []int{1, 2, 2, 3}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1013. Partition Array Into Three Parts With Equal Sum2", func(t *testing.T) {
		input := []int{6, 5, 4, -4}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1013. Partition Array Into Three Parts With Equal Sum3", func(t *testing.T) {
		input := []int{1, 3, 2}
		want := false
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1013. Partition Array Into Three Parts With Equal Sum4", func(t *testing.T) {
		input := []int{1, 2, 4, 5}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input []int) bool {
	for i := 0; i < len(input)-1; i++ {
		// 前面相等的部分直接跳过
		if input[i] == input[i+1] {
			continue
		}
		if input[i] <= input[i+1] {
			for j := i; j < len(input)-1; j++ {
				if input[j] > input[j+1] {
					return false
				}
			}
		} else {
			for j := i; j < len(input)-1; j++ {
				if input[j] < input[j+1] {
					return false
				}
			}
		}
	}
	return true
}

// 半斤八两的方法
//func solution(input []int) bool {
//	asc := true
//	dec := true
//	for i := 1; i < len(input); i++ {
//		if input[i] < input[i-1] {
//			asc = false
//		}
//		if input[i] > input[i-1] {
//			dec = false
//		}
//	}
//
//	return asc || dec
//}
