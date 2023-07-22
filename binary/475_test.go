package binary

import (
	"sort"
)

/*
冬天来了，你需要设计一个标准的加热器来加热所有的房子，所有的加热器有相同的加热半径。
给定 houses 以及 heaters 的坐标数组，返回让所有房子都可以被加热的最小半径
看起来给定的数组都是升序的（非降序）

遍历加热器然后判断当前的加热器需要的半径？

对一个加热器来说，其所需半径等于前后房子所需的半径
- 前面的半径等于当前的加热器和前一个加热器的距离的一半
- 后面的半径也类似
- 上面两个成立的条件是两个加热器 之前存在房子，如果不存在房子则对应的一边的半径可以是 0

需要使用一个全局的半径来进行更新

将房子和加热器的坐标都记录下来，记录下坐标同时记录下坐标下是房子还是加热器或者两个都有
- 一般的情况是，记录最左边的加热器，等到遇到下一个加热器的时候，此时两个加热器之间有房子就需要更新两者之间的距离的一半为当前需要的半径
- 两个加热器之间是否有房子这个要怎么快速判断呢
- 房子和加热器的关系最左边和最右边的单独判断一下就好了


*/
func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	var radius int
	// 先处理最左边的情况 最左边需要覆盖
	leftMostRadius := heaters[0] - houses[0]
	if leftMostRadius > radius {
		radius = leftMostRadius
	}
	// 中间部分，每个房子分别计算到两个加热器的距离然后取较小的就是需要的半径，最后和当前最大半径比较更新
	var prevPos int
	for _, v := range heaters {
		if prevPos != 0 {
			// 在加热器之间的房子，这里可以用二分查找快速找到开始的位置和结束的位置，只需要遍历这部分就可以了
			// 目标位置需要小于等于数值，因此是 寻找最右插入位置 这个等价于寻找**最右满足 <= target 的位置的右邻居**
			// 找到开始位置之后向后遍历
			i := bsr(houses, prevPos)
			for ; i < len(houses) && houses[i] < v; i++ {
				// 房子到加热器的距离使用的是较小的
				tempDist := houses[i] - prevPos
				if v-houses[i] < tempDist {
					tempDist = v - houses[i]
				}
				if tempDist > radius {
					radius = tempDist
				}
			}
		}
		prevPos = v
	}
	// 最右边处理
	rightMostRadius := houses[len(houses)-1] - heaters[len(heaters)-1]
	if rightMostRadius > radius {
		radius = rightMostRadius
	}
	return radius
}

func bsr(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
