package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {

	// Given a m * n matrix mat of integers, sort it diagonally in ascending order from the top-left to the bottom-right then return the sorted array.

	t.Run("1329. Sort the Matrix Diagonally", func(t *testing.T) {
		mat := [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}
		want := [][]int{{1, 1, 1, 1}, {1, 2, 2, 2}, {1, 2, 3, 3}}
		got := solution(mat)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	// t.Run("1329. Sort the Matrix Diagonally2", func(t *testing.T) {
	// 	mat := [][]int{{3, 3, 1, 1}}
	// 	want := [][]int{{1, 1, 3, 3}}
	// 	got := solution(mat)
	// 	if !IntSliceEqual(got, want) {
	// 		t.Errorf("got: %v, want: %v", got, want)
	// 	}
	// })
}

// 想到的思路是从右下角开始放置数据
// 或者换个方向，先将第一行和第一列填满，然后继续填第二行，第二列这样

// 对角线数量 m * n 如果 m == n 则是 m，否则是较小者 + 较大者 / 2 （这里向上取整）
// 按照 hint 每个结构存放一条对角线
// 对角线的最大长度是 m n 中的较小者，其中两边的数量是从 1 开始向最大长度递增，如果
// 需要看对角线总数是奇数还是偶数来判断

// 元素升序排序之后，遍历每条对角线，一次加一个元素直到对角线元素满了
// 其中第几条对角线，直接通过下标确定。
// 这里可以直接 [][]int 来存放对角线的数据就行了，但是 struct 直接记录下该对角线的长度，不需要每次都计算
// 放置元素的时候，直接用当前的轮数和对角线的数量比较，如果轮数大于数量，则表示该对角线不需要再存放了

/*
	想得限制条件太多了，不需要那么麻烦，题目只要求对角线上元素是有序就行了
	所以，只需要将元素放到一条对角线上，然后每条对角线排序，最后放回去就行
	对角线之间的元素大小顺序不做要求
	所以题目的 hints 误导人
*/

// type diagonal struct {
// 	elemNum int
// 	elem    []int
// }

// func solution(mat [][]int) [][]int {

// 	row := len(mat)
// 	col := len(mat[0])

// 	// store the element and sort
// 	allElem := make([]int, row*col)
// 	counter := 0

// 	for _, i := range mat {
// 		for _, j := range i {
// 			allElem[counter] = j
// 			counter++
// 		}
// 	}
// 	sort.Ints(allElem)

// 	counter = 0

// 	// 根据 m * n 然后初始化 struct slice
// 	structSlice, diagonalNum, less := createStructSlice(row, col)
// 	// 按照顺序将值存放到 struct
// 	for i := 0; i < less; i++ {
// 		for j := 0; j < diagonalNum; j++ {
// 			// 这里针对一个对角线，不断插入元素
// 			if structSlice[j].elemNum >= i+1 {
// 				structSlice[j].elem[i] = allElem[counter]
// 				counter++
// 			}
// 		}
// 	}

// 	var retArr [][]int

// 	for i := 0; i < row; i++ {
// 		v := make([]int, col)
// 		retArr = append(retArr, v)
// 	}

// 	fmt.Println(diagonalNum)
// 	// 如何提取 找映射关系 针对一行或者一列特殊情况处理
// 	for i := 0; i < diagonalNum; i++ {
// 		var x, y int
// 		beginX := less - i - 1
// 		if beginX < 0 {
// 			x = 0
// 			y = -beginX
// 		} else {
// 			x = beginX
// 			y = 0
// 		}
// 		for _, e := range structSlice[i].elem {
// 			retArr[x][y] = e
// 			x++
// 			y++
// 		}
// 	}

// 	return retArr
// }

// func createStructSlice(row, col int) (s []diagonal, d, l int) {
// 	var diagonalNum int
// 	var less int
// 	var lessNum int
// 	if row == col {
// 		diagonalNum = row
// 	} else if row > col {
// 		less = col
// 		if col == 1 {
// 			diagonalNum = row
// 		} else {
// 			if row/2 == 0 {
// 				lessNum = row / 2
// 			} else {
// 				lessNum = row/2 + 1
// 			}
// 			diagonalNum = col + lessNum
// 		}
// 	} else if row < col {
// 		less = row
// 		if row == 1 {
// 			diagonalNum = col
// 		} else {
// 			if col/2 == 0 {
// 				lessNum = col / 2
// 			} else {
// 				lessNum = col/2 + 1
// 			}
// 			diagonalNum = row + lessNum
// 		}
// 	}

// 	structSlice := make([]diagonal, diagonalNum)
// 	// 每条对角线处理
// 	for i := 0; i < diagonalNum; i++ {
// 		j := i + 1
// 		if j <= less {
// 			structSlice[i].elemNum = j
// 			structSlice[i].elem = make([]int, j)
// 		} else {
// 			j := diagonalNum - i
// 			structSlice[i].elemNum = j
// 			structSlice[i].elem = make([]int, j)
// 		}
// 	}

// 	return structSlice, diagonalNum, less
// }

// 这里如何从输入中将所有的
func solution(mat [][]int) [][]int {
	if len(mat) == 0 {
		return mat
	}

	m := len(mat)
	n := len(mat[0])

	tmp := make([][]int, m*n)

	// 这里是怎么存放的
	// 这个 diff
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diff := j - i
			// 这里存放位置需要前后一致
			if diff < 0 {
				diff = diff * -n
			}
			tmp[diff] = append(tmp[diff], mat[i][j])
		}
	}

	// 桶内排序
	for i := range tmp {
		sort.Slice(tmp[i], func(a, b int) bool {
			return tmp[i][a] < tmp[i][b]
		})
	}

	// 放回去
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diff := j - i
			if diff < 0 {
				diff = diff * -n
			}

			// 放回去然后取剩余的
			mat[i][j] = tmp[diff][0]
			tmp[diff] = tmp[diff][1:]
		}
	}
	return mat
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
