package string

import "testing"

func TestPro(t *testing.T) {
	t.Run("657. Robot Return to Origin", func(t *testing.T) {
		input := "UD"
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

// 需要全部都统计出来，之后再判断出现次数的独立
func solution(moves string) bool {
	vertical := 0
	horizonal := 0
	for _, v := range moves {
		switch v {
		case 'U':
			vertical++
		case 'D':
			vertical--
		case 'L':
			horizonal++
		case 'R':
			horizonal--
		}
	}
	if vertical == 0 && horizonal == 0 {
		return true
	}
	return false
}
