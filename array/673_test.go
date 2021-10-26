package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("646. Maximum Length of Pair Chain", func(t *testing.T) {
		input := []int{10, 9, 2, 5, 3, 7, 101, 18}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}

/*
	给定一个整型数组 nums，返回最长的递增子序列的数量，序列需要严格递增
	300 的变种。这里是需要统计数量。

	将 d 扩展为二维数组，d[i] 数组表示所有能成为长度为 i 的最长上升子序列的末尾元素的数值，将 d[i] = nums[j] 替换为将 nums[j] 追加到 d[i] 末尾
	d[i] 保留了历史信息，同时 d[i] 中的元素是有序的（单调非增）

	定义二维数组 cnt， cnt[i][j] 记录以 d[i][j] 为结尾的最长上升子序列的个数，计算 cnt[i][j] 可以考察 d[i-1] 和 cnt[i-1]
	所有满足 d[i-1][k] < d[i][j] 的 cnt[i-1][k] 累加到 cnt[i][j] 最终答案就是 cnt[maxLen] 的所有元素的和
*/
func solution(nums []int) int {
	d := [][]int{}
	cnt := [][]int{}
	for _, v := range nums {
		i := sort.Search(len(d), func(i int) bool { return d[i][len(d[i])-1] >= v })
		c := 1
		if i > 0 {
			k := sort.Search(len(d[i-1]), func(k int) bool { return d[i-1][k] < v })
			c = cnt[i-1][len(cnt[i-1])-1] - cnt[i-1][k]
		}
		if i == len(d) {
			d = append(d, []int{v})
			cnt = append(cnt, []int{0, c})
		} else {
			d[i] = append(d[i], v)
			cnt[i] = append(cnt[i], cnt[i][len(cnt[i])-1]+c)
		}
	}
	c := cnt[len(cnt)-1]
	return c[len(c)-1]
}

/*
	dp 类似 300
	dp[i] 表示以 nums[i] 结尾的最长上升子序列长度，cnt[i] 表示以 nums[i] 结尾的最长上升子序列个数。
	设 nums 最长上升子序列长度为 maxLen 则答案为满足 dp[i] = maxLen 的 i 对应的 cnt[i] 之和
	对 cnt[i] 等于所有满足 dp[j] + 1 = dp[i] 的 cnt[j] 之和，因此可以计算 dp[i] 的同时计算 cnt[i]

*/
func solution2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	maxLen := 0
	res := 0
	dp := make([]int, n)
	cnt := make([]int, n)
	for i, x := range nums {
		dp[i] = 1
		cnt[i] = 1
		for j, y := range nums[:i] {
			// 需要追加的时候判断
			if x > y {
				// 假如以 nums[j] 结尾的序列加上了 nums[i] 的序列长度大于 dp[i] 就需要更新 dp[i] cnt[i]
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					// 这里因为只是加上了 dp[i] 因此 cnt[i] = cnt[j]
					cnt[i] = cnt[j]
					// 等于就表示有相同的长度，需要将 cnt[j] 累加到 cnt[i] 上
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		// 这里更新 maxLen
		if dp[i] > maxLen {
			maxLen = dp[i]
			res = cnt[i]
		} else if dp[i] == maxLen {
			res += cnt[i]
		}
	}
	return res
}
