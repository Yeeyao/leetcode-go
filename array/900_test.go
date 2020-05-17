package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("900. RLE Iterator", func(t *testing.T) {
		input := []int{8, 1, 5, 2, 6}
		want := 11
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	初始化时将所有的元素根据出现数量以及数值生成
	next 将遍历生成的数组，然后返回最后的数值，
	如果数组元素不足则返回 -1。同时，数组将会减少

	这里的方法不对 A 数组进行展开
	结构体保存输入数组，当前数组开始的元素的数量索引
	以及当前开头元素已经用掉的数量
*/
type RLEIterator struct {
	usedNum     int   // 当前元素已经使用的数量
	eleNumIndex int   // 当前元素数量所在的索引
	numArr      []int // 输入数组
}

func Constructor(A []int) RLEIterator {
	return RLEIterator{
		numArr: A,
	}
}

// 先判断当前的元素是否数量足够
func (this *RLEIterator) Next(n int) int {
	// 还有剩余的元素需要删除以及还足够元素
	// 如果下一个元素的数量所在索引超过输入数组长度，就表示原来数组已经不够了
	for n > 0 && this.eleNumIndex < len(this.numArr) {
		// 如果当前元素的数量就足够
		eleLeftNum := this.numArr[this.eleNumIndex] - this.usedNum
		// 注意这里相等也要返回
		if eleLeftNum >= n {
			// 已经使用数量需要加上 n
			this.usedNum += n
			return this.numArr[this.eleNumIndex+1]
		}
		// 当前的元素数量不足了，循环中 n 要减少
		n -= eleLeftNum
		// 需要变成下一个元素，已经使用元素数量置 0，遍历到下一个元素数量索引
		this.usedNum = 0
		this.eleNumIndex += 2
	}
	return -1
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(A);
 * param_1 := obj.Next(n);
 */
