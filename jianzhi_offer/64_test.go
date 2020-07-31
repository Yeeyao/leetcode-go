package jianzhi_offer

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("64  滑动窗口的最大值", func(t *testing.T) {
		nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
		k := 3
		get := solution(nums, k)
		want := []int{3, 3, 5, 5, 6, 7}
		if reflect.DeepEqual(get, want) {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("64  滑动窗口的最大值2", func(t *testing.T) {
		nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
		k := 3
		get := solution2(nums, k)
		want := []int{3, 3, 5, 5, 6, 7}
		if reflect.DeepEqual(get, want) {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
	同 leetcode 239
	一开始找最大值，记录下最大值以及位置
	然后每次增加一个元素判断新的最大值
	单调队列，遍历下去，每次取队列的尾部元素并将队列的首部元素出队列
*/
func solution(nums []int, k int) []int {
	var res []int
	queue := make([]int, k)
	for _, n := range nums {
		enQueue(queue, k, n)
		if len(queue) == k {
			res = append(res, queue[k-1])
		}
	}
	return res
}

/*
	入队列，找到第一个可以存放的位置
*/
func enQueue(queue []int, k, ele int) {
	qLen := len(queue)
	// 找到第一个可以插入的位置
	i := 0
	for ; i < qLen && queue[i] < ele; i++ {
	}
	// 找不到可以插入的位置
	if i == 0 {
		// 无法入队列
		if qLen == k {
			return
		} else {
			// 后面的元素向后移动
			for j := qLen - 1; j >= 0; j-- {
				queue[j+1] = queue[j]
			}
			queue[0] = ele
		}
	} else {
		// 原来的队列还没有满，后面的元素全部向后移动
		if qLen < k {
			for j := qLen; j > i; j-- {
				queue[j+1] = queue[j]
			}
			queue[i] = ele
		} else {
			// 原来队列已经满了 前面的需要向前移动
			for j := i - 1; j > 0; j-- {
				queue[j-1] = queue[j]
			}
			queue[i-1] = ele
		}
	}
}

/*
	多次分配空间的版本
*/
func solution2(nums []int, k int) []int {
	var res []int
	queue := make([]int, k)
	for _, n := range nums {
		queue2 := enQueue2(queue, k, n)
		if len(queue2) == k {
			res = append(res, queue2[k-1])
		}
	}
	return res
}

/*
	入队列，找到第一个可以存放的位置
	因为这里需要再次使用 queue 所以需要多次更新
*/
func enQueue2(queue []int, k, ele int) []int {
	qLen := len(queue)
	// 找到第一个可以插入的位置
	i := 0
	for ; i < qLen && queue[i] < ele; i++ {
	}
	// 将前后连接
	queue1 := queue[0:i]
	queue1 = append(queue1, ele)
	queue1 = append(queue1, queue[i:]...)
	queue = queue1[0:k]
	return queue
}
