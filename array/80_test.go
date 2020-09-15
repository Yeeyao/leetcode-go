package array

/*
Given a sorted array nums, remove the duplicates in-place
such that duplicates appeared at most twice and return the new length.
Do not allocate extra space for another array,
you must do this by modifying the input array in-place with O(1) extra memory.
给定的已排序元素，每个元素最多出现两次，需要移除超过该次数的元素，不使用额外空间

*/

/*
	更简单的方法 i 记录最终数组的元素数量
	遍历每个元素
		如果 i < 2 或者 n > nums[i-2](当前元素大于上上个元素)
			nums[i] = n i++
	最终返回 i
*/

func removeDuplicates(nums []int) int {
	i := 0
	for _, n := range nums {
		if i < 2 || n > nums[i-2] {
			nums[i] = n
			i++
		}
	}
	return i
}

/*
使用一个额外布尔值保存每个当前元素的出现次数
如果当前元素和上一个元素相等
	布尔值是否已经为 true，是则 count++
	不是则元素直接复制
如果不相等，直接将布尔值设置为 false 并继续赋值
*/
func removeDuplicates2(nums []int) int {
	numsLen := len(nums)
	count := 0
	isTwice := false
	for i := 0; i < numsLen; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			if isTwice == false {
				nums[i-count] = nums[i]
				isTwice = true
			} else {
				count++
			}
		} else {
			isTwice = false
			nums[i-count] = nums[i]
		}
	}
	return numsLen - count
}
