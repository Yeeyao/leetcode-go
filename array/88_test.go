package array

import (
	"fmt"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("88. Merge Sorted Array", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		m := 3
		nums2 := []int{2, 5, 6}
		n := 3
		want := []int{1, 2, 2, 3, 5, 6}
		solution(nums1, m, nums2, n)
		if !IntSliceEqual(nums1, want) {
			t.Errorf("got: %v, want: %v", nums1, want)
		}
	})

	t.Run("88. Merge Sorted Array2", func(t *testing.T) {
		nums1 := []int{}
		m := 0
		nums2 := []int{1}
		n := 1
		want := []int{1}
		solution(nums1, m, nums2, n)
		if !IntSliceEqual(nums1, want) {
			t.Errorf("got: %v, want: %v", nums1, want)
		}
	})
}

func solution(nums1 []int, m int, nums2 []int, n int) {
	totalCount := m + n - 1
	i, j := m-1, n-1
	for ; i >= 0 && j >= 0; totalCount-- {
		if nums1[i] > nums2[j] {
			nums1[totalCount] = nums1[i]
			i--
		} else {
			nums1[totalCount] = nums2[j]
			j--
		}
	}
	for ; i >= 0; i-- {
		nums1[totalCount] = nums1[i]
		totalCount--
	}
	fmt.Println(i, j)
	for ; j >= 0; j-- {
		nums1[totalCount] = nums2[j]
		totalCount--
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
