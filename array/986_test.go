package array

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("977. Squares of a Sorted Array", func(t *testing.T) {
		A := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
		B := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
		want := [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}}
		got := solution(A, B)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
986. Interval List Intersections
[ques](https://leetcode.com/problems/interval-list-intersections/)
Given two lists of closed intervals, each list of intervals is pairwise disjoint and in sorted order.
Return the intersection of these two interval lists.
更简单直观的方法
[ref](https://leetcode-cn.com/problems/interval-list-intersections/solution/qu-jian-lie-biao-de-jiao-ji-by-leetcode/)
思路
	在两个数组给定的所有区间中，假设拥有最小末端点的区间是 A[0]
	然后，在数组 B 的区间中， A[0] 只可能与数组 B 中的至多一个区间相交。
	（如果 B 中存在两个区间均与 A[0] 相交，那么它们将共同包含 A[0] 的末端点，但是 B 中的区间应该是不相交的，所以存在矛盾）todo:
算法
	如果 A[0] 拥有最小的末端点，那么它只可能与 B[0] 相交。然后我们就可以删除区间 A[0]，因为它不能与其他任何区间再相交了。
	如果 B[0] 拥有最小的末端点，那么它只可能与区间 A[0] 相交，然后我们就可以将 B[0] 删除，因为它无法再与其他区间相交了。
	我们用两个指针 i 与 j 来模拟完成删除 A[0] 或 B[0] 的操作。

	求完一个交集区间，较早结束的子区间，不可能再和其他子区间有重叠，它的指针要移动
	较长的子区间还可能和别人重叠，它的指针暂时不动
	时间和空间复杂度 O(M + N)
*/
func solution(A, B [][]int) [][]int {
	res := [][]int{}
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		start := max(A[i][0], B[j][0])
		end := min(A[i][1], B[j][1])
		if start <= end {
			res = append(res, []int{start, end})
		}
		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

/*
986. Interval List Intersections
[ques](https://leetcode.com/problems/interval-list-intersections/)
Given two lists of closed intervals, each list of intervals is pairwise disjoint and in sorted order.
Return the intersection of these two interval lists.

给定两个闭区间列表，列表内区间维持升序同时没有交集，返回两个列表的集合的交集列表
brute force：
	遍历其中一个列表 A，然后对每个区间 Ai，判断另一个列表 B 的区间
	如果 Bi 起始位置大于 Ai 结束位置或者 Bi 结束位置小于 Ai 起始位置则直接跳过该 Bi，这里是跳过，不是进入下一个
		注意这里不需要每次都从另一个列表的第一个区间开始，这里开始的区间可以通过遍历过程不断更新
		如果 Bi 结束位置大于 Ai 结束位置，则不需要跳到下一个 Bi
	出现相交就保存两个起始位置和两个结束位置的两个较小者作为相交区间
	上述过程一直遍历直到遍历完第一个列表
*/
/*func solution(A, B [][]int) [][]int {
	aLen, bLen := len(A), len(B)
	if aLen == 0 || bLen == 0 {
		return [][]int{}
	}
	i, j := 0, 0
	res := make([][]int, 0)
	for i < aLen {
		AiBegin, AiEnd := A[i][0], A[i][1]
		BjBegin, BjEnd := B[j][0], B[j][1]
		if BjBegin > AiEnd {
			i++
		} else if BjEnd < AiBegin {
			j++
		} else {
			addBegin := max(AiBegin, BjBegin)
			addEnd := min(AiEnd, BjEnd)
			res = append(res, []int{addBegin, addEnd})
			i++
			j++
			if A[i][0] == B[j-1][1] {
				res = append(res, []int{A[i][0], A[i][0]})
			} else if B[j-1][0] == A[i][1] {
				res = append(res, []int{A[i][1], A[i][1]})
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}*/
