package hash_table

import "testing"

func TestPro(t *testing.T) {
	t.Run("leetcode 961 N-Repeated Element in Size 2N Array", func(t *testing.T) {
		input := []int{8, 2, 3, 3}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
36. Valid Sudoku 给定一个 9*9 的数组以及部分位置上面的数字 1-9，判断这个二维数组是否是一个合法的数独数组
直接创建 9 * 3 个 哈希表，然后遍历每个元素，将每个元素保存到对应的哈希表中，保存前判断每个位置是否已经有元素了，如果有则直接返回 false
	这里判断使用 map[int]int 然后每个元素先累加，然后判断是否已经存在
使用二维数组简化处理
遍历完所有元素，直接返回 true。所以有更好的方法吗？为什么这题是 medium
这里的一个 trick 是，元素所在的 box 是 (i/3)*3+j/3
*/
func solution(board [][]byte) bool {
	var row, col, box [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				pos := board[i][j] - '1'
				boxIndex := (i/3)*3 + j/3
				if row[i][pos] == 1 {
					return false
				} else {
					row[i][pos]++
				}
				if col[j][pos] == 1 {
					return false
				} else {
					col[j][pos]++
				}
				if box[boxIndex][pos] == 1 {
					return false
				} else {
					box[boxIndex][pos]++
				}
			}
		}
	}
	return true
}
