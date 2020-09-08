package string

/*
79. Word Search
Given a 2D board and a word, find if the word exists in the grid.
The word can be constructed from letters of sequentially adjacent cell,
where "adjacent" cells are those horizontally or vertically neighboring.
The same letter cell may not be used more than once.
给定二维数组的单词数组，在其中找到是否可以组成给定的单词，字母通过上下或者左右连接，每个字母只能使用一次
回溯加上剪枝，这里需要每个字母都作为开始来处理处理完需要改变
*/
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if helper(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func helper(board [][]byte, i, j, k int, word string) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 || board[i][j] != word[k] {
		return false
	}
	temp := board[i][j]
	board[i][j] = '0'
	res := helper(board, i-1, j, k+1, word) || helper(board, i, j-1, k+1, word) ||
		helper(board, i+1, j, k+1, word) || helper(board, i, j+1, k+1, word)
	board[i][j] = temp
	return res

}
