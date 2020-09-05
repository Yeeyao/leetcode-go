package array

/*
74. Search a 2D Matrix
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
    Integers in each row are sorted from left to right.
    The first integer of each row is greater than the last integer of the previous row.

行中从左到右递增，列中从上到下递增
二维数组，每行开头第一个元素都大于上一个的最后一个元素
从第一行最后一个元素开始判断，如果 target 大于元素，行数递增
如果小于就将列号递减，等于直接返回
*/

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	col := len(matrix[0])
	i, j := 0, col-1
	for i < row && j >= 0 {
		if target == matrix[i][j] {
			return true
		}
		if target > matrix[i][j] {
			i++
		} else {
			j--
		}
	}
	return false
}
