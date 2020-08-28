package string

/*
	647. Palindromic Substrings
	Given a string, your task is to count how many palindromic substrings in this string.
	The substrings with different start indexes or end indexes are counted as different substrings
	even they consist of same characters.

	给定一个字符串，计算字符串中有多少回文子串，只要开头或者结尾不同就可以认为是不同的字符串
	回溯法统计回文字符串 这里难道每个开头都进行判断吗，每个开头都判断的话，怎么加速，毕竟
	这里只需要开头不同，然后都向最后遍历就可以
	[ref](https://leetcode-cn.com/problems/palindromic-substrings/solution/hui-wen-zi-chuan-by-leetcode-solution/)

	计算子字符串是否回文的方式有枚举所有子字符串然后每个判断是否回文，这种方法，枚举出需要 O(N^2) 每个比较加上枚举就变成 O(N^3)

	另一种更好的方法是，枚举每个可能的回文中心，然后向两边增长，当两边的字母不满足回文就直接返回，否则继续增长
	这种枚举加上判断只需要 O(N^2)
	同时，需要处理枚举中心，考虑回文串长度是奇数，则枚举中心就是中间元素，否则，是中间的两个元素

		长度为 n 时，将生成 2n - 1 组回文中心
	第一个循环 i 从 0 到 2n-1
		初始化 l, r := i / 2, i / 2 + i % 2
		第二个循环判断 l, r 没有越界同时 l r 对应元素相等
			内部 l 向前，r 向后，然后计数 + 1
	最后返回计数

*/
func countSubstrings(s string) int {
	sLen := len(s)
	res := 0
	for i := 0; i < 2*sLen-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < sLen && s[l] == s[r] {
			l--
			r++
			res++
		}
	}
	return res
}

/*
	manacher
	一开始在前面插入 $ 结尾插入 ! 然后每个字符前后插入 #
	初始化 f slice 以及 iMax, rMax, ans 作为当前最大的回文长度的中心位置和半径以及结果
	初始化三个都是 0
	循环从 1 - n
		先计算 f[i] 如果 i < rMax 就取较小值，否则直接置 1
		然后更新扩展
		最后更新 iMax rMax
		更新 ans

*/
func countSubstrings(s string) int {
	n := len(s)
	t := "$#"
	// 插入特殊符号
	for i := 0; i < n; i++ {
		t += string(s[i]) + "#"
	}
	n = len(t)
	t += "!"

	f := make([]int, n)
	iMax, rMax, ans := 0, 0, 0
	for i := 1; i < n; i++ {
		// 初始化 f[i]
		if i <= rMax {
			f[i] = min(rMax-i+1, f[2*iMax-i])
		} else {
			f[i] = 1
		}
		// 中心扩展
		for t[i+f[i]] == t[i-f[i]] {
			f[i]++
		}
		// 维护 iMax rMax
		if i+f[i]-1 > rMax {
			iMax = i
			rMax = i + f[i] - 1
		}
		ans += f[i] / 2
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
