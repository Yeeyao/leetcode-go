package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("32 丑数", func(t *testing.T) {
		n := 10
		get := solution(n)
		want := 12
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	丑数
	我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
	递推
		xn+1 = 2 * xa | 3 * xa | 5 * xa
		因为 xn+1 需要最接近 xn 所以需要取上面的最小值
		递推中，如果 a b c 某个数成为最小值，其数值需要 + 1
	初始
		a b c = 0 dp[0] = 1
		返回 dp[n - 1]
*/
func solution(n int) int {
	a, b, c := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		n1, n2, n3 := 2*dp[a], 3*dp[b], 5*dp[c]
		dp[i] = min(min(n1, n2), n3)
		// 需要注意这里的递增是所有的都需要判断
		if dp[i] == n1 {
			a++
		}
		if dp[i] == n2 {
			b++
		}
		if dp[i] == n3 {
			c++
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
