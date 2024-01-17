package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("35  数组中的逆序对", func(t *testing.T) {
		nums := []int{7, 5, 6, 4}
		get := solution(nums)
		want := 5
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("35  数组中的逆序对2", func(t *testing.T) {
		nums := []int{7, 7, 7, 5, 5, 5, 6, 4}
		get := solution(nums)
		want := 19
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("35  数组中的逆序对3", func(t *testing.T) {
		nums := []int{2, 2, 2, 2, 2, 1, 2, 2}
		get := solution(nums)
		want := 5
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
	归并排序
		对于左边的序对，如果当前的节点比右边的当前节点小，则表示右边当前节点的左边部分都是逆序对(和左边节点组成)
	0 1 2 3 4 5 6 7
	7 7 7 5 5 5 6 4

	2 2 2 2 2 1 2 2
*/
func solution(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) int {
	// 中止条件
	if l >= r {
		return 0
	}
	mid := l + (r-l)/2
	// 这里需要带上 mid
	cnt := mergeSort(nums, l, mid) + mergeSort(nums, mid+1, r)
	// 中间数组保存临时结果
	tmp := []int{}
	// 左右两个数组的开头
	i, j := l, mid+1
	// 左右两个数组都没有遍历完
	for i <= mid && j <= r {
		// 注意这里等于 如果是取消等号，假设右边的遍历下一个，左边将计算到有逆序对，但是实际是没有的
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			// 这里 j - (mid + 1) 是右边数组的偏移，也就是右边目前有多少个元素已经放入到结果数组
			// 出现左边更小的时候，因为对于左边的数组的当前元素 x 来说，右边已经遍历过的元素数量就是当前元素 x 组成的逆序对的数量了，所以这里就计算逆序对的数量
			cnt += j - (mid + 1)
			i++
		} else {
			// 如果右边的数组元素小，就将右边的保存到结果数组
			tmp = append(tmp, nums[j])
			j++
		}
	}

	// 右边的遍历完了，表示左边的剩余这部分比右边的都要大，所以左边每一个剩余的都是逆序对的一个元素，其中的每个元素组成的逆序对数量都是右边的数组的长度
	// 因为左边数组是升序的，这里剩余的部分的右边不需要计算，那就只需要计算右边数组了
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		cnt += r - (mid + 1) + 1
	}

	// 左边的遍历完了
	for ; j <= r; j++ {
		tmp = append(tmp, nums[j])
	}
	// 重新赋值
	for i := l; i <= r; i++ {
		nums[i] = tmp[i-l]
	}

	return cnt
}
