package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("354. Russian Doll Envelopes", func(t *testing.T) {
		input := [][]int{}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*

	给定二维整型数组 envelopes envelopes[i] = [wi, hi]，表示信封的宽度和高度，当一个信封的宽度和高度都比另一个大(不能等于)的时候，可以装下另一个
	返回最大的可以装下的信封数量，俄罗斯套娃。（不能旋转信封）

	有点像 300 但是复杂了，直接用 300 的 dp 好了，dp[i] 表示长度为 i 的套娃数量的最后一个信封的下标，len 表示当前的最大套娃数量
	判断 i + 1 下标的元素，如果大于 dp[i] len++ dp[len] = e[i+1] 否则，找到并更新 dp[k] = e[i+1]
	先对宽度进行排序，然后就判断了
	暴力：排序后，遍历每个信封，然后计算出每个信封之后的最大套娃数量 O(n^2)
	优化：使用一个辅助数组，对每个信封，保存每个信封的下一个最小的可以装下当前信封的下标。然后一次遍历信封列表，对每个信封，不断遍历下去，统计每个信封
	的最大套娃数量，最终得到最大的数量

	[ref](https://leetcode-cn.com/problems/russian-doll-envelopes/solution/e-luo-si-tao-wa-xin-feng-wen-ti-by-leetc-wj68/)
	分析，这里先按照 w 排序，然后如果 w 都不相同，则按照 w 排序后，可以忽略 w ，只考虑 h。变成类似 300 的问题，对于 h 如果按照升序排列，
	则会出现相同的 w，然后 h 递增的情况，这样就有问题了，因此需要保证每个 w 值，只能选择一个信封，因此 h 值按照大小降序排列

	因此，先将 w 作为第一个关键字排序，然后将 h 作为第二个关键字降序排列，然后计算最长递增序列
	这里先宽度升序排列，然后相同的宽度，高度降序排列，然后就只需要考虑高度了，因为信封需要宽度和高度都满足才能装下
	使用一个数组保存高度。
	遍历排序后信封
		根据当前高度寻找插入的位置，如果找到了且在原来的数组之内，说明高度不是最高的，因此只需要更新该位置的高度
		如果找不到或者超过了数组，表示当前是新的高度，只需要追加到数组。因为当前高度是最高的，同时，宽度因为是递增的，也是最高的
		注意，这里宽度相同的情况下，因为高度最大的先加入了数组，因此后面的会直接更新高度，而不会追加
*/

func solution(envelopes [][]int) int {
	// 这里，如果相等就只取 h 大的在队列，就可以保证不会出现相同的 w
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	f := []int{}
	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}
