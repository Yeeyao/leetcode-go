package array

import (
	"fmt"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1010. Pairs of Songs With Total Durations Divisible by 60", func(t *testing.T) {
		input := []int{30, 20, 150, 100, 40}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	一开始排序的思路，然后每个再统计，太费时间了
	当然，一开始构造所有可以被 60 整除的数字，下面用余数的性质规避了

这里利用余数的性质，同时统计满足条件的数量

	计算 time % 60
	t % 60 将得到结果 0 - 59
	统计在数组 arr 中的余数出现的频率
	我们想要知道，对每个 t ，有多少 x 满足 (t + x) % 60 == 0
	其中 x % 60 == 60 - t % 60 对大多数情况成立
	其中 60 - t % 60 将得到 1 - 60，则 (60 - t % 60) % 60 得到 0 - 59

	另一种想法是 x % 60 = (600 - t) % 60 因为元素的数值是 [1, 500]
	这里的思路就是当前的往前面的找，前面的会先保存起来数量
*/
func solution(input []int) int {
	arr := make([]int, 60)
	sum := 0
	// 先加上当前的数量，再对应递增
	// 对于当前的数值，找到当前与之相加可以整除 60 的数的数量
	// 之后，将当前数值对应的余数递增
	for _, t := range input {
		sum += arr[(600-t)%60]
		arr[t%60]++
	}
	return sum
}

// 速度改进
func solution(time []int) int {
	ct := make([]int, 60)
	// 先全部都保存
	for _, t := range time {
		ct[t%60]++
	}
	// 先中间部分全部都加起来，统计数量
	s := 0
	for i := 1; i < 30; i++ {
		s += ct[i] * ct[60-i]
	}
	// 自己可以组合的，需要另外统计
	s += ct[0] * (ct[0] - 1) / 2
	s += ct[30] * (ct[30] - 1) / 2
	return s
}
