package binary

import (
	"fmt"
)

/*

警卫在的时候珂珂不可以吃香蕉，警卫会在 H 小时后回来。
piles[i] 表示一堆香蕉中的数量，珂珂需要以固定的速度 k 根香蕉/小时速度吃，同时每堆香蕉至少需要一个小时吃完。
需要计算珂珂在警卫回来之前以最慢速度吃完所有香蕉的速度是多少

直接看题目完全想不到二分

题目可以转换为在 H 小时内以 k 的速度将所有堆的香蕉吃完。一堆香蕉的数量小于 k 也需要消耗一个小时才能吃完

这里的思路就是能力检测二分

二分解决的关键在于：
- 明确解空间。 对于这道题来说， 解空间就是 [1,max(piles)]。
- 如何收缩解空间。关键点在于如果速度 k 吃不完所有香蕉，那么所有小于等于 k 的解 都可以被排除。

实际上如果 x 不行则小于 x 的都不行
又变成 寻找最左满足 >= target，使用最左二分模板解决

*/
func minEatingSpeed(piles []int, h int) int {
	// 这里不需要排序
	// 这里 l 和 r 是香蕉的数量而不是下标
	l, r := 1, 1
	for _, v := range piles {
		if v > r {
			r = v
		}
	}

	for l <= r {
		mid := l + (r-l)/2
		// 能力检测
		if calculateHours(piles, mid, h) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func calculateHours(piles []int, mid, h int) bool {
	var res int
	for _, v := range piles {
		div := v / mid
		res += div
		if v%mid != 0 {
			res++
		}
	}
	fmt.Println(res)
	return res <= h
}
