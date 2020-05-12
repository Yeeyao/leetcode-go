package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run(" 15. 3Sum ", func(t *testing.T) {
		input := []int{-1, 0, 1, 2, -1, -4}
		want := [][]int{{-1, 0, 1}, {-1, -1, 2}}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	有重复的元素
	类似二叉
	[0, len - 2]
	i == 0 or not same elements
	lo = i + 1 hi = len - 1 sum = 0 - num[i]
	lo < hi num[lo] + num[hi] == sum add  i, lo, hi
	skip same elements
	nums[lo] + num[hi] < sum lo++ else hi--

	最外层循环从 0 到 len - 2，因为超过最大值在右边就找不到两个元素了
	对于元素 i，从 i + 1 到 数组末尾找到两个元素，使得三者之和为 0，这里 sum = 0 - nums[i]
	就变成找两个元素使得它们的和是 sum
	需要处理重复元素，元素 i 的重复值需要过滤掉
	查找的循环中，先处理重复元素，然后再将左右移动，同时，需要一直遍历直到交叉
	一个优化是，第一个元素大于 0 时，直接返回，最小的元素都大于 0

*/
func solution(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	numsLen := len(nums)
	if numsLen == 0 {
		return [][]int{}
	}
	if nums[0] > 0 {
		return [][]int{}
	}
	retArr := make([][]int, 0)
	for i := 0; i < numsLen-2; i++ {
		// 过滤掉相同数值的元素
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			low, high := i+1, numsLen-1
			twoSum := 0 - nums[i]
			// 二叉循环
			for low < high {
				// 找到满足的数
				if nums[low]+nums[high] == twoSum {
					retArr = append(retArr, []int{nums[i], nums[low], nums[high]})
					// 过滤掉重复的元素
					for low < high && nums[low] == nums[low+1] {
						low++
					}
					for low < high && nums[high] == nums[high-1] {
						high--
					}
					low++
					high--
				} else if nums[low]+nums[high] > twoSum {
					high--
				} else {
					low++
				}
			}
		}
	}
	return retArr
}

func IntSliceEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
