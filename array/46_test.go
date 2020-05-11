package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("46. Permutations", func(t *testing.T) {
		input := []int{1, 2, 3}
		want := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
		got := solution(input)
		if IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	使用 backtracking 对比 78，90 这里需要使用所有元素，然后进行排列
	这里利用循环遍历的顺序以及递归调用的顺序进行元素顺序的处理
*/
func solution(nums []int) [][]int {
	retArr := make([][]int, 0)
	numsLen := len(nums)
	solutionHelper(nums, &retArr, []int{}, numsLen)
	return retArr
}

func solutionHelper(nums []int, retArr *[][]int, solArr []int, numsLen int) {
	// 数量相等才保存
	saLen := len(solArr)
	if saLen == numsLen {
		*retArr = append(*retArr, append([]int{}, solArr...))
	}
	// 循环递归处理 每次都从头开始判断
	for i := 0; i < numsLen; i++ {
		// 判断元素是否已经存在了，存在则取下一个元素存放
		if isInSlice(solArr, nums[i]) {
			continue
		}
		// 添加元素的代码
		solArr = append(solArr, nums[i])
		solutionHelper(nums, retArr, solArr, numsLen)
		solArr = solArr[:saLen]
	}
}

// 判断元素是否存在于 slice 中
func isInSlice(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

/*
	另一种改进的 recursive 算法
	当遍历的开始位置已经最后了就保存
	对于 nums[start ... end]，前面的 nums[0 ... start - 1]是确定好的
	接下来，就是将所有元素都放到当前的 nums[start] 来得到所有迭代的情况
	这种思路好棒啊
*/
func solution2(nums []int) [][]int {
	retArr := make([][]int, 0)
	numsLen := len(nums)
	solutionHelper2(nums, &retArr, numsLen, 0)
	return retArr
}

func solutionHelper2(nums []int, retArr *[][]int, numsLen, start int) {
	// 遍历完了才保存
	if start == numsLen {
		*retArr = append(*retArr, append([]int{}, nums...))
	}
	for i := start; i < numsLen; i++ {
		// 进行一次交换后递归处理
		nums[start], nums[i] = nums[i], nums[start]
		solutionHelper2(nums, retArr, numsLen, start+1)
		// 上个交换后要还原来处理下一次遍历
		nums[i], nums[start] = nums[start], nums[i]
	}
}

func IntSliceEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
