package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("1013. Partition Array Into Three Parts With Equal Sum", func(t *testing.T) {
		input := []int{0, 1, 0, 3, 12}
		input2 := []int{1, 3, 12, 0, 0}
		solution(input)
		if !IntSliceEqual(input, input2) {
			t.Errorf("got: %v, want: %v", input, input2)
		}
	})
}

/*
	统计 0 的数量，然后将非 0 的往前移动就行，最后补 0
*/
func solution(input []int) {
	inputLen := len(input)
	zeroNum := 0
	for i := 0; i < inputLen; i++ {
		if input[i] == 0 {
			zeroNum++
		} else {
			input[i-zeroNum] = input[i]
		}
	}
	for i := inputLen - zeroNum; i < inputLen; i++ {
		input[i] = 0
	}
}

func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
