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

/*
	类似上面的，一开始是使用边界两个。
	中间的高度需要比两边的高。
	两边中，较低的可以不用管
	这里只需要更新较低的，因为之前已经利用这个最低的计算过了，
	后面再进行移动的话，一定是需要较高的才能计算出更大的面积。
	然后得到更高的情况下，之前较低的已经是不需要了计算了
	因为向中间移动，宽度是不断减小的
*/
func solution2(height []int) int {
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
		// 坐标更新这里不同
		// 这里，只需要更新较小的高度的索引，因为更高的坐标对于 h 是有影响的
		if heigt[left] < height[right] {
			left++
		} else {
			right--
		
	}
	return area
}
