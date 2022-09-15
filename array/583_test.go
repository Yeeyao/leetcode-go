package array

/*

Given two strings word1 and word2, return the minimum number of steps required to make word1 and word2 the same.
In one step, you can delete exactly one character in either string.
给定两个字符串，返回将两个字符串变成相同字符串所需的最少步骤，一个步骤允许在任意字符串中删除一个字母

类似 712，这里同样需要选择当出现两个字母不同的情况下，需要删除哪个字母的问题
`
一样分析，对 word1 的第 i 个以及 word2 的第 j 个，判断是否相等
	如果相等，就不需要删除，步骤不需要增加
	如果不相等，有两种删除方式，一种是删除 i 另一种是删除 j

需要注意这里的 dp 与上面的分析的差异

同样使用 dp[i][j] 表示 word1[:i] 和 word2[:j] 的字符串变成相同的所需要的删除的步骤
	如果相等，则 dp[i][j] = dp[i-1][j-1]
	不相等，则 dp[i][j] = min(dp[i][j-1] dp[i-1][j]) + 1

*/

func solution(word1, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= m; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[n][m]
}
