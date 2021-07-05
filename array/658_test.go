package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("658. Find K Closest Elements", func(t *testing.T) {
		input := []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}
		k, x := 3, 5
		want := []int{3, 3, 4}
		got := solution(input, k, x)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定一个已排序的数组 arr，两个整型 k 和 x 返回数组中到 x 距离最近的 k 个数字组成的数组
	结果的数组应该升序排列，这里没有说数组元素是唯一的

	这里已经排序，这里更靠近，如果两个数组到 x 的距离相等，则数字小的更加靠近
	这里 x 和数组的边界关系分为三种
		x 在数组左边界之外，直接取数组前面的 k 个元素
		x 在数组的右边界之外，直接去数组最后的 k 个元素
		x 在数组的中间位置，需要计算距离了，所以难点在这里

	还是需要二分查找，找到合适的中间位置，然后
	双指针？从中间向两边还是从两边向中间。关键这里结果数组需要升序排列
	先思考暴力的，再优化
	如果从中间向两边 应该用这种方法
		可能存在元素从左边加入，然后从右边加入的情况，那需要左右两个数组存放，最后合并好了。同时这里可以优先将距离近的存放
		这个更加好，但是如果找不到中间的元素则需要找位置，比如下一个插入的位置这种，然后再向两边将元素添加
	如果从两边向中间
		将距离远的先保存，然后距离更近的将替换。这里替换选择需要比较两边的距离，这里也需要两个数组吗？可是没法区分左右吧？
		这个感觉简单一些，替换的话，也只需要将

*/
func solution(arr []int, k, x int) []int {
	arrLen := len(arr)
	if arrLen == 0 {
		return []int{}
	}
	// x 在数组左边界之外
	if x <= arr[0] {
		return arr[:k]
	}
	// x 在数组右边界之外
	if x > arr[arrLen-1] {
		return arr[k:]
	}
	// 找到 x 应该插入的位置 LeetCode 有一题，然后从该位置向两边查找
	left, right := 0, arrLen-1
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] == x {
			left = mid
			break
		}
		if arr[mid] < x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// 直接 left 就是这个数字插入的地方或者数字所在位置，向两边查找
	/*
		这里使用左右两个子数组保存左右两边的数字
		先判断数量是否已经够了
			不够，判断两边是否超过边界，如果其中一个超过了就直接将另一个插入到结果数组
			否则，分别计算左右两边的距离，将距离较近的插入
		最后，将右边结果数组插入到左边，返回左边的
	*/
	/*
		这里的问题是，中间向两边则结果数组两个子数组都是降序的，因此合并前将元素反转一下
		[0,0,1,2,3,3,4,7,7,8]
		3
		5
		[3,4,7] [3,3,4]
	*/
	var leftArr, rightArr []int
	// 这里分两种情况，第一种是 left 等于 x，此时 leftIndex = x，另一种则 leftIndex = left - 1, rightIndex = left
	var leftIndex, rightIndex int
	if arr[left] == x {
		leftIndex, rightIndex = left, left+1
	} else {
		leftIndex, rightIndex = left-1, left
	}
	for k > 0 {
		fmt.Println(leftIndex, rightIndex)
		if leftIndex < 0 {
			rightArr = append(rightArr, arr[rightIndex])
			rightIndex++
		} else if rightIndex > arrLen-1 {
			leftArr = append(leftArr, arr[leftIndex])
			leftIndex--
		} else {
			// 距离相等，数值更小的更加靠近
			if distance(arr[leftIndex], x) <= distance(arr[rightIndex], x) {
				leftArr = append(leftArr, arr[leftIndex])
				leftIndex--
			} else {
				rightArr = append(rightArr, arr[rightIndex])
				rightIndex++
			}
		}
		k--
	}
	fmt.Println(leftArr)
	reverseArr(leftArr)
	leftArr = append(leftArr, rightArr...)
	return leftArr
}

func distance(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func reverseArr(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
