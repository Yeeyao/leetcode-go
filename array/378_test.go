package array

import (
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

*/

/*
	二叉查找方法，首先读取矩阵中的最大和最小值，作为 hi 和 lo，
		循环中 mid = (lo + hi) // 2
		计算每一行的 mid 这个数值的插入位置的总和，如果小于 k 就直接将 lo = mid + 1，否则 hi = mid
	这里其实找到插入位置的总和就是找到当前 mid 是第几个元素(按照升序)，有点暴力解法的意思。时间复杂度O(logn * n * logN)
*/
func solution(matrix [][]int, k int) int {

}

/*
	heap 方法 这里类似 23 的方法 3，就是一样的使用一个队列，一开始将一行放入到优先队列中，然后需要删除 k - 1次，头部元素就是所求的
	每次删除元素之后，需要将对应的那一列的下一个元素放入到队列中。（这里也可以先放入一列，然后一行行追加，其实就类似链表，反正行或者列都是有序的）
*/

func solution2(matrix [][]int, k int) int {

}
