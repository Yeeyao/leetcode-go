package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("37  在排序数组中查找数字出现的次数", func(t *testing.T) {
		nums := []int{1, 2, 3, 3, 3, 3, 4, 5}
		target := 3
		get := solution(nums, target)
		want := 4
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("37  在排序数组中查找数字出现的次数", func(t *testing.T) {
		nums := []int{1, 2, 3, 3, 3, 3, 4, 5}
		target := 7
		get := solution(nums, target)
		want := 0
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	类似二分查找来处理 找到后前后查找统计数量
	这里使用两次二分法分别找到左右边界
	0, 7, 3  4, 7, 5  6, 7, 7 超过最大值
	0, 7, 3  0, 3, 1  0, 1, 0  0, 0, 0 小于最小值
	nums := []int{1, 2, 3, 3, 3, 3, 4, 5}
*/
func solution(nums []int, target int) int {
	i, j := 0, len(nums)-1
	// 先找到有边界
	for i < j {
		mid := i + (j-i)/2
		// 注意这里仍旧向右移动
		if nums[mid] <= target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	// 如果找不到有边界表示不存在该元素，直接返回
	if j >= 1 && nums[j-1] != target {
		return 0
	}
	right := i
	// 找到左边界
	i, j = 0, len(nums)-1
	for i < j {
		mid := i + (j-i)/2
		if nums[mid] < target {
			i = mid + 1
			// 注意这里仍旧向右移动
		} else {
			j = mid - 1
		}
	}
	left := i
	return right - left
}
