package string

import "testing"

func TestPro(t *testing.T) {
	t.Run("293. flip game", func(t *testing.T) {
		input := "+++---+++---"
		want := "aaabcbc"
		got := decodeString(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
You are playing the following Flip Game with your friend: Given a string that contains only these two characters:
+ and -, you and your friend take turns to flip twoconsecutive "++" into "--". The game ends when a person
can no longer make a move and therefore the other person will be the winner.
Write a function to compute all possible states of the string after one valid move.

给定只包含 + 和 - 符号的字符串
将连续的 ++ 变成 -- 然后当一个人无法继续翻转的时候，另一个人算赢
编写一个函数计算字符串在一次合法的翻转后的所有可能的结果
所以这里和谁赢没有关系，只需要计算可能的结果

遍历字符串
	找到第一个 + 字符，判断下一个字符是否是 +
		如果是则可以将其翻转然后保存，然后向前一个位置
		如果不是，继续查找下一个 + 字符
*/

func solution(s string) []string {
	var res []string
	var temp []byte
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		if s[i+1] == '+' && s[i] == '+' {
			temp = []byte(s)
			temp[i] = '-'
			temp[i+1] = '-'
			res = append(res, string(temp))
		}
	}
	return res
}
