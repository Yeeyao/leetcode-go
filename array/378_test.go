package array

import (
	"container/heap"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("378. Kth Smallest Element in a Sorted Matrix", func(t *testing.T) {
		n := 10
		got := solution(n)
		want := 6
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	有 n * n 矩阵，每行和每列的元素都升序排列，找到第 k 个最小的元素
	暴力就是直接排序然后选择
	这里需要利用元素的顺序还是选择最后一个元素来判断，找到一行，
	假如是一步步来呢，先从左上角开始，判断是向下还是向右。同时可以加快速度，某些情况下
	[ref](https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/solution/you-xu-ju-zhen-zhong-di-kxiao-de-yuan-su-by-leetco/)
*/

/*
	二叉查找方法，首先读取矩阵中的最大和最小值，作为 hi 和 lo，
		循环中 mid = (lo + hi) // 2
		计算每一行的 mid 这个数值的插入位置的总和，如果小于 k 就直接将 lo = mid + 1，否则 hi = mid
	这里其实找到插入位置的总和就是找到当前 mid 是第几个元素(按照升序)，有点暴力解法的意思。时间复杂度O(logn * n * logN)
	这里的 mid 也是也 719 一样来解释，然后这里实际是计算不大于 mid 的数量
*/
func solution(matrix [][]int, k int) int {
	n := len(matrix)
	lo, hi := matrix[0][0], matrix[n-1][n-1]
	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(matrix, mid, k, n) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// 从左下角开始统计小于等于 mid 的数字数量
func check(matrix [][]int, mid, k, n int) bool {
	i, j := n-1, 0
	num := 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			num += i + 1
			j++
		} else {
			i--
		}
	}
	// 这里如果数量小于 k 表示当前的数字还不够大
	return num < k
}

/*
	heap 方法 这里类似 23 的方法 3，就是一样的使用一个队列，一开始将一行放入到优先队列中，然后需要删除 k - 1次，头部元素就是所求的
	每次删除元素之后，需要将对应的那一列的下一个元素放入到队列中。（这里也可以先放入一列，然后一行行追加，其实就类似链表，反正行或者列都是有序的）
*/

func solution2(matrix [][]int, k int) int {
	h := &IHeap{}
	// 先将第一行保存到 IHeap
	for i := 0; i < len(matrix); i++ {
		heap.Push(h, [3]int{matrix[i][0], i, 0})
	}

	// 然后就直接出队列 k - 1 次
	for i := 0; i < k-1; i++ {
		now := heap.Pop(h).([3]int)
		if now[2] != len(matrix)-1 {
			heap.Push(h, [3]int{matrix[now[1]][now[2]+1], now[1], now[2] + 1})
		}
	}
	return heap.Pop(h).([3]int)[0]
}

// 保存数值，行，列
type IHeap [][3]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
