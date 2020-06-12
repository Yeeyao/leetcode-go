package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("209. Minimum Size Subarray Sum", func(t *testing.T) {
		nums := []int{8, 2, 4, 7}
		limit := 4
		want := 2
		got := solution(nums, limit)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定一个含有 n 个正数的数组以及一个整数 s，要求找到最短的连续子数组长度使得
	子数组和大于等于 s，如果不存在就返回 0
	这里直接类似 713 如果子數組和大于 s 则将前面的都减去直到小于 s
	双指针的滑动窗口
*/
func solution(s int, nums []int) int {
	i, j, temp := 0, 0, 0
	const intMax = int(^uint(0) >> 1)
	res := intMax
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	for i < numsLen {
		temp += nums[i]
		i++
		for temp >= s {
			// 注意这里先判断再减去
			if i-j < res {
				res = i - j
			}
			temp -= nums[j]
			j++
		}
	}
	if res == intMax {
		return 0
	} else {
		return res
	}
}
