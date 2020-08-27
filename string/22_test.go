package string

/*
 Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
For example, given n = 3, a solution set is:
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
n = 1 () 一个 n = 2 ()() (()) 两个 n = 3 五个
直接暴力，使用回溯法，初始化 slice 长度为 2 * n
pos 从 0 开始，然后每次递增一位表示下一个的数值，一次递归将 pos 分别设置为 ( 以及 )
然后继续递归处理
当当前的临时结果长度等于 n * 2 以及字符串合法就保存结果并返回

暴力改进，在添加 ( 或者 ) 之前检查合法性，提前剪枝
初始化 rune slice 以及 open, close 两个符号计数器
如果 slice 长度等于 2 * n 就直接保存结果并返回
添加 ( 前判断 open 数量，同理添加 ) 前也判断
同时注意这里先添加 ( 后递归，再添加 ) 后递归
*/
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	helper(&res, []rune{}, 0, 0, 0, n)
	return res
}

func helper(res *[]string, temp []rune, pos, open, close, n int) {
	if pos == 2*n {
		*res = append(*res, string(temp))
	}
	tempLen := len(temp)
	if open < n {
		temp = append(temp, '(')
		helper(res, temp, pos+1, open+1, close, n)
		temp = temp[:tempLen]
	}
	if close < n {
		temp = append(temp, ')')
		helper(res, temp, pos+1, open, close+1, n)
		temp = temp[:tempLen]
	}
}

/*
按照括号序列的长度递归
*/
func generateParenthesis(n int) []string {

}
