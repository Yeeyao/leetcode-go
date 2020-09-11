package string

/*
5. Longest Palindromic Substring 类似 647
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
给定一个字符串，找到该字符串的最长回文子串
[ref](https://leetcode-cn.com/problems/longest-palindromic-substring/solution/zui-chang-hui-wen-zi-chuan-by-leetcode-solution/)
TODO:
*/

/*
	dp p(i, j) 表示字符串 s 第 i 到 j 个字母组成的串是否为回文串
	p(i,j) = true/false
	p(i,j) = p(i+1,j-1)&&(Si==Sj)
	p(i,i) = true p(i,i+1) = (Si==Sj)

初始化 dp 二维数组保存两个位置间字符串是否回文
遍历每个元素，以该元素作为起点
*/

func longestPalindrome(s string) string {
	sLen := len(s)
	res := ""
	dp := make([][]int, sLen)
	for i := 0; i < sLen; i++ {
		dp[i] = make([]int, sLen)
	}
	// 遍历每个元素
	for l := 0; l < sLen; l++ {
		// 每个元素作为开始向后再次遍历
		for i := 0; i+l < sLen; i++ {
			j := i + l
			// 初始值
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l+1 > len(res) {
				res = s[i : i+1+l]
			}
		}
	}
	return res
}

/*
	manacher
*/
func longestPalindrome2(s string) string {
	start, end := 0, -1
	// 字符串构造
	t := "#"
	for i := 0; i < len(s); i++ {
		t += string(s[i]) + "#"
	}
	t += "#"
	armLen := []int{}
	right, j := -1, -1
	for i := 0; i < len(s); i++ {
		var curArmLen int
		if right >= i {
			iSym := j*2 - i
			minArmLen := min(armLen[iSym], right-i)
			curArmLen = expand(s, i-minArmLen, i+minArmLen)
		} else {
			curArmLen = expand(s, i, i)
		}
		armLen = append(armLen, curArmLen)
		if i+curArmLen > right {
			j = i
			right = i + curArmLen
		}
		if curArmLen*2+1 > end-start {
			start = i - curArmLen
			end = i + curArmLen
		}
	}
	res := ""
	for i := start; i <= end; i++ {
		if s[i] != '#' {
			res += string(s[i])
		}
	}
	return res
}

// 中心扩展
func expand(s string, left, right int) int {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return (right - left - 2) / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
