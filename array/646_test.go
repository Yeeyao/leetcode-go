package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("646. Maximum Length of Pair Chain", func(t *testing.T) {
		input := []int{10, 9, 2, 5, 3, 7, 101, 18}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}

/*
	给定 n 个序对 pairs 其中 pairs[i] = [lefti, righti] lefti < righti
	p2[c,d] 将会跟随 p1[a,b] 如果 b < c 成立，按照这种方式可以组成一个 pairs 的链条，就最长的链条长度
	可以理解为序对之间的取值没有交集就可以组成 chain
	[ref](https://leetcode-cn.com/problems/maximum-length-of-pair-chain/solution/zui-chang-shu-dui-lian-by-leetcode/)

	使用贪心策略，在所有下一个数对中选择第二个数最小的添加到链表
	这里 cur 理解为当前的坐标，然后不断判断变化， 就是不断跳动

	这里先按照 right 降序排列，然后就不断比较下一个 pair 的 left 和上一个的 right
*/

func solution(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		a, b := pairs[i], pairs[j]
		return a[1] < b[1]
	})
	cur, res := -1001, 0
	for _, p := range pairs {
		if cur < p[0] {
			cur = p[1]
			res++
		}
	}
	return res
}

/*
	动态规划 在一个长度为 k，以 pairs[i] 结尾的数对链中 如果 pairs[i][1] < pairs[j][0] 则可以将数对放到链表，然后长度 k + 1
	根据第一个数对 pair 进行排序 dp[i] 保存以 pairs[i] 结尾的最长链表长度 i < j 且 pairs[i][1] < pairs[j][0] 更新
	dp[j] = max(dp[j], dp[i] + 1), 初始化 dp[i] = 1
*/
func solution2(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		a, b := pairs[i], pairs[j]
		return a[0] < b[0]
	})
	n := len(pairs)
	dp := make([]int, n)
	for i, _ := range dp {
		dp[i] = 1
	}
	// 遍历同时，需要向前计算
	for j := 1; j < n; j++ {
		// 每遍历一个元素，都需要从头开始判断是否可以加到前面的元素的链表中
		for i := 0; i < j; i++ {
			if pairs[i][1] < pairs[j][0] {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
	}
	// 最后需要全部判断
	res := 0
	for _, v := range dp {
		if v > res {
			res = v
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
