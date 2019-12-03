package array

import (
	"testing"
)

func TestPro(t *testing.T) {

	t.Run("766. Toeplitz Matrix", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		queries := [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}
		want := []int{8, 6, 2, 4}
		got := solution(array, queries)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	先全部相加，之后判断
	如果该元素加之前
		是偶数
			本次相加前需要先减去该数
		是奇数
			直接相加
	本次相加
		新的数字是偶数，直接相加
		新的数字数奇数，不相加

*/
func solution(array []int, queries [][]int, ) []int {
	arrayLen := len(array)
	retArr := make([]int, arrayLen)
	arrayEleSum := 0
	for j := 0; j < arrayLen; j++ {
		if array[j]%2 == 0 {
			arrayEleSum += array[j]
		}
	}
	for i := 0; i < len(queries); i++ {
		arrayAdd := queries[i][0]
		arrayEle := queries[i][1]
		beforeAdd := array[arrayEle]
		array[arrayEle] += arrayAdd
		afterAdd := array[arrayEle]
		if beforeAdd%2 == 0 {
			arrayEleSum -= beforeAdd
		}
		if afterAdd%2 == 0 {
			arrayEleSum += afterAdd
		}
		retArr[i] = arrayEleSum
	}
	return retArr
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
