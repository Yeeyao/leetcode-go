package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("840. Magic Squares In Grid", func(t *testing.T) {
		input := [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	根据观察，中间的元素必须是 5，偶数必须在角上，奇数需要在边
	同时，需要的顺序 43816729，顺时针或者逆时针
	以及，和需要是 15
*/
func solution(input [][]int) int {
	inputLen := len(input)
	if inputLen < 3 || len(input[0]) < 3 {
		return 0
	}
	sum := 0
	for i := 0; i < inputLen-2; i++ {
		for j := 0; j < len(input[0])-2; j++ {
			if isMagic(i, j, input) {
				sum++
			}
		}
	}
	return sum
}

func isMagic(i, j int, input [][]int) bool {
	// 中间是否是 5
	if input[i+1][j+1] != 5 {
		return false
	}
	// 角落的元素需要是 偶数
	if input[i][j]%2 != 0 || input[i+2][j]%2 != 0 ||
		input[i][j+2]%2 != 0 || input[i+2][j+2]%2 != 0 {
		return false
	}
	// 中间的元素需要是 奇数
	if input[i+1][j]%2 == 0 || input[i][j]+1%2 == 0 ||
		input[i+1][j+2]%2 == 0 || input[i+2][j+1]%2 == 0 {
		return false
	}
	// 和需要是 15 只需要检查上方的外围部分
	if input[i][j]+input[i][j+1]+input[i][j+2] != 15 ||
		input[i+2][j]+input[i+2][j+1]+input[i+2][j+2] != 15 ||
		input[i][j]+input[i+1][j]+input[i+2][j] != 15 {
		return false
	}
	return true
}
