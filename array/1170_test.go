package array

import (
	"fmt"
	"sort"
	"testing"
)

func TestPro(t *testing.T) {

	t.Run("1170. Compare Strings by Frequency of the Smallest Character", func(t *testing.T) {
		queries := []string{"cbd"}
		words := []string{"zaaaz"}
		want := []int{1}
		got := solution(queries, words)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1170. Compare Strings by Frequency of the Smallest Character2", func(t *testing.T) {
		queries := []string{"bbb", "cc"}
		words := []string{"a", "aa", "aaa", "aaaa"}
		want := []int{1, 2}
		got := solution(queries, words)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}

// 利用排序减少比较次数
func solution(queries, words []string) []int {
	var res, w []int
	for _, v := range words {
		w = append(w, count(v))
	}
	// 降序排列
	sort.Sort(sort.Reverse(sort.IntSlice(w)))
	fmt.Println(w)
	var q int
	for _, v := range queries {
		q = count(v)
		num := 0
		// 前面是降序排列
		for _, f2 := range w {
			if q < f2 {
				num++
			} else {
				break
			}
		}
		res = append(res, num)
	}
	return res
}

func count(s string) int {
	count, i := [26]int{}, 0
	for _, b := range s {
		count[b-'a']++
	}
	for count[i] == 0 {
		i++
	}
	return count[i]
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
