package string

/*
763. Partition Labels
A string S of lowercase English letters is given.
We want to partition this string into as many parts as possible so that each letter appears in at most one part,
and return a list of integers representing the size of these parts.
给定小写英文字母组成的字符串，我们想要将该字符串进行划分，
使得每个字母尽量只出现在一个划分部分，返回每个部分的大小数组
究竟要怎么划分，字母只出现在一个部分
保存一个当前字符数组，后面的字符如果在就加入到当前字符数组，如果不在？？？

[ref](https://leetcode-cn.com/problems/partition-labels/solution/hua-fen-zi-mu-qu-jian-by-leetcode/)
贪心 策略就是不断地选择从最左边起最小区间，可以从第一个字母开始分析。
假设第一个字母为 a，第一个区间一定包含最后一次出现的 a，这两个 a 之间可能还有其他字母，会让区间变大
因此，对于遇到的每个字母，找该字母最后一次出现的位置，用来更新最小区间

定义数组 last[char] 表示 char 字符最后一次出现的下标，定义 first 和 last 表示当前区间的首尾
如果遇到的字符最后一次出现的位置下标大于 j 就让 j = last[c] 来扩展区间
遍历到当前区间的末尾时，i == j 就把当前区间加入答案, first = i + 1 找下一个区间

创建 last 数组，并初始化 0 到 S 的长度初始化
初始化 first, last 为 0 以及结果数组
遍历 S 字符串
	先用 j,last 数组对应位置两个较大值更新 j
	如果 i == j 将结果保存到结果数组 然后 first = i + 1 作为新的开始

TODO: 这个思路感觉是数学题了
*/
func partitionLabels(S string) []int {
	res := make([]int, 0)
	lastSlice := make([]int, 26)
	// 注意这里就是每个字母的最后出现位置数组
	for i, _ := range S {
		lastSlice[S[i]-'a'] = i
	}
	// 当前处理的数组开始结束位置
	first, last := 0, 0
	for i, _ := range S {
		// 每次更新最后位置
		if lastSlice[S[i]-'a'] > last {
			last = lastSlice[S[i]-'a']
		}
		// 这里的意思是，如果当前的元素就是某个字母的最后位置，就表示处理的这段字符串的字符就满足了
		if i == last {
			res = append(res, last-first+1)
			first = i + 1
		}
	}
	return res
}
