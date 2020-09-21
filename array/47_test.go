package array

import (
	"reflect"
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run(" 47. Permutations II ", func(t *testing.T) {
		input := []int{1, 1, 2}
		want := [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}
		got := solution(input)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	[ref](https://leetcode-cn.com/problems/permutations-ii/solution/quan-pai-lie-ii-by-leetcode-solution/)

	这里类似 46 对于 idx 的位置，需要保证第 idx 的位置，重复的数字只会出现一次
	因此对原数组进行排序，保证相同的数字都相邻，然后每次填入的数一定是这个数所在的重复数集合
	中从左向右的第一个未被填入的数字。这里 visited 保存的是索引
	if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
		continue
	}
	同时注意这里每次都是每个元素都遍历然后判断是否已经被使用，因为存在 visited 数组
*/

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	numsLen := len(nums)
	perm := []int{}
	visited := make([]bool, numsLen)
	var backtrack func(int)
	backtrack = func(index int) {
		if index == numsLen {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			// 当前索引的元素已经使用了或者
			// 当前元素等于上一个元素同时上一个元素没有被使用
			// 这里需要上一个元素被使用了才能使用当前元素
			if visited[i] || i > 0 && !visited[i-1] && v == nums[i-1] {
				continue
			}
			// 递归前以及恢复处理
			perm = append(perm, v)
			visited[i] = true
			backtrack(index + 1)
			visited[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
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
