package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("414. Third Maximum Number", func(t *testing.T) {
		input := []int{17, 13, 11, 2, 3, 5, 7}
		want := []int{2, 13, 3, 11, 5, 17, 7}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	开始，所有牌都是向下的。重复下面操作直到全部翻开
		1. 将最上面的牌翻开并从桌子上拿开
		2. 如果桌子上还有牌，将下一张牌放在桌子底部
		3. 如果还有没有翻开的牌，回到第一步并重复，否则停止
	这里是将原来的数组进行重排，使得上述的过程将对所有的元素从小到大操作
	一些观察：
		相邻大小的元素需要中间相隔一个元素 这里指从小到大的元素
		1, 3, 5... 位置的元素需要按照当前数组的前面的升序排列
		2, 4, 6... 位置的元素，需要区分奇数和偶数数量来处理
			对这里来说，如果是偶数，那就是这种情况自身的重复处理 要求是 a c b d
			如果是奇数，要求是 b a c

	先排序，然后对前半部分和后半部分分别处理
*/

/*
	这里是先计算升序列表，然后把列表中的元素放到返回的结果数组中
	两个列表 一个是排序后的元素列表 一个是所有的索引的升序列表
	先排序，之后我们需要将它们放回到数组中，仅仅需要处理它们的索引
	使用一个队列来模拟整个过程
	1. 首先获得在顶部的索引
	2. 将索引的下一个索引存放到底部
	3. 重复 n 次，得到结果数组

	将第一个元素放到结果的 slice 然后将其下一个元素放到 slice 尾部，重复 n 次，n 为输入长度
*/
func solution(input []int) []int {
	inputLen := len(input)
	sort.Ints(input)
	retArr := make([]int, inputLen)
	indexArr := make([]int, inputLen)
	for i, _ := range indexArr {
		indexArr[i] = i
	}
	for i := 0; i < inputLen; i++ {
		retArr[indexArr[0]] = input[i]
		if len(indexArr) > 1 {
			next := indexArr[1]
			indexArr = indexArr[2:]
			indexArr = append(indexArr, next)
		}
	}
	return retArr
}

/*
	先排序，之后将排序后的元素存放在第二个可以存放的空位中
	跟上面的思路类似，只是上面的需要元素移动一下
*/
func solution(input []int) []int {
	inputLen := len(input)
	sort.Ints(input)
	retArr := make([]int, inputLen)
	retArr[0] = input[0]
	for i, p := 1, 0; i < inputLen; i++ {
		for j := 0; j < 2; {
			p %= inputLen
			if retArr[p] == 0 {
				j++
			}
			p += 1
		}
		p -= 1
		retArr[p] = input[i]
	}
	return retArr
}

func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
