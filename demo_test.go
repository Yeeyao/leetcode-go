package array

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("quick sort", func(t *testing.T) {
		A := []int{0, 9, 7, 2, 4, 5}
		qsort(A)
		////if (get, want) {
		////	t.Errorf("got: %v, want: %v", get, want)
		//}
	})

	t.Run("binary search", func(t *testing.T) {
		nums := []int{0, 1, 2, 3, 4, 5}
		get := binarySearch(nums, 3)
		want := 3
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("sqrt", func(t *testing.T) {
		input := 2.0
		precise := 0.0000001
		want := 1.414
		got := sqrtNew(input, precise)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("sqrt2", func(t *testing.T) {
		input := 4.0
		precise := 0.0000001
		want := 2.0
		got := sqrtNew(input, precise)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("sqrtN", func(t *testing.T) {
		input := 2.0
		precise := 0.0000001
		want := 1.414
		got := sqrtNt(input, precise)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("sqrtN2", func(t *testing.T) {
		input := 4.0
		precise := 0.0000001
		want := 2.0
		got := sqrtNt(input, precise)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("fast mul", func(t *testing.T) {
		n := 2
		m := 4
		want := 8
		got := fastMul(n, m)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("fast fact", func(t *testing.T) {
		n := 2.0
		m := 4
		want := 16.0
		got := fastFact(n, m)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	newton sqrt
*/
func newtonSqrt(e float64) float64 {
	if e < 0 {
		return 0.0
	}
	guess := e
	for guess*guess > e {
		fmt.Println(guess)
		guess = (guess + e/guess) / 2
	}
	return guess
}

//func abs(c float64) float64 {
//	if c < 0 {
//		return -c
//	}
//	return c
//}

/*
quicksort
*/
//func quickSortHelper(arr []int, left, right, pivotIndex int) int {
//	// 交换最后的元素和 pivotIndex 让 pivotIndex 元素不需要参与分割
//	pivotVal := arr[pivotIndex]
//	arr[right], arr[pivotIndex] = arr[pivotIndex], arr[right]
//	// 遍历元素，将小于 pivotVal 的元素向右存放
//	for i := left; i < right; i++ {
//		if arr[i] <= pivotVal {
//			arr[i], arr[pivotIndex] = arr[pivotIndex], arr[i]
//			pivotIndex++
//		}
//	}
//	// 交换回来
//	arr[right], arr[pivotIndex] = arr[pivotIndex], arr[right]
//	return pivotIndex
//}
//
//func quickSort2(arr []int, left, right int) {
//	if right > left {
//		pivotIndex := left
//		pivotIndexN := quickSortHelper(arr, left, right, pivotIndex)
//		quickSortHelper(arr, left, pivotIndexN-1, pivotIndexN)
//		quickSortHelper(arr, pivotIndexN+1, right, pivotIndexN)
//	}
//}
//
//func quickSort(arr []int) {
//	quickSort2(arr, 0, len(arr)-1)
//}

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	// pick pivotIndex
	pivotIndex := rand.Int() % len(a)

	// 先交换 pivotIndex 这个不需要参与遍历
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	for i := range a {
		// 元素小于 pivotIndex 元素就存放到左边
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	// 将 pivotIndex 元素存放到合适的位置
	a[left], a[right] = a[right], a[left]
	qsort(a[:left])
	qsort(a[left+1:])
	return a
}

func qsort2(nums []int) {
	quickSort2(nums, 0, len(nums)-1)
}

func quickSort2(nums []int, l, r int) {
	if r > l {
		// 初始 pivot index
		pivotIndex := l
		pivotIndexN := quickSort3(nums, l, r, pivotIndex)
		quickSort2(nums, l, pivotIndexN-1)
		quickSort2(nums, pivotIndexN+1, r)
	}
}

func quickSort3(nums []int, l, r, pivotIndex int) int {
	// 先将 pivot 存放到最右边
	pivot := nums[pivotIndex]
	nums[r], nums[pivotIndex] = nums[pivotIndex], nums[r]
	// pivot 计数
	storedIndex := l
	for i := l; i < r; i++ {
		if nums[i] < pivot {
			nums[i], nums[storedIndex] = nums[storedIndex], nums[i]
			storedIndex++
		}
	}
	nums[storedIndex], nums[r] = nums[r], nums[storedIndex]
	return storedIndex
}

/*
	使用栈来保存每次遍历的左右位置
*/
//func quickSortDe(nums []int) {
//	l, r := 0, len(nums-1)
//	st := Stack{}
//	st.push((l, r))
//	for len(st) {
//		l, r = st.pop()
//		if l < r {
//			pivotIndexN := partitionDe(nums, l, r)
//			st.Push(l, pivotIndexN-1)
//			st.Push(pivotIndexN+1, r)
//		}
//	}
//}
//
//func partitionDe(nums []int, l, r int) {
//	pivot := nums[r]
//	storeIndex := l
//	for i := l; i < r; i++ {
//		if nums[i] < pivot {
//			nums[i], nums[storeIndex] = nums[storeIndex], nums[i]
//			storeIndex++
//		}
//	}
//	nums[storeIndex], nums[r] = nums[r], nums[storeIndex]
//	return storeIndex
//}

/*
	二分查找
*/
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

/*
	牛顿法求平方根
*/
func sqrtNew(input, precise float64) float64 {
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

/*
	牛顿法求平方根
*/
func sqrtNt(number, precise float64) float64 {
	guess := number / 2
	for abs(number-guess*guess) > precise {
		guess = (guess + number/guess) / 2
	}
	return guess
}

func abs(number float64) float64 {
	if number < 0 {
		return -number
	}
	return number
}

/*
	fast mul
*/
func fastMul(n, m int) int {
	if n == 0 || m == 0 {
		return 0
	}
	res := 0
	for m > 0 {
		if m&1 > 0 {
			res += n
		}
		n <<= 1
		m >>= 1
	}
	return res
}

/*
	fast fact
*/
func fastFact(base float64, exponent int) float64 {
	if base == 0 {
		return 0.0
	}
	// 先将指数转换为正数
	if exponent < 0 {
		exponent = -exponent
		base = 1 / base
	}
	res := 1.0
	// 这里需要统计指数的每个字段是否是 1，计算 k * x^n + k * x^n-1
	// 一个个位数来判断
	for exponent > 0 {
		// 当前位数非 0
		if exponent&1 > 0 {
			res *= base
		}
		// x^n 递增同时指数右移
		base *= base
		exponent >>= 1
	}
	return res
}
