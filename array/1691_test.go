package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1691. Maximum Height by Stacking Cuboids ", func(t *testing.T) {
		input := []int{10, 9, 2, 5, 3, 7, 101, 18}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}

/*
	给定 n 个立方体，第 i 个 cuboids[i] = [widthi, lengthi, heighti] 选择一个立方体的子集并将它们放在一起
	对 i，j 两个立方体，如果 i 的三维都小于等于 j 则 i 可以放到 j 上，可以对立方体进行旋转来实现放置，返回堆叠的立方体的最大总高度
	排序之后，需要判断  small[i] <= small[j] and mid[i] <= mid[j] and big[i] <= big[j]
	dp[i] 表示第 i 个立方体堆叠的总高度
	有点类似 646

	354 加强版，多了一个维度，可以旋转
	因为可以旋转，这里直接对三维进行排序，取最小的两个作为一个平面（高度可以忽略了），然后就可以化简为 1996 的问题。
	但是这样是否正确？正确的吧？因为这里想要高度最大，感觉是正确的？

	立方体三维排序 O(n) 然后本身排序 O(nlogn)排序两次，最后和 1996 一样遍历
	直接将前两个作为较小的，然后最大的作为高度
	[ref](https://leetcode.com/problems/maximum-height-by-stacking-cuboids/discuss/970293/JavaC%2B%2BPython-DP-Prove-with-Explanation)
*/
func solution(cuboids [][]int) int {
	for i, _ := range cuboids {
		sort.Ints(cuboids[i])
	}
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] != cuboids[j][0] {
			return cuboids[i][0] < cuboids[j][0]
		}
		if cuboids[i][1] != cuboids[j][1] {
			return cuboids[i][1] < cuboids[j][1]
		}
		return cuboids[i][2] < cuboids[j][2]
	})
	n := len(cuboids)
	res := 0
	dp := make([]int, n)
	/*
		这里遍历到每个元素，都需要更新前面的 dp，646 是判断能否加入到子序列中组成更长的链条
		这里是判断当前元素能否加到前面的堆叠中
	*/
	for j := 0; j < n; j++ {
		// 初始化为当前的高度
		dp[j] = cuboids[j][2]
		for i := 0; i < j; i++ {
			if cuboids[i][0] <= cuboids[j][0] && cuboids[i][1] <= cuboids[j][1] && cuboids[i][2] <= cuboids[j][2] {
				dp[j] = max(dp[j], dp[i]+cuboids[j][2])
			}
		}
		res = max(res, dp[j])
	}
	return res
}

/*
 */
func solutionError(cuboids [][]int) int {
	for i, _ := range cuboids {
		sort.Ints(cuboids[i])
	}
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] != cuboids[j][0] {
			return cuboids[i][0] < cuboids[j][0]
		}
		if cuboids[i][1] != cuboids[j][1] {
			return cuboids[i][1] < cuboids[j][1]
		}
		return cuboids[i][2] < cuboids[j][2]
	})
	f := make([]int, 0)
	for i, _ := range cuboids {
		if j := sort.SearchInts(f, cuboids[i][1]); j < len(f) {
			f[j] = i
		} else {
			f = append(f, i)
		}
	}
	height := 0
	for i, _ := range f {
		height += cuboids[f[i]][2]
	}
	return height
}

func solutionError2(cuboids [][]int) int {
	for i, _ := range cuboids {
		sort.Ints(cuboids[i])
	}
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] != cuboids[j][0] {
			return cuboids[i][0] < cuboids[j][0]
		}
		if cuboids[i][1] != cuboids[j][1] {
			return cuboids[i][1] < cuboids[j][1]
		}
		return cuboids[i][2] < cuboids[j][2]
	})
	res := 0
	curWidth, curLength, curHeight := -1, -1, -1
	for _, cur := range cuboids {
		// 这里的问题在于，排序和 646 不同，因此不适用
		if curWidth <= cur[0] && curLength <= cur[1] && curHeight <= cur[2] {
			curWidth, curLength, curHeight = cur[0], cur[1], cur[2]
			res += curHeight
		}
	}
	return res
}
