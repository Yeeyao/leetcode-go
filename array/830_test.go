package array

import (
	"testing"
)

func TestPro(t *testing.T) {

	t.Run("830. Positions of Large Groups", func(t *testing.T) {
		input := "abbxxxxzzy"
		want := [][]int{{3, 6}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("830. Positions of Large Groups", func(t *testing.T) {
		input := "abcdddeeeeaabbbcd"
		want := [][]int{{3, 5}, {6, 9}, {12, 14}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("118. Pascal's Triangle3", func(t *testing.T) {
		input := "aaa"
		want := [][]int{{0, 2}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	直接统计每个字母的第一次出现以及最后一次出现位置和出现次数
	注意审题，这里是统计连续的
*/
func solution(input string) [][]int {
	var retArr [][]int
	limit := 3
	inputLen := len(input)
	for i, endPos := 0, 0; endPos < inputLen; i = endPos {
		for endPos < inputLen && input[endPos] == input[i] {
			endPos++
		}
		// 最后的加多了一
		if endPos-i >= limit {
			retArr = append(retArr, []int{i, endPos - 1})
		}
	}
	return retArr
}

//func solution(input string) [][]int {
//	var retArr [][]int
//	limit := 3
//	inputLen := len(input)
//	i, endPos := 0, 0
//	for endPos < inputLen {
//		for endPos < inputLen && input[endPos] == input[i] {
//			endPos++
//		}
//		// 最后的加多了一
//		if endPos-i >= limit {
//			retArr = append(retArr, []int{i, endPos - 1})
//		}
//		i = endPos
//	}
//	return retArr
//}

func IntSliceEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
