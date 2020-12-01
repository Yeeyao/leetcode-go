package main

import "fmt"

func main() {
	src := []int{9, 8, 7, 1, 2, 3, 4, 5, 6}
	quickSort(src)
	fmt.Println(src)
}

/*
	这里形参和返回值的命名重复了，会报错
*/
//func t(i int) (i int) {
//	defer i++
//	return 1
//}

// 注意这里递归调用 qsHelper
func quickSort(nums []int) {
	qsHelper(nums, 0, len(nums)-1)
}

func qsHelper(nums []int, left, right int) {
	if right > left {
		pivotIndex := left
		pivotIndexN := qs(nums, left, right, pivotIndex)
		qsHelper(nums, left, pivotIndexN-1)
		qsHelper(nums, pivotIndexN+1, right)
	}
}

func qs(nums []int, left, right, pivotIndex int) int {
	pivotVal := nums[pivotIndex]
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	pCount := left
	for i := left; i < right; i++ {
		if nums[i] < pivotVal {
			nums[i], nums[pCount] = nums[pCount], nums[i]
			pCount++
		}
	}
	nums[pCount], nums[right] = nums[right], nums[pCount]
	return pCount
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	1		1
   2 3	   2 4

判断根节点是否 nil,是就直接返回这个
遍历判断，
	如果当前数值大于当前节点数值，
		判断当前节点的右子树是否 nil
			是，就直接插入到新建的右子树，直接返回
			不是则继续向右边遍历
	如果当前数值小于当前节点数值
		判断当前节点左子树是否 nil
			是，就直接插入到新建的左子树，直接返回
			不是则继续向右边遍历
*/
func bstInsert(x int, root *TreeNode) {
	if root == nil {
		root = &TreeNode{Val: x}
	}
	inPos := root
	for inPos != nil {
		if x == inPos.Val {
			return
		} else if x > inPos.Val {
			if inPos.Right == nil {
				inPos.Right = &TreeNode{Val: x}
				return
			}
			inPos = inPos.Right
		} else {
			if inPos.Left == nil {
				inPos.Left = &TreeNode{Val: x}
				return
			}
			inPos = inPos.Left
		}
	}
}

func longestPath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(longestPath(root.Left), longestPath(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bobbleSort(arr []int) {
	arrLen := len(arr)
	sorted := false
	for i := 0; i < arrLen; i++ {
		sorted = true
		for j := 0; j < arrLen-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				sorted = false
			}
		}
		if sorted {
			break
		}
	}

}

// 递归计算左右部分然后调用合并
func mergeSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen < 2 {
		return arr
	}
	mid := arrLen / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)

}

func merge(left, right []int) []int {
	tempArr := make([]int, len(left)+len(right))
	i, j, index := 0, 0, 0
	// 一个死循环，其中一个数组遍历完，直接复制剩下的部分然后退出循环
	// 每次更新完判断是否遍历完
	for {
		if left[i] > right[j] {
			tempArr[index] = right[j]
			j++
			index++
			if j == len(right) {
				copy(tempArr[index:], left[i:])
				break
			}
		} else {
			tempArr[index] = left[i]
			i++
			index++
			if i == len(left) {
				copy(tempArr[index:], right[j:])
				break
			}
		}
	}
	return tempArr
}
