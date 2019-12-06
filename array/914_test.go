package array

import (
	"fmt"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("914. X of a Kind in a Deck of Cards", func(t *testing.T) {
		input := []int{1, 2}
		want := false
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("914. X of a Kind in a Deck of Cards2", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 4, 3, 2, 1}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("914. X of a Kind in a Deck of Cards3", func(t *testing.T) {
		input := []int{1, 1, 1, 2, 2, 2, 3, 3}
		want := false
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("914. X of a Kind in a Deck of Cards4", func(t *testing.T) {
		input := []int{1, 1}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("914. X of a Kind in a Deck of Cards5", func(t *testing.T) {
		input := []int{1, 1, 2, 2, 2, 2}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	使用另一个 map 来统计数量就好了
	全部统计完然后比较数量
*/
func solution(input []int) bool {
	countMap := make(map[int]int)
	for _, v := range input {
		countMap[v]++
	}
	divisor := 0
	for _, v := range countMap {
		divisor = gcd(v, divisor)
	}
	return divisor > 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
