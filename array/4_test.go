package array

import (
	"testing"
)

func TestPro(t *testing.T) {

	t.Run("leetcode 238  Product of Array Except Self1", func(t *testing.T) {
		nums1 := []int{1}
		nums2 := []int{3}
		want := 2.0
		got := solution(nums1, nums2)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("leetcode 238  Product of Array Except Self2", func(t *testing.T) {
		nums1 := []int{1, 5}
		nums2 := []int{2, 3, 4}
		want := 3.0
		got := solution(nums1, nums2)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("leetcode 238  Product of Array Except Self3", func(t *testing.T) {
		nums1 := []int{1, 3, 4}
		nums2 := []int{2}
		want := 2.5
		got := solution(nums1, nums2)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("leetcode 238  Product of Array Except Self4", func(t *testing.T) {
		nums1 := []int{1, 2, 4, 5, 6, 7, 8}
		nums2 := []int{3}
		want := 4.5
		got := solution(nums1, nums2)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("leetcode 238  Product of Array Except Self5", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
		nums2 := []int{3, 4}
		want := 6.5
		got := solution(nums1, nums2)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("leetcode 238  Product of Array Except Self6", func(t *testing.T) {
		nums1 := []int{}
		nums2 := []int{1}
		want := 1.0
		got := solution(nums1, nums2)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
元素数量是奇数或者偶数计算中间数的判断
当遇到中间的数的时候，如果总数是奇数，则直接取该数并返回了
如果总数是偶数，则需要再找下一个元素
还需要确定的是，当前 i 是哪个数组中的，直接存储当前的元素值
*/
func solution(nums1, nums2 []int) float64 {
	nums1Len, nums2Len := len(nums1), len(nums2)
	if nums1Len == 0 {
		medianNum := nums2Len / 2
		if nums2Len%2 == 0 {
			return (float64(nums2[medianNum-1]) + float64(nums2[medianNum])) / 2
		} else {
			return float64(nums2[medianNum])
		}
	}
	if nums2Len == 0 {
		medianNum := nums1Len / 2
		if nums1Len%2 == 0 {
			return (float64(nums1[medianNum-1]) + float64(nums1[medianNum])) / 2
		} else {
			return float64(nums1[medianNum])
		}
	}
	totalLen := nums1Len + nums2Len
	// 这里需要区别是两个元素还是一个元素
	medianIndex := totalLen / 2
	nums1Index, nums2Index := 0, 0
	// 当前中位数值
	medianNum, medianNum2 := 0, 0
	// 这里只是遍历到中间元素的前一个
	for i := 0; i < medianIndex; i++ {
		// 第一个数组已经遍历完了
		if nums1Index > nums1Len-1 {
			medianNum = nums2[nums2Index]
			nums2Index++
			// 防止将上一个次数设置
			medianNum2 = nums2[nums2Index]
		} else if nums2Index > nums2Len-1 {
			// 第二个数组已经遍历完了
			medianNum = nums1[nums1Index]
			nums1Index++
			medianNum2 = nums1[nums1Index]
		} else if nums1[nums1Index] < nums2[nums2Index] {
			// 两个都没遍历完
			medianNum = nums1[nums1Index]
			nums1Index++
			if nums1Index < nums1Len {
				// 这里需要找到下一个较小的元素
				if nums1[nums1Index] < nums2[nums2Index] {
					medianNum2 = nums1[nums1Index]
				} else {
					medianNum2 = nums2[nums2Index]
				}
			} else {
				medianNum2 = nums2[nums2Index]
			}
		} else {
			medianNum = nums2[nums2Index]
			nums2Index++
			if nums2Index < nums2Len {
				// 这里需要找到下一个较小的元素
				if nums2[nums2Index] < nums1[nums1Index] {
					medianNum2 = nums2[nums2Index]
				} else {
					medianNum2 = nums1[nums1Index]
				}
			} else {
				// 第二个数组已经遍历完了
				medianNum2 = nums1[nums1Index]
			}
		}
	}
	if totalLen%2 == 0 {
		return (float64(medianNum) + float64(medianNum2)) / 2
	} else {
		// 因为遍历的过程没有遍历到中间的元素的，所以奇数时取的是下一个中位数
		return float64(medianNum2)
	}
}
