package array

import (
	"container/heap"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("668. Kth Smallest Number in Multiplication Table", func(t *testing.T) {
		m, n, k := 3, 3, 5
		got := solution(m, n, k)
		want := 6
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	binary search 应该类似 378，左下角开始统计，坐上部分总是小于右下部分
	这里每次统计一列，因为一列中，
		当前行的元素如果小于 mid，则它以及该列的上面的元素都小于 mid。
		当前行的元素如果大于 mid，则它需要将行数递减来向上找更小的元素。
	同时，下一列的当前行的元素一定大于当前行的元素，因此，当前列中将行数递减，和后面的列的也不冲突

*/
func solution(m, n, k int) int {
	low, high := 1, m*n
	for low < high {
		mid := low + (high-low)/2
		if isEnough(m, n, k, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// 查找当前的数值数量是否足够
func isEnough(m, n, k, mid int) bool {
	i, j := m, 1
	count := 0
	// 这里以列为主来移动
	for j < n+1 && i > 0 {
		// 如果满足就可以直接将当前列的该行以及上面的元素数量统计，然后统计下一列
		if i*j <= mid {
			count += i
			j++
			// 如果大于，则需要行数递减来判断更小的元素
		} else {
			i--
		}
	}
	return count >= k
}

/*
	这里应该和 373 类似的，只不过是换成了乘法
	matrix[i][j] = i * j
	TLE
*/
func solution2(m, n, k int) int {
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
