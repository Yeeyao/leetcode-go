package array

import (
	"math/rand"
	"testing"
	"time"
)

func TestPro(t *testing.T) {
	t.Run("215. Kth Largest Element in an Array", func(t *testing.T) {
		input := []int{1, 3, 4, 2, 2}
		k := 2
		want := 3
		got := findKthLargest(input, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("215. Kth Largest Element in an Array2", func(t *testing.T) {
		input := []int{2, 2, 2, 2, 2}
		k := 2
		want := 2
		got := findKthLargest(input, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("215. Kth Largest Element in an Array3", func(t *testing.T) {
		input := []int{2}
		k := 1
		want := 2
		got := findKthLargest(input, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("215. Kth Largest Element in an Array4", func(t *testing.T) {
		input := []int{2, 1}
		k := 2
		want := 1
		got := findKthLargest(input, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
Find the kth largest element in an unsorted array.
Note that it is the kth largest element in the sorted order, not the kth distinct element.
在一个未排序数组中找到第 k 大的元素 可以用 quick select 将前 k 个分到一个子数组，然后找子数组的最小值
*/
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, targetIndex int) int {
	q := randomPartition(nums, left, right)
	if q == targetIndex {
		return nums[q]
	}
	if q < targetIndex {
		return quickSelect(nums, q+1, right, targetIndex)
	}
	return quickSelect(nums, left, q-1, targetIndex)
}

func randomPartition(nums []int, left, right int) int {
	i := rand.Int()%(right-left+1) + left
	nums[i], nums[right] = nums[right], nums[i]
	return partition(nums, left, right)
}

func partition(nums []int, left, right int) int {
	pivotVal := nums[right]
	pivotCount := left
	for i := left; i < right; i++ {
		if nums[i] < pivotVal {
			nums[i], nums[pivotCount] = nums[pivotCount], nums[i]
			pivotCount++
		}
	}
	nums[pivotCount], nums[right] = nums[right], nums[pivotCount]
	return pivotCount
}

//func partition2(nums []int, left, right int) int {
//	pivotVal := nums[right]
//	i := left - 1
//	for j := left; j < right; j++ {
//		if nums[j] <= pivotVal {
//			i++
//			nums[i], nums[j] = nums[i], nums[j]
//		}
//	}
//	nums[i+1], nums[right] = nums[right], nums[i+1]
//	return i + 1
//}

/*
	可以用 heap
*/
