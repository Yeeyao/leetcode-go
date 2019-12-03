package main

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("sqrt", func(t *testing.T) {
		input := 2.0
		precise := 0.0000001
		want := 1.414
		got := solution(input, precise)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("sqrt2", func(t *testing.T) {
		input := 4.0
		precise := 0.0000001
		want := 2.0
		got := solution(input, precise)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input, precise float64) float64 {
	a, b := 1.0, input
	for {
		tmp := (a + b) / 2
		if tmp*tmp > input {
			// 较大
			if tmp*tmp-input < precise {
				return tmp
			} else {
				b = tmp
				continue
			}
		} else {
			// 较小
			if input-tmp*tmp < precise {
				return tmp
			} else {
				a = tmp
				continue
			}
		}
	}
}
