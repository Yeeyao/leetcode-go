package array

import (
	"testing"
)

/*
300. Longest Increasing Subsequence
Given an unsorted array of integers, find the length of longest increasing subsequence.
There may be more than one LIS combination, it is only necessary for you to return the length.
Your algorithm should run in O(n^2) complexity.
[ref](https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/zui-chang-shang-sheng-zi-xu-lie-by-leetcode-soluti/)
求最长递增子序列，这里不要求连续
brute force 每个作为开头，然后向后
*/

func TestPro(t *testing.T) {
	t.Run("300. Longest Increasing Subsequence", func(t *testing.T) {
		input := []int{10, 9, 2, 5, 3, 7, 101, 18}
		want := 4
		got := lengthOfLIS(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	动态规划
	dp[i] 为考虑前 i 个元素，包含 i 结尾的最长上升子序列长度
	从小到大计算 dp[] 的数值
	dp[i] = max(do[j]) + 1 0 <= j < i nums[i] > nums[j]

	输入长度如果是 0 直接返回 0 dp[i] 表示以 i 为结尾的最长连续子数组的序列长度
	创建 dp 数组，初始化 dp[0] = 1 res = 1
	从 1 开始遍历每个输入元素，对每个当前元素，向前判断每个当前元素大于前面元素的 dp 值取最大的 dp 值
	当前元素的 dp 值是找到的最大 dp 值 + 1，然后更新全局最大 dp 值
*/
func lengthOfLIS(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	dp := make([]int, numsLen)
	dp[0] = 1
	res := 1
	for i := 1; i < numsLen; i++ {
		maxDp := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j] > maxDp {
				maxDp = dp[j]
			}
		}
		dp[i] = maxDp + 1
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

/*
	贪心 + 二分查找
	考虑一个简单的贪心，若需要使上升子序列尽可能长，必须让它增长尽可能慢，所以希望
	每次在上升子序列最后加上那个数尽可能小

	维护数组 d[i] 表示长度为 i 的最长上升子序列的末尾元素最小值，len 记录当前
	最长长度，起始值 len = 1, d[1] = num[0]
	依次遍历 nums[] 每个元素，更新 d[] 以及 len 的值
		如果 nums[i] > d[len] 则 len = len + 1，d[len + 1] = nums[i]
		否则在 d[1...len] 中找满足 d[i-1] < nums[j] < d[i] 的下标 i 更新 d[i] = nums[j]

	当前最长上升子序列长度为 len，顺序遍历 nums，对 nums[i]
		若 nums[i] > d[len] 加入 d 数组末尾，更新 len = len + 1
		否则在 d 中二分查找，找到第一个比 nums[i] 小的 d[k] 更新 d[k + 1] = nums[i]
*/

func lengthOfLIS2(nums []int) int {
	len, n := 1, len(nums)
	if n == 0 {
		return 0
	}
	d := make([]int, n+1)
	d[len] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > d[len] {
			len++
			d[len] = nums[i]
		} else {
			left, right, pos := 1, len, 0
			for left <= right {
				mid := left + (right-left)/2
				if nums[i] > d[mid] {
					pos = mid
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			d[pos+1] = nums[i]
		}
	}
	return len
}
