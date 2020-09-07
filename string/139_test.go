package string

/*
139. Word Break
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
determine if s can be segmented into a space-separated sequence of one or more dictionary words.

给定一个非空字符串 s 以及一个字典 wordDict 包含非空单词的列表
判断 s 能否被分解成字典中的一个或者多个单词 最好是将字典构造出前缀子树 Trie

brute force 先将字典单词排序，然后每个单词取判断，如果 s 遍历完，返回 true 否则返回 false
还有一个问题，如果单词间存在某个单词是另一个单词的前缀

DP
dp[i] 表示字符串 s 前 i 个字符组成的字符串 s[0...i-1] 能否被空格拆分为若干个字典中出现的单词，
从前向后计算考虑转移方程，每次转移需要枚举包含位置 i - 1 的最后一个单词，看它是否出现在字典中
以及去除该部分的字符串是否合法
dp[i] = dp[j] && check(s[j...i-1])
初始值 dp[0] = true 判断单词是否存在于字典中，使用 map 保存每个存在于字典的单词
第一个循环 i 从 1 开始遍历输入字符串结尾（包含）
	第二个循环 j 从 0 到 i 遍历
		如果 dp[j] 以及当前 j - i 的字符串存在于字典中，dp[i] = true break 跳出循环
*/
func wordBreak(s string, wordDict []string) bool {
	wdSet := make(map[string]bool)
	for _, w := range wordDict {
		wdSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// 可以划分的依据是前面的 j 结尾的字符串和 j 到 i 的字符串都需要存在于字典中
			// 找到就直接跳出循环
			if dp[j] && wdSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
