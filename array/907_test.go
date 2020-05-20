package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1386. Cinema Seat Allocation", func(t *testing.T) {
		A := []int{3, 1, 2, 4}
		want := 17
		got := solution(A)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	找所有子数组的最小值的和
	类似 78 不对，这里要求所有的子数组的需要需要是连续的
	TLE
*/
func solution(A []int) int {
	sum := 0
	ALen := len(A)
	solutionHelper(A, &sum, []int{}, 0, ALen)
	return sum % (1000000000 + 7)
}

func solutionHelper(A []int, sum *int, tempArr []int, start, ALen int) {
	tempArrLen := len(tempArr)
	if tempArrLen > 0 {
		tempMin := 30001
		for _, v := range tempArr {
			if v < tempMin {
				tempMin = v
			}
		}
		*sum += tempMin
	}
	for i := start; i < ALen; i++ {
		if tempArrLen == 0 {
			tempArr = append(tempArr, A[i])
			solutionHelper(A, sum, tempArr, i+1, ALen)
			tempArr = tempArr[:tempArrLen]
		} else {
			if i+1 <= ALen {
				tempArr = append(tempArr, A[i])
				solutionHelper(A, sum, tempArr, i+1, ALen)
			}
			break
		}
	}
}

/*

 */
func solution2(A []int) int {

}
