package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("14 链表中倒数第k个结点", func(t *testing.T) {
		n := 3
		want := 4
		got := solution(n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	输入一个链表，输出该链表中倒数第k个结点
	将链表入栈，然后 pop k 次元素后得到的就是所求
 */
func solution(nums []int) {

}
