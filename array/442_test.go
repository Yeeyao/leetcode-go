package array

import "testing"

func TestPro(t *testing.T) {
	t.Run(" 442. Find All Duplicates in an Array ", func(t *testing.T) {
		input := []int{4, 3, 2, 7, 8, 2, 3, 1}
		want := []int{2, 3}
		got := solution(input)
		if IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	这里，使用数组的下标和内容来存放已经存在的数组元素
	对一个元素检查它的值
		如果是 1 表示已经出现过两次了，直接将其置 1
		如果是 0 表示已经出现过一次了，将其置 1
		如果不是上面两个值，表示这个值没有出现过，则将这个值置 0 然后将原来的值继续进行判断
	将他作为目标位置的下标，检查对应位置的元素的值，如果是 0 表示已经
*/

/*
	当前元素的值作为下标去访问对应的值
	为 0 表示之前出现过一次了
*/
// func solution(nums []int) []int {
// 	numsLen := len(nums)
// 	for i := 0; i < numsLen; i++ {
// 		// 有一个元素存放在合适的位置就停止。
// 		for nums[i] != i+1 {
// 			// 有一个元素存放在合适的位置就停止。
// 			// 这里要求是 nums[i] 存放在 nums[nums[i] - 1]
// 			if nums[nums[i]-1] == nums[i] {
// 				break
// 			}
// 			// 这里为什么需要这样交换 这个位置的数值和数值的位置的数值
// 			// 这里交换时
// 			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
// 		}
// 	}
// 	retArr := []int{}
// 	for i, v := range nums {
// 		if v != i+1 {
// 			retArr = append(retArr, v)
// 		}
// 	}
// 	return retArr
// }

/*
	这种思路是，元素值对应的下标值是元素值-1,遍历元素，然后获得该元素的索引，
	之后，判断该索引的值是否已经是负数，是则表示已经出现过了，直接保存到返回数组
	不是，则不需要保存。然后将该下标值的数值置为相反数
	注意这里题目的中元素最多只会出现两次
*/
func solution(nums []int) []int {
	var retArr []int
	for _, v := range nums {
		// 下标值需要取非负数
		var value int
		if v < 0 {
			value = -v
		} else {
			value = v
		}
		index := value - 1

		if nums[index] < 0 {
			retArr = append(retArr, value)
		}
		nums[index] = -nums[index]
	}
	return retArr
}

func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
