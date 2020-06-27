package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("23 二叉搜索树的后序遍历序列", func(t *testing.T) {
		input := []int{1, 6, 3, 2, 5}
		get := solution(input)
		want := []int{3, 9, 20, 15, 7}
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("23 二叉搜索树的后序遍历序列2", func(t *testing.T) {
		input := []int{1, 3, 2, 6, 5}
		get := solution(input)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同
	直接利用数组判断重构的二叉搜索树是否满足 左右根
	[ref](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/solution/mian-shi-ti-33-er-cha-sou-suo-shu-de-hou-xu-bian-6/)

	主要的问题是找到根节点和左右子树的关系，一开始，根节点是最后的节点

	递归判断
	i >= j 子树节点为 1，直接返回 true
	1.遍历 [i, j] 区间的元素，找到第一个大于根节点的元素，索引为 m，
	其中左子树 [i, m-1]，右子树 [m, j - 1] 根 j
	2.大小判断，[i, m-1] 区间节点要小于 p[j]; [m, j-1] 节点要大于 p[j]
*/
func solution(arr []int) bool {
	arrLen := len(arr)
	return solutionHelper(arr, 0, arrLen-1)
}

func solutionHelper(arr []int, i, j int) bool {
	// 两个索引交叉就表示这里递归下去的子树满足条件
	if i >= j {
		return true
	}
	// 找到第一个大于根节点的元素，这个元素是右子树的根节点 m
	// 其中该节点的左边是左子树的元素
	p := i
	for arr[p] < arr[j] {
		p++
	}
	m := p
	// 遍历当前根节点右子树的元素
	for arr[p] > arr[j] {
		p++
	}
	// 如果第一个条件成立，就表示当前的根节点左右子树的元素都满足
	return (p == j) && solutionHelper(arr, i, m-1) && solutionHelper(arr, m, j-1)
}
