package array

import (
	"reflect"
	"sort"
	"testing"
)

/*
18. 4Sum
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums
such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.
*/
func TestPro(t *testing.T) {
	t.Run("454. 4Sum II", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		target := 3
		want := [][]int{}
		got := solution(nums, target)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	应该使用 hashMap
*/

func solution(nums1, nums2, nums3, nums4 []int) int {
	count := 0
	countMap := make(map[int]int, 0)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			countMap[n1+n2]++
		}
	}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			if times, ok := countMap[0-n3-n4]; ok {
				count += times
			}
		}
	}
	return count
}

/*
	给定四个整型数组 nums1, nums2, nums3, nums4 长度都是 n，返回 tuples(i, j, k, l) 的数量
	其中 0 <= i, j, k, l < n nums1[i] + nums2[j] + nums3[k] + nums4[l] = 0
	应该也是先排序，然后类似 4-sum 的处理，重复的数值也可以，因为这里是下标组合的数量

	TLE
*/
func solution2(nums1, nums2, nums3, nums4 []int) int {
	// 先进行排序
	sort.Ints(nums1)
	sort.Ints(nums2)
	sort.Ints(nums3)
	sort.Ints(nums4)
	n := len(nums1)
	if nums1[0]+nums2[0]+nums3[0]+nums4[0] > 0 {
		return 0
	}
	count := 0
	for i := 0; i < n; i++ {
		// 剩余和
		targetSum := 0 - nums1[i]
		isRemain := targetSum >= 0 && targetSum <= nums2[n-1]+nums3[n-1]+nums4[n-1] ||
			targetSum < 0 && targetSum >= nums2[0]+nums3[0]+nums4[0]
		if isRemain {
			for j := 0; j < n; j++ {
				targetSum = 0 - nums1[i] - nums2[j]
				// 剩余和
				isRemain = targetSum >= 0 && targetSum <= nums3[n-1]+nums4[n-1] ||
					targetSum < 0 && targetSum >= nums3[0]+nums4[0]
				if isRemain {
					for k := 0; k < n; k++ {
						targetSum = 0 - nums1[i] - nums2[j] - nums3[k]
						// 剩余和
						isRemain = targetSum >= 0 && targetSum <= nums4[n-1] || targetSum < 0 && targetSum >= nums4[0]
						if isRemain {
							left, right := 0, n-1
							for left < right {
								mid := left + (right-left)/2
								midVal := nums4[mid]
								if targetSum > midVal {
									left = mid + 1
								} else {
									right = mid
								}
							}
							if targetSum == nums4[left] {
								count++
								sameCount := 0
								tempMid := left - 1
								for tempMid > 0 && nums4[tempMid] == nums4[left] {
									sameCount++
									tempMid--
								}
								count += sameCount
								tempMid = left + 1
								for tempMid < n && nums4[tempMid] == nums4[left] {
									sameCount++
									tempMid++
								}
								count += sameCount
							}
						}
					}
				}
			}
		}
	}
	return count
}
