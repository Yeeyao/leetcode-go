package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run(" 1011. Capacity To Ship Packages Within D Days ", func(t *testing.T) {
		weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		D := 5
		want := 15
		got := solution(weights, D)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	这个题目的意思是，给定一个箱子重量的数组，指定的天数
	使用一艘船来搬运，要求在指定天数内全部可以搬运完，同时需要按照给定的箱子顺序搬运。
	每天可以搬运多个箱子，求满足条件下，船的最小承重
	将箱子数组分成 D 部分，找所有部分中的最大重量和

	不考虑天数情况下，结果是在 [max(A), sum(A)] 中，第一个表示 输入的最大元素

	直接考虑每天的重量
		怎么判断是否已经满足要求了 停止
		怎么更新 累加当前的重量
		什么时候停止 左右越界

	找到最大的重量
	binary search
	从左边开始找连续的和，计算 mid，记录当前所需的元素数量以及当前累加和
	如果当前累加和加上当前元素大于 mid 就表示还需要多一天

	这里是利用二叉搜索的思想，找到满足条件的一天的重量
	针对每个重量，遍历整个给定的数组，找该重量下，数组所有箱子搬完所需的天数，
	如果大于给定的 D 则表示重量太小了，需要增加。如果小于给定的 D 就表示满足条件，
	但是题目要求找到最小的一天的重量，所以还要继续判断下去

*/

func solution(weights []int, D int) int {
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}
	for left < right {
		mid := (left + right) / 2
		days := 1
		tempSum := 0
		// notes 重新开始遍历找连续的和
		for _, w := range weights {
			// 已经超过当天的重量了，需要增加当前的天数
			tempSum += w
			if tempSum > mid {
				days++
				tempSum = w
			}
		}
		// 当前的累加和已经超过给定的天数了，表示当前的每天重量太小了
		// notes
		if days > D {
			left = mid + 1
		} else {
			// 虽然重量已经满足了，但是需要继续更新来尝试找到更小的
			right = mid
		}
	}
	return left
}
