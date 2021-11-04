package array

import (
	"testing"
)

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
	// 对每个 i 作为 c 位置，循环内，所有的 d-c 以及 a+b 都计算了，先计算 d-c，然后统计 a+b
	for i := n - 2; i > 1; i-- {
		// 这里是先确定了 c，然后计算出 nums[d]-nums[c]，每个 i 后面的都需要计算出 d-c
		for j := i + 1; j < n; j++ {
			countMap[nums[j]-nums[i]]++
		}
		// 这边计算 a+b 因为 c 的位置是 i，因此这里是 i-1 和 i-2，每个 i-1 之前的都计算 a+b，然后计数
		for j := i - 2; j >= 0; j-- {
			count += countMap[nums[j]+nums[i-1]]
		}
	}
	return count
}
