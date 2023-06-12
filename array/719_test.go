package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("719. Find K-th Smallest Pair Distance", func(t *testing.T) {
		want := 7
		nums := []int{1, 2, 3}
		k := 1
		got := solution(nums, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	a, b 之间的距离是 a 到 b 的绝对距离
	给定 nums 以及 k，找到所有元素中的 nums[i] 到 nums[j] 的第 k 个最小的距离，其中 0 <= i < j < nums.length
	感觉还是 DP 了？毕竟还是需要全部计算的，同时 nums 不能变换元素位置，这里只是减少了重复计算距离
	同时，查找第 k 小的元素好像也有一种方法，但我忘了

	暴力：1.计算所有的距离，2.保存到数组然后数组排序，返回第 k 个
	其中 1 可以使用 DP 来将已经计算的结果保存 2 好像可以用单调栈？不对，

	[binary search](https://leetcode.com/problems/find-k-th-smallest-pair-distance/discuss/109075/Java-solution-Binary-Search)
	关键在于只需要关注距离大小
		先将数组排序，然后计算最大最小的距离，然后就统计数组中到某个距离的序对的数量，这里也就表示这个距离的大小（是第几个），
	这里需要找到第 k 个，到这个距离的序对数量也应该是 k

	力扣加加从堆的解法过渡到配合二分法。

这里如果直接使用堆处理，则还是需要 n^2 的比较。
这里使用多路归并的思路，则是记录下当前的差值，然后多路是，代码遍历是以所有的相邻为基础，然后记录下 from, to
- 1 获取一个当前的数值
- 2 找到当前数值的剩下的其他数值进行差值计算
- 也就是最外面的循环确定减数，里面则是当前减数确定的情况下，确定被减数
- 这里的问题是，如果 k 很大，则还是需要 n^2 次比较，因此提到使用二分法解决(需要单调性)

这里提到求第 k 小本质是求不大于其本身的有 k - 1 个的那个数，比如第 1 小就是求不大于其本身的有 0 个数。这个题目，最大的差值计作 max_diff
就可以问，数对差小于 max_diff 有几个，小于 max_diff - 1 的有几个。。。不断找直到找到小于 x 的有 k - 1 个。但是也有问题
1. 小于 x 的有 k - 1 个的数可能不止 1 个
2. 无法确定小于 x 就有 k - 1 个数存在，比如 [1,1,1,1,2] 求第 3 大，则小于 x 的两个数不存在
因此将思路调整为求小于等于 x 有 k 个的，使用二分法最左模板

*/
func solution(nums []int, k int) int {
	numsLen := len(nums)
	sort.Ints(nums)
	low := nums[1] - nums[0]
	for i := 1; i < numsLen-1; i++ {
		temp := nums[i+1] - nums[i]
		if temp < low {
			low = temp
		}
	}
	high := nums[numsLen-1] - nums[0]
	for low < high {
		mid := low + (high-low)/2
		if countPairs(nums, mid) < k {
			low = mid + 1
		} else {
			high = mid
		}
	}
	/*
		这里 mid 明明是 low 和 high 计算出来的，为什么 low 一定是序对的距离
		countPairs(nums, mid) == k 这里没有提前返回
		因为对于  countPairs(nums, mid) < k 来说，如果 mid + 1 这个距离不存在的话，则
		countPairs(nums, mid) 和 countPairs(nums, mid + 1) 这两个数值是相等的，导致 low 以及 high 不断变化
	*/

	return low
}

// 返回第一个比 key 大的元素的索引
func upperBound(nums []int, low, high, key int) int {
	if nums[high] <= key {
		return high + 1
	}
	for low < high {
		mid := low + (high-low)/2
		if key >= nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

// 返回距离小于等于 mid 的 pairs 的数量
func countPairs(nums []int, mid int) int {
	n, res := len(nums), 0
	for i := 0; i < n; i++ {
		res += upperBound(nums, i, n-1, nums[i]+mid) - i - 1
	}
	return res
}
