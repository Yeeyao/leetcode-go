package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run(" 47. Permutations II ", func(t *testing.T) {
		input := []int{1, 1, 2}
		want := [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}
		got := solution(input)
		if IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	对比 46 这里需要去重处理

	使用一个布尔数组来保存总结果中当前元素是否已经使用的信息，如果被使用了，就跳过该元素
	注意这里的布尔数组也是每次递归传递下去，类似中间结果数组
	同时，如果当前元素与上一个元素相等且上一个元素没有被使用，也跳过当前元素
	每次递归前设置当前元素为使用，递归后改回来
	当前元素和前面的元素相等，仅当前一个元素已经被使用了，我们才能使用当前元素

	如果中间数组的长度等于给定数组，就需要保存结果
	循环从 0 开始处理
*/
func solution(nums []int) [][]int {
	retArr := make([][]int, 0)
	numsLen := len(nums)
	if numsLen == 0 {
		return retArr
	}
	// 排序是因为代码需要判断前后元素是否相等
	sort.Ints(nums)
	usedArr := make([]bool, numsLen)
	solutionHelper(nums, &retArr, usedArr, []int{}, numsLen)
	return retArr
}

func solutionHelper(nums []int, retArr *[][]int, usedArr []bool, solArr []int, numsLen int) {
	// 数量相等才保存
	saLen := len(solArr)
	if saLen == numsLen {
		*retArr = append(*retArr, append([]int{}, solArr...))
	}
	// 循环递归处理 每次都从头开始判断
	for i := 0; i < numsLen; i++ {
		// 判断当前元素是否已经被使用了
		if usedArr[i] {
			continue
		}
		// 当前元素和前面的元素相等，仅当前一个元素已经被使用了，我们才能使用当前元素
		if i > 0 && nums[i] == nums[i-1] && !usedArr[i-1] {
			continue
		}
		// 添加元素的代码
		usedArr[i] = true
		solArr = append(solArr, nums[i])
		solutionHelper(nums, retArr, usedArr, solArr, numsLen)
		usedArr[i] = false
		solArr = solArr[:saLen]
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
