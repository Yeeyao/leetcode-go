package array

import (
	"fmt"
	"testing"
)

/*
	找到所有单数的行，然后减去偶数的列这样处理
*/
func TestPro(t *testing.T) {
	t.Run("1252. Cells with Odd Values in a Matrix", func(t *testing.T) {
		indice := [][]int{{0, 1}, {1, 1}}
		got := solution(2, 3, indice)
		want := 6
		fmt.Println(got)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1252. Cells with Odd Values in a Matrix2", func(t *testing.T) {
		indice := [][]int{{1, 1}, {0, 0}}
		got := solution(2, 2, indice)
		want := 0
		fmt.Println(got)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

// method one slower one
func solution(n, m int, indices [][]int) int {
	rowCounter := make(map[int]int)
	colCounter := make(map[int]int)
	rowOddCounter := 0
	colOddCounter := 0
	for _, v := range indices {
		rowCounter[v[0]]++
		colCounter[v[1]]++
	}

	for _, v := range rowCounter {
		if v%2 != 0 {
			rowOddCounter++
		}
	}

	for _, v := range colCounter {
		if v%2 != 0 {
			colOddCounter++
		}
	}
	return rowOddCounter*m + colOddCounter*n - 2*rowOddCounter*colOddCounter
}

// method two faster
func solution(n, m int, indices [][]int) int {
	rowCounter := make(map[int]int)
	colCounter := make(map[int]int)
	oddNum := 0
	for _, v := range indices {
		// 针对行
		rowCounter[v[0]]++
		// 针对列
		colCounter[v[1]]++
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			num := rowCounter[i] + colCounter[j]
			if num%2 != 0 {
				oddNum++
			}
		}
	}
	return oddNum
}
