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
	这里也类似 327 题目
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
		nums := []int{1, 2, 3, 4, 5}
		get := reversePairs(nums)
		want := 0
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	// 这里出现负数需要怎么处理呢？
	t.Run("逆序对4", func(t *testing.T) {
		nums := []int{-5, -5}
		get := reversePairs(nums)
		want := 1
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("逆序对5", func(t *testing.T) {
		nums := []int{-5, -5, -10, -2, -6}
		get := reversePairs(nums)
		want := 6
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("逆序对6", func(t *testing.T) {
		nums := []int{2147483647, 2147483647, -2147483647, -2147483647, -2147483647, 2147483647}
		get := reversePairs(nums)
		want := 9
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
这里类似 327 题目，这里维护 i,j 两个指针，对于任意给定的 i，我们不断向右移动 j，直到 nums[i] <= 2*nums[j]（也就是不满足条件），
这里就是对于左边数组的当前位置元素，只要元素数值满足条件，就继续找右边数组的元素
这时意味着以 i 为左端点的翻转对数量为 j - m - 1，这里就是右边的已经移动了的元素的数量
上面的过程不断重复就是结果
*/

func reversePairs(nums []int) int {
	n := len(nums)
	// 元素数量不足就直接返回
	if n <= 1 {
		return 0
	}
	// 右边部分数组
	n1 := append([]int(nil), nums[:n/2]...)
	// 左边部分数组
	n2 := append([]int(nil), nums[n/2:]...)
	// 分别计算左右两边的数组的 cnt
	cnt := reversePairs(n1) + reversePairs(n2)
	// 需要计算翻转序对的数量，对左边数组的每个元素，统计右边数组的下标移动并计入总数
	// 同时因为两个数组已经是有序的，所以可以加

	// 这里只需要统计左边的数组的每个元素
	j := 0
	for _, v := range n1 {
		// 只要左边数组的当前元素满足条件，右边数组的元素就需要移动
		for j < len(n2) && v > 2*n2[j] {
			j++
		}
		// 不满足了就统计已经满足的右边数组的元素的数量
		cnt += j
	}

	// 然后归并排序
	p1, p2 := 0, 0
	for i := range nums {
		if p1 < len(n1) && p2 < len(n2) {
			if n1[p1] < n2[p2] {
				nums[i] = n1[p1]
				p1++
			} else {
				nums[i] = n2[p2]
				p2++
			}
		} else if p1 < len(n1) {
			nums[i] = n1[p1]
			p1++
		} else {
			nums[i] = n2[p2]
			p2++
		}
	}
	return cnt
}

// deprecated
func reversePairsWrong(nums []int) int {
	return mergeSortWrong(nums, 0, len(nums)-1)
}

/*
	这个没有考虑到负数的情况
	错误太多了，直接看答案了。。。
	这里是归并排序的过程中统计，针对这个题目是不合适的？
*/
func mergeSortWrong(nums []int, l, r int) int {
	// 终止条件
	if l >= r {
		return 0
	}
	mid := l + (r-l)/2
	// 递归
	cnt := mergeSortWrong(nums, l, mid) + mergeSortWrong(nums, mid+1, r)
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

	// 因为存在元素是负数的情况，所以还需要计算左边的
	for ; j <= r; j++ {
		temp = append(temp, nums[j])
		// 如果当前的右边数组第一个元素等于左边数组的最后一个元素，则需要向前计算
		equalCnt := 0
		leftIndex := mid
		// 这里左边的也需要计算了
		for leftIndex >= 0 && nums[leftIndex] > 2*nums[j] {
			leftIndex--
			equalCnt++
		}
		// 这里需要计算左边的相同的元素数量
		cnt += mid - leftIndex
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
