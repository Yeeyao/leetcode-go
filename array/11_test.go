package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run(" 11. Container With Most Water ", func(t *testing.T) {
		input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		want := 49
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	类似 1014 ？
	给定数组表示每个高度，然后找到两个高度，使得它们组成的容器可以装最多的水
	min(a[i], a[j]) * (j - i) 最大 i < j
	从两边向中间遍历，因为两边的宽度是最大的，
	所以，中间的高度需要高于最低的才需要重新计算

初始化 left right 两个指针
循环条件是 left < right
	比较左右指针的大小取较小的作为高度
	得到高度乘以左右指针的差值得到面积然后更新最大面积
	两个循环过滤掉左右两边小于当前高度的元素
*/
func solution(height []int) int {
	left, right := 0, len(height)-1
	h, area := 0, 0
	for left < right {
		// 更新最小度
		if height[left] < height[right] {
			h = height[left]
		} else {
			h = height[right]
		}
		// 面积更新
		areaT := h * (right - left)
		if areaT > area {
			area = areaT
		}
		// 过滤掉中间的小于或者等于当前最小高度的元素
		for left < right && height[left] <= h {
			left++
		}
		for left < right && height[right] <= h {
			right--
		}
	}
	return area
}
