package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("486. Predict the Winner", func(t *testing.T) {
		want := false
		got := solution()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*

给定一个整型数组 nums，两个玩家使用这个数组进行游戏，P1 和 P2
两个人轮流，P1 开始，两个都从分数 0 开始，玩家从数组的两端挑选一个数字加入到自己的总和中，
当数组元素为空的时候，游戏结束。
如果 P1 可以赢得游戏就返回 true，如果两个人分数相同，P1 也算赢。假定两个人都使用最优选择
数组最大长度是 20，最大数值是 10^7

分析

这里选择怎么分析，如果每次都选择当前两端的最大值，是局部最优，但是不是全局最优。
比如 2 1 4 3，P1 一开始按照局部最优，选择 3，则 P2 会选择 4，导致 P1 不是选择最大的
因此这里还是暴力到 DP，暴力遍历则是每次都有两端的选择，然后每个选择都可以记录当前的总和

max/min problem

[ref](https://www.youtube.com/watch?v=Tw1k46ywN6E&list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp&t=3622s)
这里首先分析了，如果元素数量是偶数，则先手的一定可以通过计算偶数元素的和以及奇数元素的和来选择其中一个作为自己的最终和，先手的一定不会输
假设 V(i, j) 表示从位置 i 到 j 的元素的数值总和，轮到 P1 选择。对下一步，可以选择 Vi 或者 Vj

- 假设选择 Vi，则 P2 可以选择 Vi+1 或者 Vj
- 假设选择 Vj，则 P2 可以选择 Vi 或者 Vj-1

同时，对于 P2 他也选择最优的策略，则可以知道

- 选择了 Vi，对 V(i+1,j)，我们至少可以保证 min(V(i+2,j), V(i+1,j-1))
- 选择了 Vj，对 V(i,j-1)，同样保证得到 min(V(i+1,j-1), V(i,j-2))

最终得到的 DP
V(i, j) = max( (min(V(i+2, j), V(i+1, j-1)) +  Vi), ( min(V(i+1, j-1), V(i,j-2)) + Vj))

*/
var dp [][]int

func solution(nums []int) bool {
	// 偶数则一定赢
	numsLen := len(nums)
	if numsLen%2 == 0 {
		return true
	}
	numsSum := 0
	for _, v := range nums {
		numsSum += v
	}
	dp := make([][]int, numsLen)
	for i, _ := range dp {
		dp[i] = make([]int, numsLen)
		for j, _ := range dp[i] {
			dp[i][j] = -1
		}
	}
	sum := util(nums, dp, 0, numsLen-1)
	// 计算的和是否可以赢，和是否超过原来总和的一半
	return 2*sum >= numsSum
}

func util(nums []int, dp [][]int, i, j int) int {
	// 索引以及是否计算的提前返回
	if i > j {
		return 0
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	// 公式计算
	a := nums[i] + min(util(nums, dp, i+2, j), util(nums, dp, i+1, j-1))
	b := nums[j] + min(util(nums, dp, i+1, j-1), util(nums, dp, i, j-2))
	dp[i][j] = max(a, b)
	return dp[i][j]
}
