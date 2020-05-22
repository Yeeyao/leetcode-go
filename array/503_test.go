package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("503. Next Greater Element II", func(t *testing.T) {
		nums := []int{1, 2, 1}
		want := []int{2, -1, 2}
		got := solution(nums)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("503. Next Greater Element II2", func(t *testing.T) {
		nums := []int{1, 2, 1}
		want := []int{2, -1, 2}
		got := solution2(nums)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	相比 496 数组变成了循环数组 使用两个栈？
	原来是只有一个单调递减栈。先保存前面的部分，再保存后面的部分
	保存前面的部分，就遍历一次，向前找第一个比当前元素大的。这里使用
	保存后面的部分同 496

	同时，注意数组的元素可重复
	都不是，只是简单的 loop twice 就可以。emm，子集验证一下吧
	这里可以知道，遍历一遍的时候，找不到下一个最大元素的元素都会保存到栈
	因此，在栈中存在这些元素的情况下，再遍历一次，就实现了循环的遍历
*/
func solution(nums []int) []int {
	numsLen := len(nums)
	retArr := make([]int, numsLen)
	for i, _ := range retArr {
		retArr[i] = -1
	}
	// 栈中保存的是元素位置 注意这里栈的空间
	// 主要注意栈顶的变化，因为循环后面将栈顶位置 + 1
	// 所以循环前访问栈顶元素的位置应该 - 1
	st := make([]int, numsLen*2)
	stTop := 0
	for i := 0; i < 2*numsLen; i++ {
		// 非空且当前元素比栈顶元素大，当前元素就是栈顶元素的后面第一个比它大的元素
		// 因为栈顶元素是先遍历到的 这里需要注意使用的是 for 只要元素不断大于栈顶
		// 就不断更新
		for stTop > 0 && nums[i%numsLen] > nums[st[stTop-1]] {
			// 这里要保存元素的值
			retArr[st[stTop-1]] = nums[i%numsLen]
			stTop--
		}
		st[stTop] = i % numsLen
		stTop++
	}
	return retArr
}

/*
	这里跟上面类似，只是将循环分成了两个来处理
*/
func solution2(nums []int) []int {
	numsLen := len(nums)
	retArr := make([]int, numsLen)
	// 初始化都是 -1
	for i, _ := range retArr {
		retArr[i] = -1
	}
	// 栈中保存的是元素位置 注意这里栈的空间
	// 主要注意栈顶的变化，因为循环后面将栈顶位置 + 1
	// 所以循环前访问栈顶元素的位置应该 - 1
	st := make([]int, numsLen)
	stTop := 0
	for i := 0; i < numsLen; i++ {
		// 非空且当前元素比栈顶元素大，当前元素就是栈顶元素的后面第一个比它大的元素
		// 因为栈顶元素是先遍历到的 这里需要注意使用的是 for 只要元素不断大于栈顶
		// 就不断更新
		for stTop > 0 && nums[i] > nums[st[stTop-1]] {
			// 这里要保存元素的值
			retArr[st[stTop-1]] = nums[i]
			stTop--
		}
		st[stTop] = i
		stTop++
	}
	// 注意这里第二趟就不需要将新的元素入栈了，因为第一趟需要入栈的已经入栈了
	for i := 0; i < numsLen; i++ {
		// 非空且当前元素比栈顶元素大，当前元素就是栈顶元素的后面第一个比它大的元素
		// 因为栈顶元素是先遍历到的 这里需要注意使用的是 for 只要元素不断大于栈顶
		// 就不断更新
		for stTop > 0 && nums[i] > nums[st[stTop-1]] {
			// 这里要保存元素的值
			retArr[st[stTop-1]] = nums[i]
			stTop--
		}
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
