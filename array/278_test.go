package array

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("278. First Bad Version", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		target := 8
		got := solution(nums, target)
		want := []int{3, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", nums, want)
		}
	})
}

/*
	bad version 会导致后面的 version 都是 bad
	作为一个项目经理，你有一个版本列表 [1,2,...,n] 需要找到第一个 bad version
	有一个 isBadVersion(version) 函数用来返回 version 是否 bad
	现在需要实现一个函数来找到第一个 bad version 同时要求调用上面的函数的次数尽量小
	变种二分查找，找到左边第一个 bad version
*/
func solution(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		// 需要向左边移动，类似 34 只是判断条件变成了判断是否 bad version
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
