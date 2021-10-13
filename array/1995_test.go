package array

import (
	"testing"
)

/*
18. 4Sum
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums
such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.
*/
func TestPro(t *testing.T) {
	t.Run("1995. Count Special Quadruplets", func(t *testing.T) {
		nums := []int{35, 15, 38, 1, 10, 26}
		want := 1
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1995. Count Special Quadruplets2", func(t *testing.T) {
		nums := []int{1, 1, 1, 3, 5}
		want := 4
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定一个索引从 0 开始的数组 nums，返回数组中不同的四元组数量（a,b,c,d）满足 nums[a] + nums[b] + nums[c] == nums[d] a < b < c < d
	这里是索引组成的元组，因此元素数值是允许重复的
	还是参考 454 的，直接使用 hashmap 但是这里需要三个元素

TODO: 参考
https://leetcode.com/problems/count-special-quadruplets/discuss/1451080/JavaPython-O(n2)-solution-with-explanation

a + b + c = d ==> a + b = d - c
将数组拆分为两个部分，下标 [0,i-1] [i,n-1]
对每个 i 其实这里就是将 i 作为第三个位置
	- 计算所有可能的 d-c 然后放到 hashMap， d-c = nums[j]-nums[i] j [i+1,n-1]
	- 计算前面部分的 a+b 然后检查是否存在于 hashMap a+b = nums[j]+nums[i-1] j [0,i-2]

*/
func solution(nums []int) int {
	count := 0
	n := len(nums)
	countMap := make(map[int]int)
	// 这里外部循环有一个，内部有两个 for 循环，确定了最后一个 d，然后前面的确定 c
	// i 作为第三个位置，向后计算第四个位置的所有的和，然后确定前一个位置作为第二个位置，向前计算第一个位置来求和
	for i := n - 2; i > 1; i-- {
		// 这里是先确定了 c，然后计算出 nums[d]-nums[c]
		for j := i + 1; j < n; j++ {
			countMap[nums[j]-nums[i]]++
		}
		for j := i - 2; j >= 0; j-- {
			count += countMap[nums[j]+nums[i-1]]
		}
	}
	return count
}
