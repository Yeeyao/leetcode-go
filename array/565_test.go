package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("565. Array Nesting", func(t *testing.T) {
		input := []int{5, 4, 0, 3, 1, 6, 2}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	找数组元素的最长唯一的路径长度
	注意 A 的元素数值都是唯一的
	当出现重复元素的时候，这里就已经不存在循环了，所以需要跳出来
	为什么不需要将已经遍历过的元素作为开头。
	因为已经遍历过了的话，就表示至少其他的开头得到的循环比从这个开头的循环元素数量多了
	对每个元素，遍历下去
		判断元素是否为 -1 或者等于其索引值，是则跳过该元素
		否则将其值设置为 -1，同时当前长度 +1 以及将元素置 -1 。
*/
func solution(nums []int) int {
	maxLen := 1
	// 每个元素遍历
	for i, j := range nums {
		tempLen := 0
		// 需要跳过已经设置成 -1 的元素以及等于自身的也需要跳过
		if j == -1 || j == i {
			continue
		}
		// 从当前元素出发遍历下去
		for k := j; nums[k] >= 0; tempLen++ {
			temp := nums[k]
			// 将已经访问的标记为 -1
			nums[k] = -1
			k = temp
		}
		if tempLen > maxLen {
			maxLen = tempLen
		}
	}
	return maxLen
}
