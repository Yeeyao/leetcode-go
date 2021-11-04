package array

import (
	"reflect"
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("18. 4Sum", func(t *testing.T) {
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
	参考 15 需要注意这里没有说 target 一定是正数
*/
func solution(nums []int, target int) [][]int {
	var res [][]int
	// 先进行排序
	sort.Ints(nums)
	numsLen := len(nums)
	for i := 0; i < numsLen-3; i++ {
		for j := i + 1; j < numsLen-2; j++ {
			targetSum := target - nums[i] - nums[j]
			// 二分查找
			left, right := j+1, numsLen-1
			for left < right {
				// 太小了
				if nums[left]+nums[right] < targetSum {
					left++
					// 太大了
				} else if nums[left]+nums[right] > targetSum {
					right--
					// 刚好等于
				} else {
					tempRes := make([]int, 4)
					tempRes[0], tempRes[1], tempRes[2], tempRes[3] = nums[i], nums[j], nums[left], nums[right]
					// 结果保存
					res = append(res, tempRes)
					// left 部分去重
					for left < right && nums[left] == tempRes[2] {
						left++
					}
					// right 部分去重
					for left < right && nums[right] == tempRes[3] {
						right--
					}
				}
			}
			// 内层循环元素去重
			for j < numsLen-2 && nums[j] == nums[j+1] {
				j++
			}
		}
		// 最外层循环去重
		for i < numsLen-3 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
