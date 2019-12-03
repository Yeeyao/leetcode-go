package array

import (
	"sort"
	"testing"
)

/*
	先排序，完了之后两两比较，计算两个元素的差值，当有更小的差值则需要
	保存新的差值元素值
*/
func TestPro(t *testing.T) {

	t.Run("1200. Minimum Absolute Difference", func(t *testing.T) {
		input := []int{4, 2, 1, 3}
		want := [][]int{{1, 2}, {2, 3}, {3, 4}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1200. Minimum Absolute Difference2", func(t *testing.T) {
		input := []int{3, 8, -10, 23, 19, -4, -14, 27}
		want := [][]int{{-14, -10}, {19, 23}, {23, 27}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1200. Minimum Absolute Difference3", func(t *testing.T) {
		input := []int{1, 3, 6, 10, 15}
		want := [][]int{{1, 3}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1200. Minimum Absolute Difference3", func(t *testing.T) {
		input := []int{40, 11, 26, 27, -20}
		want := [][]int{{26, 27}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(arr []int) [][]int {
	arrLen := len(arr)
	var retArr [][]int
	sort.Ints(arr)
	minDiff := arr[1] - arr[0]
	for i := 0; i < arrLen-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff < minDiff {
			minDiff = diff
			retArr = retArr[:0]
			retArr = append(retArr, []int{arr[i], arr[i+1]})
		} else if diff == minDiff {
			retArr = append(retArr, []int{arr[i], arr[i+1]})
		}
	}
	return retArr
}

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
