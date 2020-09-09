package string

/*
3. Longest Substring Without Repeating Characters
Given a string s, find the length of the longest substring without repeating characters.
*/

/*
	滑动窗口
	[ref](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/wu-zhong-fu-zi-fu-de-zui-chang-zi-chuan-by-leetc-2/)
初始化 appear map 记录字符是否遇到
res, right 初始化为 0， -1
循环遍历每个字符，判断是否需要删除左边出现的字符的 map 数值
	循环移动右指针
	更新最大长度
返回最大长度
*/

func lengthOfLongestSubstring(s string) int {
	// 记录是否存在字符的 map
	// 注意这里 right 初始值和后面的一致
	appear := make(map[byte]bool)
	res, right := 0, -1
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		// 左边指针移动后，需要在 map 中删除对应的字符出现次数
		if i != 0 {
			delete(appear, s[i-1])
		}
		// 右指针移动
		for right+1 < sLen && !appear[s[right+1]] {
			appear[s[right+1]] = true
			right++
		}
		if right-i+1 > res {
			res = right - i + 1
		}
	}
	return res
}

/*
输入字符串含字母，点，符号和空格
找到最长的无重复字母字串
从第一个字母开始，使用临时字符串保存，如果遇到了 点，符号和空格或者相同的字母就保存比较结果
*/
//func lengthOfLongestSubstring(s string) int {
//	appear := make([]int, 52)
//	begin, end := 0, 0
//	res := 0
//	for i, _ := range s {
//		if isAppear(s[i], appear) {
//			if end == 0 {
//				continue
//			}
//			if end-begin+1 > res {
//				res = end - begin
//			}
//			appear = make([]int, 52)
//			begin, end = i+1, i+1
//		} else {
//			end++
//		}
//	}
//	return res
//}
//
//func isAppear(str byte, appear []int) bool {
//	if str >= 'a' && str <= 'z' {
//		if appear[str-'a'] > 0 {
//			return true
//		} else {
//			appear[str-'a']++
//			return false
//		}
//	} else if str >= 'A' && str <= 'Z' {
//		if appear[str-'A'] > 0 {
//			return true
//		} else {
//			appear[str-'A']++
//			return false
//		}
//	} else {
//		return true
//	}
//}
