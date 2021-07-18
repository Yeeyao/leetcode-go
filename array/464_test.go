package array

import (
	"fmt"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("464. Can I Win", func(t *testing.T) {
		maxChoosableIntegrr := 10
		desiredTotal := 10
		want := false
		got := solution(maxChoosableIntegrr, desiredTotal)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	100 game 的变种。两个玩家执行加法，使用从 1 到 10 的任意数字，第一个将总和加到超过 100 的玩家赢

	这里规则改成玩家不能重复使用数字。给定两个数字，最大的可选择数字以及目标和，两个玩家都使用最优的玩法？返回第一个玩家能否一定赢得游戏

	maxChoosableInteger [1,20] desiredTotal [0,300]
	这里最优的理解是，能赢就一定选择可以赢得数字。但是问题是不能赢应该选择什么，选择对手下一次即使选择剩余数字中的最大数字也不能赢的数字
	需要思考最优是什么，至少是让自己下次不会输，比如现在选择一个，那要让对方下次选择的时候不能赢。但是本次的选择又会影响到未来。
	有点像 剪枝的问题 min-max 问题

	这里最终还是需要 dp
	暴力的解法需要先想清楚

	[ref](https://leetcode.com/problems/can-i-win/discuss/95277/Java-solution-using-HashMap-with-detailed-explanation)
	这里描述使用 top-down DP 来暴力模拟每个可能的状态
	top-down DP 的关键策略是，我们需要避免重复解决子问题。我们应该使用一些策略来记住子问题的结果，这样再次遇到它们就可以马上知道结果
	通过应用 memo？我们最多可以每个子问题都计算一次，存在 O(2^n) 个子问题，因此时间复杂度是 O(2^n) ，如果没有 memo，时间复杂度将是 O(n!)

	对这个问题，关键是哪个是游戏的状态，为了任何状态确定一个唯一的结果，我们需要知道
		1. 没有选择过的数字
		2. 需要达到的目标和
	其次，1，2 是相关的，因为我们从原始的目标和中通过选择数字来见效，因此，问题就变成了如何使用 1 来描述状态

	这里的解法中，使用了一个布尔数组来记录哪个数字已经被选择了，那我们能够使用一个hashMap 来记住子问题的结果吗？使用 Map<boolean[], Boolean>
	显然是不能的，因为如果我们使用 boolean[] 作为 key，对 boolean[] 的引用不回显示 boolean[] 的实际内容？TODO: 没看懂

`	因此在问题的描述提到 maxChoosableInteger 不超过 20，也就是 boolean[] 数组的长度将小于 20 因此是可以使用 Integer 来表示 boolean[] 数组
	使用整型的饿位数表示数字是否被选择，因此可以使用 Map<Integer, Boolean> 记录子问题的结果
*/
func solution(maxChoosableInteger, desiredTotal int) bool {
}
