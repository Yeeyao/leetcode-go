package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("65  矩阵中的路径", func(t *testing.T) {
		board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
		word := "ABCCED"
		get := solution(board, word)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。
如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]]

这个类似走迷宫吧
但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
用回溯法和剪枝，从第一行的每个字符开始遍历
但是可以从矩阵的任何一个格子开始

中止条件：
	返回 false 行或者列索引越界 当前矩阵元素与目标字符不同 当前矩阵元素已经访问过
	返回 true 字符 word 全部匹配 k = len(word) - 1
递推
	保存 board[i][j] 到临时元素并标记当前元素为'/' 表示已经访问
	递归调用四个方向
	还原
返回
	最后返回 res
*/
func solution(board [][]byte, word string) bool {
	// 这里每个元素都需要作为开始元素一次，只要有一个满足就直接返回
	for i, row := range board {
		for j, _ := range row {
			if dfs(board, word, i, j, 0) == true {
				return true
			}
		}
	}
	return false
}

/*
	dfs 遍历
*/
func dfs(board [][]byte, word string, i, j, k int) bool {
	// 全部匹配
	if k == len(word) {
		return true
	}
	// 越界检查以及字符检查
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] {
		return false
	}
	temp := board[i][j]
	// 标记当前元素已经访问过
	board[i][j] = '0'
	// 四个方向递归处理
	res := dfs(board, word, i-1, j, k+1) || dfs(board, word, i+1, j, k+1) ||
		dfs(board, word, i, j-1, k+1) || dfs(board, word, i, j+1, k+1)
	board[i][j] = temp
	return res
}
