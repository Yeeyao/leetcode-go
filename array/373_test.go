package array

import (
	"container/heap"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("373. Find K Pairs with Smallest Sums", func(t *testing.T) {
		want := [][]int{}
		nums1 := []int{1, 2, 3}
		nums2 := []int{1, 2, 3}
		k := 1
		got := solution(nums1, nums2, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定两个升序排序的整型数组 nums1, nums2 以及整型数字 k，定义 pair (u, v)，其中一个元素从第一个数组中获取，另一个元素从第二个数组中获取
	返回 pair (u, v) 和最小的 k 对。
	暴力方法是直接将枚举所有的序对，然后按照总和排序

	优先队列做法
		count = 0，count < k，递增，
		每次将 i + j == count 的元素放入到优先队列中，然后取出头部元素保存到结果，这里队列元素 {sum, i, j} i,j 记录 sum 的索引
	二叉查找，一般有序的数组都可以使用二叉查找，这里同样计算最小和最大的和，然后和小于等于 mid 的序对的数量和 k 比较

*/

// 优先队列 5% 主要问题是内存访问不友好导致频繁切页
type ele [][3]int

func (e ele) Len() int           { return len(e) }
func (e ele) Less(i, j int) bool { return e[i][0] < e[j][0] }
func (e ele) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func (e *ele) Push(x interface{}) {
	*e = append(*e, x.([3]int))
}

func (e *ele) Pop() interface{} {
	old := *e
	n := len(old)
	x := old[n-1]
	*e = old[0 : n-1]
	return x
}

func solution(nums1, nums2 []int, k int) [][]int {
	h := &ele{}
	count := 0
	n, m := len(nums1), len(nums2)
	if k > n*m {
		k = n * m
	}
	res := make([][]int, 0)
	for count < k {
		for i := 0; i < n && i < k; i++ {
			dest := count - i
			if dest >= 0 && dest < m {
				heap.Push(h, [3]int{nums1[i] + nums2[dest], i, dest})
			}
		}
		topE := heap.Pop(h).([3]int)
		res = append(res, []int{nums1[topE[1]], nums2[topE[2]]})
		count++
	}
	return res
}
