package binary

import "testing"

/*
	给定整型数组 nums，返回数组里面的反转序对的数量

A reverse pair is a pair (i, j) where:
    0 <= i < j < nums.length and
    nums[i] > 2 * nums[j].

*/

/*
	类似 剑指 35 题目，先看看 35 的题目，这里只是将 35 的条件判断修改了一下罢了
	brute force 就是遍历每个然后从当前的下一个开始直到最后一个进行判断
	能否将之前比较过的 pairs 利用起来？
	需要使用归并排序
*/

func TestPro(t *testing.T) {
	t.Run("逆序对", func(t *testing.T) {
		nums := []int{1, 3, 2, 3, 1}
		get := reversePairs(nums)
		want := 2
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("逆序对2", func(t *testing.T) {
		nums := []int{2, 4, 3, 5, 1}
		get := reversePairs(nums)
		want := 3
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("逆序对3", func(t *testing.T) {
		nums := []int{2, 2, 2, 2, 2, 1, 2, 2}
		get := reversePairs(nums)
		want := 5
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) int {
	// 终止条件
	if l >= r {
		return 0
	}
	mid := l + (r-l)/2
	// 递归
	cnt := mergeSort(nums, l, mid) + mergeSort(nums, mid+1, r)
	// 结果数组
	var temp []int
	// 归并
	i, j := l, mid+1

	// 两个都没有遍历完
	for i <= mid && j <= r {
		// 左边数组当前元素的比右边的小，需要计算当前元素可以组成的逆序对了
		if nums[i] <= nums[j] {
			// 这里 j 不用算上
			cnt += calReversePairs(nums[i], nums[mid+1:j])
			temp = append(temp, nums[i])
			i++
		} else {
			temp = append(temp, nums[j])
			j++
		}
	}

	// 如果左边遍历完了，右边的就是递增的，不需要计算 cnt 了
	for ; j <= r; j++ {
		temp = append(temp, nums[j])
		j++
	}

	// 如果右边遍历完了，那左边的剩余元素都要比右边的大了，这些元素和右边的元素可能组成逆序对
	for ; i <= mid; i++ {
		temp = append(temp, nums[i])
		cnt += calReversePairs(nums[i], nums[mid+1:])
	}

	// 将数组结果进行处理
	for i := l; i <= r; i++ {
		nums[i] = temp[i-l]
	}

	return cnt
}

func calReversePairs(val int, nums []int) int {
	cnt := 0
	// 因为数组有序，如果出现不满足大小关系的就直接 break 好了
	for _, v := range nums {
		if val > 2*v {
			cnt++
		} else {
			break
		}
	}
	return cnt
}
