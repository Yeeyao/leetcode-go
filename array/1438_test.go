package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1438. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit", func(t *testing.T) {
		nums := []int{8,2,4,7}
		limit := 4
		want := 2
		got := solution(nums, limit)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1438. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit2", func(t *testing.T) {
		nums := []int{10,1,2,4,7,2}
		limit := 5
		want := 4
		got := solution(nums, limit)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1438. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit3", func(t *testing.T) {
		nums := []int{4,2,2,2,4,4,2,2}
		limit := 0
		want := 3
		got := solution(nums, limit)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	[lee](https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/discuss/609771/JavaC%2B%2BPython-Deques-O(N))
	使用 dequeue 找子数组的最大最小值
 */
func solution(nums []int, limit int) int {

}

/*
	给定 nums 数组以及 limit，要求返回最长的子数组长度，子数组元素要求满足任意两个元素的差值不超过 limit
	不能直接修改原数组 可以直接用回溯法将所有的子数组都枚举出来然后判断，但是没有必要
	直接用双指针 每次遍历记录下最大最小值然后跟 limit 比较
	先更新最大最小值，然后比较 limit
		如果超过了则需要重新计算子数组
		没有超过，则需要更新最大长度
*/
//func solution(nums []int, limit int) int {
//	numsLen := len(nums)
//	count, res := 0, 0
//	tempMin, tempMax := nums[0], nums[0]
//	for i := 0; i < numsLen; i++ {
//		if nums[i] < tempMin {
//			tempMin = nums[i]
//		}
//		if nums[i] > tempMax {
//			tempMax = nums[i]
//		}
//		if tempMax-tempMin > limit {
//			// 这里超过限制，从前面的第一个元素开始丢弃，直到子数组满足条件
//			// 考虑到这种情况，当前的方法还是太麻烦
//			// 注意这里超过限制之后，count 初始化是 1，表示当前的元素
//			tempMin, tempMax, count = nums[i], nums[i], 1
//		} else {
//			count++
//			if count > res {
//				res = count
//			}
//		}
//	}
//	return res
//}
