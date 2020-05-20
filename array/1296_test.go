package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1296. Divide Array in Sets of K Consecutive Numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 3, 4, 4, 5, 6}
		k := 4
		want := true
		got := solution(nums, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1296. Divide Array in Sets of K Consecutive Numbers2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		k := 3
		want := false
		got := solution(nums, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1296. Divide Array in Sets of K Consecutive Numbers2", func(t *testing.T) {
		nums := []int{3, 3, 2, 2, 1, 1}
		k := 3
		want := true
		got := solution(nums, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	思路类似下面的，但是使用 struct slice 来替代 map 无法顺序遍历的问题
*/

// 数值，已出现数量
type NumCounter struct {
	num   int
	count int
}

func solution(nums []int, k int) bool {
	// 无法划分
	if len(nums)%k != 0 {
		return false
	}
	// 因为后面需要按顺序访问，所以这里需要排序
	sort.Ints(nums)
	// 计数的 slice
	ncSlice := []NumCounter{}
	nc := NumCounter{nums[0], 0}
	for _, v := range nums {
		if v == nc.num {
			nc.count++
		} else {
			ncSlice = append(ncSlice, nc)
			nc.num = v
			nc.count = 1
		}
	}
	// 最后一个需要保存
	ncSlice = append(ncSlice, nc)
	// 元素判断
	return solutionHelper(ncSlice, k)
}

// 元素判断
func solutionHelper(ncSlice []NumCounter, k int) bool {
	// 所有元素遍历完了
	ncSliceLen := len(ncSlice)
	if ncSliceLen == 0 {
		return true
	}
	if ncSliceLen < k {
		return false
	}
	// 计数处理 一样的，在 k 的子数组中，
	// 所有元素的数量需要减去第一个元素的数量
	next, firstCnt := ncSlice[0].num, ncSlice[0].count
	for i := 0; i < k; i++ {
		// 不连续或者元素不足
		if ncSlice[i].num != next || ncSlice[i].count < firstCnt {
			return false
		}
		next++
		ncSlice[i].count -= firstCnt
	}
	// 找到下一个子数组的开始点
	for next = 0; next < ncSliceLen; next++ {
		if ncSlice[next].count != 0 {
			break
		}
	}
	// 下一个子数组处理
	return solutionHelper(ncSlice[next:], k)
}

/*
	类似 846

	map c 记录 A 中每个元素出现的次数
	遍历 map c 注意这里按照数值的升序遍历，
	当前数值为 i 的元素数量大于 0
	就需要将所有 i + k - 1 到 i 的元素的数量都 - i 数值元素的数量
	这里就是有多少 i 的数值的元素，对应的从 i 到 i + k - 1 的元素也需要那么多个，
	才能组成满足题目要求的那么多个数组
	不够即得到的数值小于 0，就需要直接返回 false

	上面的思路用来判断是否含有连续的升序的元素

	类似 954
	注意这里的循环处理

	这种方法的缺陷是，如果给定的数组数值很大，会没法分配到足够的空间，然后如果
	只是给定部分最大的数值，前面的空间就等于浪费了
*/
func solution2(nums []int, k int) bool {
	// 无法划分
	if len(nums)%k != 0 {
		return false
	}
	// 先将每个元素数值计数
	maxNum := 0
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	counter := make([]int, maxNum+1)
	for _, v := range nums {
		counter[v]++
	}
	for i, v := range counter {
		if v > 0 {
			for j := k - 1; j >= 0; j-- {
				// 需要注意越界问题
				if i+j >= maxNum {
					return false
				} else {
					counter[i+j] -= v
					if counter[i+j] < 0 {
						return false
					}
				}
			}
		}
	}
	return true
}
