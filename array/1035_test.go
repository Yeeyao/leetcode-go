package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1035. Uncrossed Lines", func(t *testing.T) {
		A := []int{1, 4, 2}
		B := []int{1, 2, 4}
		want := 2
		got := solution(A, B)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定两个数组，将数组元素对应相等的数进行连接，
	要求所有的连接线段不能相交，求最大的连接线数量
	首先，相同的元素且在相同位置的应该都可以直接连接
	相交的判断：相邻两条线 a[x1], a[y1] b[x2] b[y2]
	x1 x2 y1 y2 大小关系不一致 x1 > x2 y1 < y2 或者 x1 < x2 y1 > y2
	难点在于找到合适的连线让相交最少？

	为啥想到，遍历第一个数组元素，然后分别只向左或者向右找，最后比较大小 这样如果是八字形
	就只能统计到一条连线，实际是两条
	先从一个方向找，然后，连接了的标记为 1，之后换个方向找，需要判断相交
	上面都不对

	其他人的方法是，找最长的相同子数组，使用 DP 找不相交的连线就可以这样转换，妙啊，直接统计最长相同子数组的长度
	初始化需要多一个长度，然后两个索引从 1 开始，判断 i - 1 与 j - 1 是否相等，如果是就将 dp[i][j] = 1 + dp[i-1][j-1]，否则，就用 dp[i-1][j] 和 dp[i][j-1] 的较大者更新 dp[i][j]
	dp[i][j] 表示对 A 到 i 以及 对 B 到 j 元素，当前的最大连续子数组长度
	这里的 dp[i][j] 数值关系还是好难理解啊
*/
func solution(A, B []int) int {
	aLen := len(A)
	bLen := len(B)
	dp := make([][]int, aLen+1)
	for i, _ := range dp {
		dp[i] = make([]int, bLen+1)
	}
	for i := 1; i <= aLen; i++ {
		for j := 1; j <= bLen; j++ {
			// 这里容易理解
			if A[i-1] == B[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
				// 后面这里，因为不相等，所以只能从上一个最大值里面获取
			} else if dp[i-1][j] > dp[i][j-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[aLen][bLen]
}
