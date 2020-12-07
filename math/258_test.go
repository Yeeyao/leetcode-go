package math

import "testing"

func TestPro(t *testing.T) {
	t.Run("258. Add Digits", func(t *testing.T) {
		num := 9
		want := true
		got := solution(num)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定一个非负整型 num，重复将它的位数累加直到和只有一位
	最后返回该只有一位的和 不是直接将所有位数累加然后对 9 取余？
	不使用循环或者递归 4 个 hint...
	[ref](https://www.zhihu.com/question/30972581)
	数根
*/
func solution(num int) int {
	return (num-1)%9 + 1
}
