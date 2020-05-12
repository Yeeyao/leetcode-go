package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run(" 289. Game of Life ", func(t *testing.T) {
		input := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
		want := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}
		solution(input)
		if !IntSliceEqual(input, want) {
			t.Errorf("got: %v, want: %v", input, want)
		}
	})
}

/*
	这里每个判断倒还好 难点在于不能一直变化下去，即一个位置变化了，
	不能直接影响其他位置
	所以大概需要复制一份数组，或者其他方法使用的位存储原来的数组信息
	坐标的关系，自己画个坐标系看看
*/
func solution(board [][]int) {
	row := len(board)
	col := len(board[0])
	src := make([][]int, row)
	for key, value := range board {
		src[key] = make([]int, col)
		copy(src[key], value)
	}
	// 需要判断周围 8 个格子
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			nearLives := 0
			// alive numbers
			if i > 0 {
				// 左上角
				if j > 0 {
					if src[i-1][j-1] == 1 {
						nearLives++
					}
				}
				// 正上方
				if src[i-1][j] == 1 {
					nearLives++
				}
				// 右上角
				if j < col-1 {
					if src[i-1][j+1] == 1 {
						nearLives++
					}
				}
			}
			// 左边
			if j > 0 {
				if src[i][j-1] == 1 {
					nearLives++
				}
			}
			// 右边
			if j < col-1 {
				if src[i][j+1] == 1 {
					nearLives++
				}
			}
			if i < row-1 {
				// 右下角
				if j < col-1 {
					if src[i+1][j+1] == 1 {
						nearLives++
					}
				}
				// 正下方
				if src[i+1][j] == 1 {
					nearLives++
				}
				// 左下方
				if j > 0 {
					if src[i+1][j-1] == 1 {
						nearLives++
					}
				}
			}
			// 当前元素 alive
			if src[i][j] == 1 {
				// condition 1
				if nearLives < 2 {
					board[i][j] = 0
				}
				// condition 3
				if nearLives > 3 {
					board[i][j] = 0
				}
			}
			// 当前元素 die
			if src[i][j] == 0 {
				if nearLives == 3 {
					board[i][j] = 1
				}
			}
		}
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
