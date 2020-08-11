package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("52 正则表达式匹配", func(t *testing.T) {
		s := "aa"
		p := "a"
		get := solution(s, p)
		want := false
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("52 正则表达式匹配2", func(t *testing.T) {
		s := "aa"
		p := "a*"
		get := solution(s, p)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("52 正则表达式匹配3", func(t *testing.T) {
		s := "ab"
		p := ".*"
		get := solution(s, p)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("52 正则表达式匹配4", func(t *testing.T) {
		s := "aab"
		p := "c*a*b*"
		get := solution(s, p)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("52 正则表达式匹配5", func(t *testing.T) {
		s := "mississippi"
		p := "mis*is*p*."
		get := solution(s, p)
		want := false
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
请实现一个函数用来匹配包含'. '和'*'的正则表达式。
模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。
在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。
同 leetcode 10
遍历 p 来匹配 s

	直接 DP 主串 A，长度 n，匹配 B，长度 m
关注 B 的最后一个字符，有三种可能：正常字符，*，.
	如果是正常字符，判断当前字符，然后比较前一个
	如果是 . 直接比较前一个
	如果是 * 表示 B[m-2] = c 可以重复 0 或者多次
		如果 A[n-1] 是 0 个 c，直接匹配 B[m-3] A[n-1]
		如果 A[n-1] 为 c 或者 c = '.'，A 向前，A[n-2] 同样和 B[m-1]
[ref](https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof/solution/zhu-xing-xiang-xi-jiang-jie-you-qian-ru-shen-by-je/)

创建长度是 (n + 1) * (m + 1) 的 dp [][]bool
遍历 n, m 的每个元素 两层
	空串默认是 true
	判断 B 的最后一个字符是 '*'
		不是
			如果 i > 0 且最后两个字符相等或者 B 的最后字符是 . dp[i][j] = dp[i=1][j-1]
		是
			如果 j >= 2 dp[i][j] = dp[i][j] || dp[i][j-2]
			如果 i >= 1 且 j >= 2 以及 B 中 * 前的字符和 A 的字符相同或者 A 的字符是 .
				dp[i][j] = dp[i][j] || dp[i-1][j]
最后返回 dp[n][m]
*/
func solution(s, p string) bool {
	byteS, byteP := []byte(s), []byte(p)
	n, m := len(byteS), len(byteP)
	dp := make([][]bool, n+1)
	for i, _ := range dp {
		dp[i] = make([]bool, m+1)
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			// 判断空主串
			if j == 0 {
				dp[i][j] = i == 0
			} else {
				// 非 *
				if byteP[j-1] != '*' {
					if i > 0 && (byteS[i-1] == byteP[j-1] || byteP[j-1] == '.') {
						dp[i][j] = dp[i-1][j-1]
					}
				} else {
					if j >= 2 {
						dp[i][j] = dp[i][j] || dp[i][j-2]
					}
					if i >= 1 && j >= 2 && (byteS[i-1] == byteP[j-2] || byteP[j-2] == '.') {
						dp[i][j] = dp[i][j] || dp[i-1][j]
					}

				}
			}
		}
	}
	return dp[n][m]
}
