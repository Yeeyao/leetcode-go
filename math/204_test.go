package math

import "testing"

func TestPro(t *testing.T) {
	t.Run("204. count primes", func(t *testing.T) {
		num := 9
		want := true
		got := solution(num)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
264. Ugly Number II
Count the number of prime numbers less than a non-negative number, n.
直接计算好了
-[ref](https://leetcode-cn.com/problems/count-primes/solution/ji-shu-zhi-shu-by-leetcode-solution/)

*/

/*
埃氏筛
如果 x 是质数，则 2x, 3x, ... 是合数。因此可以对合数进行标记。
从小到大遍历每个数，如果这个数为质数，则将其所有的倍数都标记为合数（除了该质数本身），即 0，这样在运行结束的时候我们即能知道质数的个数。
这种方法显然不会将质数标记成合数；另一方面，当从小到大遍历到数 x 时，倘若它是合数，则它一定是某个小于 x 的质数 y 的整数倍，
故根据此方法的步骤，我们在遍历到 y 时，就一定会在此时将 x 标记为 isPrime[x]=0。因此，这种方法也不会将合数标记为质数。
同时，不需要从 2x 开始标记，只需要从 x * x 开始标记，因为前面的过程已经标记了
判断过程中就进行标记
*/
func solution(n int) int {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	res := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			res++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return res
}

// 超时
func solution2(n int) int {
	res := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			res++
		}
	}
	return res
}

func isPrime(n int) bool {
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return false
}
