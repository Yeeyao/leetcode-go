package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1386. Cinema Seat Allocation", func(t *testing.T) {
		reservedSeats := [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}}
		n := 3
		want := 4
		got := solution2(n, reservedSeats)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1386. Cinema Seat Allocation2", func(t *testing.T) {
		reservedSeats := [][]int{{2, 1}, {1, 8}, {2, 6}}
		n := 2
		want := 2
		got := solution2(n, reservedSeats)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	TLE
*/
func solution(n int, reservedSeats [][]int) int {
	count := 0
	reserved := make(map[int][]int)
	for _, v := range reservedSeats {
		reserved[v[0]] = append(reserved[v[0]], v[1])
	}
	for i := 1; i <= n; i++ {
		// 当前行已经占用的座位索引
		leftLastTwoEmpty := !isMember(2, 3, reserved[i])
		midFirstTwoEmpty := !isMember(4, 5, reserved[i])
		midLastTwoEmpty := !isMember(6, 7, reserved[i])
		rightFirstTwoEmpty := !isMember(8, 9, reserved[i])
		if leftLastTwoEmpty && midFirstTwoEmpty {
			count++
		}
		if midLastTwoEmpty && rightFirstTwoEmpty {
			count++
		}
		if midFirstTwoEmpty && midLastTwoEmpty && count == 0 {
			count++
		}
	}
	return count
}

func isMember(a, b int, arr []int) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if a == v {
			return true
		}
		if b == v {
			return true
		}
	}
	return false
}

/*
	use bit vector
	2,3,4,5 可以使用这个形式表示 (1 << 2) | (1 << 3) | (1 << 4) | (1 << 5) = 60
	6,7,8,9 以及 4,5,6,7 同样可以使用这样的形式
	因此，对于保留的数组的每一行，直接用 0 和 每一列做位或保存，用一个数字保存下
	每一列的保留座位信息
	[参考](https://leetcode.com/problems/cinema-seat-allocation/discuss/546451/Java-Straightforward-solution-(bitwise))
*/
func solution2(n int, reservedSeats [][]int) int {
	sum := 0
	reserved := make(map[int]int)
	for _, v := range reservedSeats {
		key, value := v[0], v[1]
		// map 中存在 需要将 1 左移然后做位或
		if old, ok := reserved[key]; ok {
			reserved[key] = old | 1<<value
		} else {
			// 0 做位或等于本身
			reserved[key] = 1 << value
		}
	}
	for _, v := range reserved {
		count := 0
		// 当前行已经占用的座位索引 先检查两边再检查中间
		if v&60 == 0 {
			count++
		}
		if v&960 == 0 {
			count++
		}
		if v&240 == 0 && count == 0 {
			count = 1
		}
		sum += count
	}
	// n 小于的部分就直接每行 2 个座位
	return sum + 2*(n-len(reserved))
}
