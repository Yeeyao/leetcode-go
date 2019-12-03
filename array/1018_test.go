package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("1018. Binary Prefix Divisible By 5", func(t *testing.T) {
		input := []int{0, 1, 1}
		want := []bool{true, false, false}
		got := solution(input)
		if !BoolSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1018. Binary Prefix Divisible By 5 2", func(t *testing.T) {
		input := []int{0, 1, 1, 1, 1, 1}
		want := []bool{true, false, false, false, true, false}
		got := solution(input)
		if !BoolSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1018. Binary Prefix Divisible By 5 3", func(t *testing.T) {
		input := []int{1, 1, 1, 0, 1}
		want := []bool{false, false, false, false, false}
		got := solution(input)
		if !BoolSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1018. Binary Prefix Divisible By 5 4", func(t *testing.T) {
		input := []int{1, 1, 0, 0, 0, 1, 0, 0, 1}
		want := []bool{false, false, false, false, false, false, false, false, false}
		got := solution(input)
		if !BoolSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	利用取模的性质
	Use the fact that (ab + c)%d is same as ((a%d)(b%d) + c%d)%d.
	We now have the relation new_number%5 = ((old_number%5)*2 + d)%5;
*/
func solution(input []int) []bool {
	var retArr []bool
	num := 0
	for _, v := range input {
		num = (num*2 + v) % 5
		retArr = append(retArr, num == 0)
	}
	return retArr
}

/*
	数字太大了，就会溢出
*/
//func solution(input []int) []bool {
//	var retArr []bool
//	sum := 0
//	for _, v := range input {
//		sum = sum*2 + v
//		if sum%5 == 0 {
//			retArr = append(retArr, true)
//		} else {
//			retArr = append(retArr, false)
//		}
//	}
//	return retArr
//}

func BoolSliceEqual(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
