package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1386. Cinema Seat Allocation", func(t *testing.T) {
		reservedSeats := [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}}
		n := 3
		want := 4
		got := solution(n, reservedSeats)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1386. Cinema Seat Allocation2", func(t *testing.T) {
		reservedSeats := [][]int{{2, 1}, {1, 8}, {2, 6}}
		n := 2
		want := 2
		got := solution(n, reservedSeats)
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
*/
func solution2(n int, reservedSeats [][]int) int {

}
