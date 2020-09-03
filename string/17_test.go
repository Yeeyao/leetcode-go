package string

/*
[ref](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/solution/dian-hua-hao-ma-de-zi-mu-zu-he-by-leetcode-solutio/)
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

给定一个包含 2-9 的按键，其中每个数字可以表示若干个字母，求给定按键可以组成的所有字母组合
首先构建每个数字到字母的映射，然后将给定字符串的顺序进行排列

先创建 map 进行数字和字母的匹配
直接回溯法 默认索引从 0 开始，对输入的每个数字字符串，以每个为开头进行字符串构建
*/
var res []string
var ncMap = map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}

func letterCombinations(digits string) []string {
	digitsLen := len(digits)
	if digitsLen == 0 {
		return []string{}
	}
	res = []string{}
	helper(digits, 0, "")
	return res
}

func helper(digits string, index int, temp string) {
	if index == len(digits) {
		res = append(res, temp)
	}
	digit := string(digits[index])
	chars := ncMap[digit]
	charsLen := len(chars)
	for i := 0; i < charsLen; i++ {
		helper(digits, index+1, temp+string(chars[i]))
	}
}
