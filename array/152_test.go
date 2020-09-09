package array

/*
152. Maximum Product Subarray
Given an integer array nums,
find the contiguous subarray within an array (containing at least one number) which has the largest product.

给定整型数组，找到数组连续子数组的最大乘积
这里还有正负号需要考虑 使用两个变量，一个是当前最大负数，一个是当前最大正数
因为是乘法，然后整型，所以只要不是乘以 0 一直乘是最大的
[ref](https://leetcode.com/problems/maximum-product-subarray/discuss/48230/Possibly-simplest-solution-with-O(n)-time-complexity)
*/

// 无敌了
/*
	初始化 res 为第一个元素，然后 iMax, iMin 分别维护当前的最大正数乘积和最小负数乘积
	从第二个元素开始遍历数组
		如果当前元素小于 0，将 iMax, iMin 互换，
		iMax = max(nums[i], iMax*nums[i])
		iMin = min(nums[i], iMin*nums[i])
		最后用 iMax 以及 res 较大值更新 res
*/
func maxProduct(nums []int) int {
	numsLen := len(nums)
	res := nums[0]
	for i, iMax, iMin := 1, res, res; i < numsLen; i++ {
		if nums[i] < 0 {
			iMax, iMin = iMin, iMax
		}
		iMax = max(nums[i], iMax*nums[i])
		iMin = min(nums[i], iMin*nums[i])
		res = max(res, iMax)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProduct2(nums []int) int {
	numsLen := len(nums)
	mp, mn := -1, 1
	tmp, tmn := 1, 1
	for i := 0; i < numsLen; i++ {
		if nums[i] >= 0 {
			tmp *= nums[i]
			if tmp > mp {
				mp = tmp
			}
			if tmp == 0 {
				tmp = 1
			}
		} else {
			tmn *= nums[i]
			if tmn < mn {
				mn = tmn
			}
			if tmn == 0 {
				tmn = 1
			}
		}
	}
	return mp
}
