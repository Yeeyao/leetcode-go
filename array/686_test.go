package array

import (
	"container/heap"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("686. Kth Smallest Number in Multiplication Table", func(t *testing.T) {
		m, n, k := 3, 3, 5
		got := solution(m, n, k)
		want := 6
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	这里应该和 373 类似的，只不过是换成了乘法
	matrix[i][j] = i * j
	TLE
*/

func solution(m, n, k int) int {
	h := &IHeap{}
	// 先将第一列保存到 IHeap
	if m == 1 || n == 1 {
		return k
	}
	for i := 1; i <= m && i < k; i++ {
		heap.Push(h, [3]int{i, i, 1})
	}
	// 然后就直接出队列 k - 1 次
	for i := 1; i < k && len(*h) > 1; i++ {
		now := heap.Pop(h).([3]int)
		if now[2] != n {
			heap.Push(h, [3]int{now[1] * (now[2] + 1), now[1], now[2] + 1})
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
