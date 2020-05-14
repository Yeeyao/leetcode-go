package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run(" 287. Find the Duplicate Number ", func(t *testing.T) {
		input := []int{1, 3, 4, 2, 2}
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	注意题目说只有一个重复的元素，然后可能重复多次
	这里只需要找到重复的数字
	用新数组来计数吗？题目要求 O(1) 空间
	low = 1 high = len - 1
	同时，注意这里的数组元素大小是 1 到 n
	讨论中争论的是，统计数组元素时，如果左边的超过一半大于中间元素，那左边的就一定存在重复的元素
	然后就缩小范围到元素较多的区间继续找
	但是，这不能说明右边就不存在重复的元素了。
	类似鸽笼原理以及二分搜索
*/
func solution(nums []int) int {
	numsLen := len(nums)
	low, high := 0, numsLen-1
	for low < high {
		count := 0
		mid := low + (high-low)/2
		for _, v := range nums {
			// 这里和下面的以及返回值需要一致
			if v <= mid {
				count++
			}
		}
		if count <= mid {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

/*
	这种方法类似 Linked List Cycle II
	对于 nums[i] 不断进行 nums[nums[i]] 访问，如果存在重复元素则会一直循环访问
	第一步。一个指针每次进行一步的访问，另一个指针进行两步的访问，如果数组存在重复元素，
	则两个指针将会在一个点上相遇。
	第二步，其中一个指针在相遇的点，另一个从 0 位置移动。两个指针都每次移动一步，
	最后相遇的点就是循环的开始点，这个就是重复的元素
*/
func solution2(nums []int) int {
	numsLen := len(nums)
	if numsLen > 1 {
		// 找到相遇的点
		slow := nums[0]
		fast := nums[nums[0]]
		for slow != fast {
			slow = nums[slow]
			fast = nums[nums[fast]]
		}
		// 一起移动找到循环开始的点
		fast = 0
		for slow != fast {
			slow = nums[slow]
			fast = nums[fast]
		}
		return fast
	}
	return -1
}
