package jianzhi_offer

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("32 把数组排成最小的数", func(t *testing.T) {
		intSlice := []int{3, 30, 34, 5, 9}
		get := solution(intSlice)
		want := "3033459"
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	直接对所有的元素进行排序然后按照顺序处理
	对于元素 x, y
		如果 x + y > y + x 则 x > y
		如果 x + y < y + x 则 x < y
	这里可以使用快速排序或者内置的函数处理
	TODO: 递归和迭代都要试试
*/
func solution(nums []int) string {
	qsort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	var s string
	for _, i := range nums {
		s += strconv.Itoa(i)
	}
	return s
}

func qsort(nums []int, l, r int) {
	if r > l {
		// 可以随机选取 pivot index
		pivotIndex := l
		pivotNewIndex := partition(nums, l, r, pivotIndex)
		qsort(nums, l, pivotNewIndex-1)
		qsort(nums, pivotNewIndex+1, r)
	}
}

func partition(nums []int, l, r, pivotIndex int) int {
	// 先将 pivot 交换到最右边
	pivot := nums[pivotIndex]
	nums[r], nums[pivotIndex] = nums[pivotIndex], nums[r]
	storeIndex := l
	// 移动
	for i := l; i < r; i++ {
		if strconv.Itoa(nums[i])+strconv.Itoa(pivot) < strconv.Itoa(pivot)+strconv.Itoa(nums[i]) {
			nums[i], nums[storeIndex] = nums[storeIndex], nums[i]
			storeIndex++
		}
	}
	nums[storeIndex], nums[r] = nums[r], nums[storeIndex]
	return storeIndex
}

func qsort2(nums []int, l, r int) {
	if r > l {
		// 可以随机选取 pivot index
		pivotIndex := l
		pivotNewIndex := partition2(nums, l, r, pivotIndex)
		qsort2(nums, l, pivotNewIndex-1)
		qsort2(nums, pivotNewIndex+1, r)
	}
}

func partition2(nums []int, l, r, pivotIndex int) int {
	// 先将 pivot 交换到最右边
	pivot := nums[pivotIndex]
	nums[r], nums[pivotIndex] = nums[pivotIndex], nums[r]
	storeIndex := l
	// 移动
	for i := l; i < r; i++ {
		if nums[i] < pivot {
			nums[i], nums[storeIndex] = nums[storeIndex], nums[i]
			storeIndex++
		}
	}
	nums[storeIndex], nums[r] = nums[r], nums[storeIndex]
	return storeIndex
}
