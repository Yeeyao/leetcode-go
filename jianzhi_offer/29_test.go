package jianzhi_offer

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("29  最小 k 个数", func(t *testing.T) {
		intSlice := []int{1, 2, 3}
		k := 2
		get := solution(intSlice, k)
		want := []int{1, 2}
		if !reflect.DeepEqual(get, want) {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
	使用类似二分查找，这里是需要最小 k 个数量，所以每次都一定要划分完
		划分后
			左边部分加上已经存在的部分等于 k 就直接加起来并返回
			左边部分加上已经存在的部分小于 k，就直接加起来并继续
			左边部分加上已经存在的部分大于 k，继续在左边部分划分
*/
func solution(intSlice []int, k int) []int {
	randomizedSelected(intSlice, 0, len(intSlice)-1, k)
	var ret []int
	for i := 0; i < k; i++ {
		ret = append(ret, intSlice[i])
	}
	return ret
}

func partition(nums []int, l, r int) int {
	// 因为上一个函数随机了 pivot 并存放在这里，所以直接将最右边作为 pivot
	pivot := nums[r]
	// 因为后面是先加所以这里先 - 1
	i := l - 1
	for j := l; j < r; j++ {
		// 元素小于 pivot 交换
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// 将 pivot 存放到位置
	nums[i+1], nums[r] = nums[r], nums[i+1]
	// 返回这次划分的数量
	return i + 1
}

// 随机划分
func randomizedPartition(nums []int, l, r int) int {
	// 随机 pivot
	i := rand.Int()%(r-1+l) + 1
	nums[r], nums[i] = nums[i], nums[r]
	return partition(nums, l, r)
}

// 快排划分
func randomizedSelected(nums []int, l, r, k int) {
	if l >= r {
		return
	}
	pos := randomizedPartition(nums, l, r)
	// 已经划分的数量
	num := pos - 1 + l
	if k == num {
		return
	} else if k < num {
		// 划分数量太多，继续划分
		randomizedSelected(nums, l, pos-1, k)
	} else {
		// 划分数量不够，继续划分
		randomizedSelected(nums, pos+1, r, k-num)
	}
}
