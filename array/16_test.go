package array

import "sort"

/*
16. 3Sum Closest
Given an array nums of n integers and an integer target, find three integers in nums such that
the sum is closest to target. Return the sum of the three integers.
You may assume that each input would have exactly one solution.
    3 <= nums.length <= 10^3
    -10^3 <= nums[i] <= 10^3
    -10^4 <= target <= 10^4

先排序，然后遍历每个元素，每个元素内部进行二分查找，如果结果相等就直接返回结果，否则更新
*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	numsLen := len(nums)
	res := nums[0] + nums[1] + nums[2]
	lom := abs(res - target)
	for i := 0; i < numsLen-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			left, right := i+1, numsLen-1
			for left < right {
				// 如果有相等的，直接就返回了
				tempSum := nums[i] + nums[left] + nums[right]
				if tempSum == target {
					return target
				} else if tempSum > target {
					tempLom := tempSum - target
					if tempLom < lom {
						lom = tempLom
						res = tempSum
					}
					right--
				} else {
					tempLom := target - tempSum
					if tempLom < lom {
						lom = tempLom
						res = tempSum
					}
					left++
				}
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
