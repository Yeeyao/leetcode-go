package array

import (
	"testing"
	"sort"
	"fmt"
)

func TestPro(t *testing.T) {
	t.Run("581. Shortest Unsorted Continuous Subarray", func(t *testing.T) {
		input := []int{2, 6, 4, 8, 10, 9, 15}
		want := 5
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("581. Shortest Unsorted Continuous Subarray2", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		want := 0
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("581. Shortest Unsorted Continuous Subarray3", func(t *testing.T) {
		input := []int{2, 1}
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("581. Shortest Unsorted Continuous Subarray4", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		want := 0
		got := findUnsortedSubarray2(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
/*
排序
	如果复制原数组后进行排序，然后逐个元素判断，如果相等就表示元素已经在排好序的位置了，不需要处理；
	不相等，则需要更新坐标的最小值和最大值，这两个之间的元素就是需要排序的

	先复制一份 nums 然后排序
	初始化 start, end 为 数组长度和 0
	遍历排序后数组，如果原数组和当前数组的相同位置数值不同就更新最大的位置和最小的位置
	最后如果结束位置大于开始位置，差值大于等于 0 就返回两者差 + 1 否则返回 0
*/
func findUnsortedSubarray(nums []int)int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	begin, end := len(nums), 0
	for i := 0; i < len(nums); i++{
		if nums[i] != sortedNums[i]{
			if i < begin {
				begin = i
			}
			if i > end {
				end = i
			}
		}
	}
	if end - begin >= 0 {
		return end - begin + 1
	}
	return 0
}


/*
不使用额外空间
	无序子数组中最小元素的正确位置可以决定左边界，最大元素的正确位置可以决定右边界。
	遍历数组查找，第一次从开头向结尾找，如果元素升序就跳过（因为满足了升序），如果元素开始降序，就直接找到降序后的最小值
	同理，从结尾向开头遍历，如果元素降序就跳过，开始升序，直接找升序后的最大值
	再从头开始遍历，找到第一个大于最小元素的位置，同理从尾部向前遍历，找到第一个小于最大元素的位置
	两个位置之间的元素数量就是所求
*/
func findUnsortedSubarray2(nums []int)int {
	numsLen := len(nums)
	// 是否改变了排序
	flag := false
	const intMax = int(^uint(0) >> 1)
	const intMin = ^intMax
	min, max := intMax, intMin
	for i := 1; i < numsLen; i++{
		if nums[i - 1] > nums[i] {
			flag = true
		}
		if flag {
			if nums[i] < min {
				min = nums[i]
			}
		}
	}
	flag = false
	for i := numsLen-2; i >= 0; i--{
		if nums[i] > nums[i+1] {
			flag = true
		} 
		if flag {
			if nums[i] > max {
				max = nums[i]
			}
		}
	}
	// 这两个初始值需要都是 0
	var left, right int
	for left := 0; left < numsLen; left++{
		if min < nums[left] {
			break
		}
	}
	for right := numsLen - 1; right >= 0; right-- {
		if max > nums[right] {
			break
		}
	}
	// 注意这里如果是 0 就返回 0
	if right - left <= 0 {
		return 0
	}
	return right - left + 1
}

/*
	找到最短连续的未排序的子数组，该数组有序则整体数组有序
	从左到右遍历，找到比当前 max 小的最近一个元素
	从右到做遍历，找到比当前 min 大的最近一个元素
	其实是右边需要找到最小的数值，左边找最大的数值，它们组成的子数组排序就行
	[参考](https://leetcode.com/problems/shortest-unsorted-continuous-subarray/discuss/103057/Java-O(n)-Time-O(1)-Space)
*/
func solution(input []int) int {
	inputLen := len(input)
	// subarray index
	left, right := -1, -2
	// 一开始是设置成边界的两个元素值
	min, max := input[inputLen-1], input[0]
	// 注意这里的遍历顺序，左右两个指针会交叉
	for i := 1; i < inputLen; i++ {
		rightIndex := inputLen - i - 1
		// 先更新最大最小值
		if input[i] > max {
			max = input[i]
		}
		if input[rightIndex] < min {
			min = input[rightIndex]
		}
		// 更新子数组的开始结束位置索引
		/*
			这里的意思是，需要排序的数组的结束位置，是在当前最大值的左边的
			开始位置，是在当前最小值的右边的。所以需要不断更新
		*/
		if input[i] < max {
			right = i
		}
		if input[rightIndex] > min {
			left = rightIndex
		}
	}
	return right - left + 1
}

func solution2(input []int) int {
	inputLen := len(input)
	if inputLen == 0 || inputLen == 1 {
		return 0
	}
	// 找到比最大值小的最右边元素
	const intMax = int(^uint(0) >> 1)
	const intMin = ^intMax
	max, end := intMin, -2
	for i, v := range input {
		if v > max {
			max = v
		}
		if v < max {
			end = i
		}
	}
	// 找到比最小值大的最左边元素
	min, begin := intMax, -1
	for i := inputLen - 1; i >= 0; i-- {
		if input[i] < min {
			min = input[i]
		}
		if input[i] > min {
			begin = i
		}
	}
	return end - begin + 1
}
