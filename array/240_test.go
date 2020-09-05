package array

/*
240. Search a 2D Matrix II
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
    Integers in each row are sorted in ascending from left to right.
    Integers in each column are sorted in ascending from top to bottom.
一行中元素从左到右递增，一列中元素从上到下递增

对比 74 是行首元素大于上一行的行尾元素 这里没有这个特性
每行进行二分查找，如果多行的开头和结尾都超过 target 效率就高，然后如果一行的数据很多，内存可能要频繁刷新
*/

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 || len(matrix[0]) == 0 {
		return false
	}
	for _, m := range matrix {
		if binarySearch(m, target) {
			return true
		}
	}
	return false
}

func binarySearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left] == target
}

/*
	直接顺序遍历居然是最快的，就离谱
*/
func searchMatrix2(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 || len(matrix[0]) == 0 {
		return false
	}
	col := len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}
