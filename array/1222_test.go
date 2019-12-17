package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1222. Queens That Can Attack the King", func(t *testing.T) {
		input := [][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}}
		king := []int{0, 0}
		want := [][]int{{0, 1}, {1, 0}, {3, 3}}
		got := solution(input, king)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	找到可以吃 king 的 queens 其中后者可以上下以及斜对角线移动，
	同时，路径上有其他的 queens 会阻碍

	8 * 8 棋盘 其实就八个方向。每个方向都保存最近的就行。
	使用一个 slice 来保存每个位置的当前最近的 每遍历一个 queen 就更新最近的
	这里保存最近的方式是，直接遍历每个方向，从靠近 king 到远离的顺序去查找坐标点是否出现在
	queens 中。
*/
func solution(input [][]int, king []int) [][]int {
	var retArr [][]int
	directArr := []int{-1, 0, 1}
	xArr := make([]int, 8)
	for i, _ := range xArr {
		xArr[i] = i + 1
	}
	for _, v := range directArr {
		for _, w := range directArr {
			// 八个方向 这里只是将下面的合并在一起了
			// 同时，因为 king 本身不会出现在 queens 中，所以 0,0 的坐标也不会影响判断
			for _, k := range xArr {
				x, y := king[0]+v*k, king[1]+w*k
				addArr := []int{x, y}
				if inArr(input, addArr) {
					retArr = append(retArr, addArr)
					break
				}
			}
		}
	}
	return retArr
}

func inArr(arr [][]int, dest []int) bool {
	for _, v := range arr {
		if dest[0] == v[0] && dest[1] == v[1] {
			return true
		}
	}
	return false
}

// brute force
func queensAttacktheKing(queens [][]int, king []int) [][]int {
	out := make([][]int, 8)
	// element counter
	index := 0
	//left
	for i := king[0] - 1; i >= 0; i-- {
		if queenExists(i, king[1], queens) {
			out[index] = []int{i, king[1]}
			index++
			break
		}
	}
	//right
	for i := king[0] + 1; i < 8; i++ {
		if queenExists(i, king[1], queens) {
			out[index] = []int{i, king[1]}
			index++
			break
		}
	}
	//top
	for i := king[1] - 1; i >= 0; i-- {
		if queenExists(king[0], i, queens) {
			out[index] = []int{king[0], i}
			index++
			break
		}
	}
	//bottom
	for i := king[1] + 1; i < 8; i++ {
		if queenExists(king[0], i, queens) {
			out[index] = []int{king[0], i}
			index++
			break
		}
	}
	//left-top
	for i, j := king[0]-1, king[1]-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if queenExists(i, j, queens) {
			out[index] = []int{i, j}
			index++
			break
		}
	}
	//left-bottom
	for i, j := king[0]-1, king[1]+1; i >= 0 && j < 8; i, j = i-1, j+1 {
		if queenExists(i, j, queens) {
			out[index] = []int{i, j}
			index++
			break
		}
	}
	//right-bottom
	for i, j := king[0]+1, king[1]+1; i < 8 && j < 8; i, j = i+1, j+1 {
		if queenExists(i, j, queens) {
			out[index] = []int{i, j}
			index++
			break
		}
	}
	//right-top
	for i, j := king[0]+1, king[1]-1; i < 8 && j >= 0; i, j = i+1, j-1 {
		if queenExists(i, j, queens) {
			out[index] = []int{i, j}
			index++
			break
		}
	}
	return out[:index]
}

func queenExists(i, j int, queens [][]int) bool {
	for a := 0; a < len(queens); a++ {
		if queens[a][0] == i && queens[a][1] == j {
			return true
		}
	}
	return false
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
