package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("846. Hand of Straights", func(t *testing.T) {
		nums := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
		W := 3
		want := true
		got := solution(nums, W)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("846. Hand of Straights2", func(t *testing.T) {
		nums := []int{1, 1, 2, 2, 3, 3}
		W := 2
		want := false
		got := solution2(nums, W)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	这个跟 1296 一样？应该说 1296 跟这个一样？
	这里有个问题是，如果数值很大，那可能没有办法分配到足够的空间，
	然后，也有可能浪费很多空间
*/

type NumCounter struct {
	num   int
	count int
}

func solution(hand []int, W int) bool {
	if len(hand)%W != 0 {
		return false
	}
	// 计数
	sort.Ints(hand)
	var ncSlice []NumCounter
	nc := NumCounter{hand[0], 0}
	for _, v := range hand {
		if v == nc.num {
			nc.count++
		} else {
			// 注意这里添加的是上一个元素
			ncSlice = append(ncSlice, nc)
			nc.num = v
			nc.count = 1
		}
	}
	ncSlice = append(ncSlice, nc)
	// 元素判断
	return solutionHelper(ncSlice, W)
}

func solutionHelper(ncSlice []NumCounter, W int) bool {
	ncSliceLen := len(ncSlice)
	if ncSliceLen == 0 {
		return true
	}
	if ncSliceLen < W {
		return false
	}
	next, FirstCnt := ncSlice[0].num, ncSlice[0].count
	for i := 0; i < W; i++ {
		if ncSlice[i].num != next || ncSlice[i].count < FirstCnt {
			return false
		}
		next++
		ncSlice[i].count -= FirstCnt
	}
	for next = 0; next < ncSliceLen; next++ {
		if ncSlice[next].count != 0 {
			break
		}
	}
	return solutionHelper(ncSlice[next:], W)
}

/*
	另一种方法 先排序，然后遍历
	判断是否被遍历了，计数器置 0 更新需要求的元素数值
	从当前位置遍历，判断是否是要求的数值，
	是则计数器 + 1，更新要求的，遍历置 -1 判断计数器数值
	不是则退出循环，判断数量是否足够
*/
func solution2(nums []int, k int) bool {
	numsLen := len(nums)
	// 无法划分
	if numsLen%k != 0 {
		return false
	}
	// 排序
	sort.Ints(nums)
	for i := 0; i < numsLen; i++ {
		if nums[i] != -1 {
			count := 0
			last := nums[i] - 1
			for j := i; j < numsLen; j++ {
				if nums[j] == last+1 {
					count++
					last++
					nums[j] = -1
					if count == k {
						break
					}
				}
			}
			// 里面的循环找不到足够的子数组数值
			if count != k {
				return false
			}
		}
	}
	return true
}

func solution2(hand []int, W int) bool {
	handLen := len(hand)
	// 无法划分
	if handLen%W != 0 {
		return false
	}
	// 排序
	sort.Ints(hand)
	// last 是上一个元素的数值
	// 内部循环可能会跳过重复元素，所以这里是一个个遍历
	for i := 0; i < handLen; i++ {
		// 是否需要判断，已经判断过的直接跳过
		if hand[i] != -1 {
			// 初始化计数器，所需要数值
			count := 0
			last := hand[i] - 1
			// 这里从开始位置向后面查找目标元素
			for j := i; j < handLen; j++ {
				// 注意，这里相等的时候才进行更新，
				// 不相等，需要向后面查找对应的目标元素
				if hand[j] == last+1 {
					count++
					last++
					hand[j] = -1
					// 满足数量就退出循环
					if count == W {
						break
					}
				}
			}
			// 最后的判断
			if count != W {
				return false
			}
		}
	}
	return true
}
