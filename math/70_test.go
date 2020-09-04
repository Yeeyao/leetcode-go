package math

/*
	70 Climbing Stairs 爬楼梯
	f(n) 表示第 n 步时的所有路径
	f(n) = f(n-1) + f(n-2) 因为从 f(n-1) 到 f(n) 路径唯一，然后 f(n-2) 同理，所以这里不同路径就是前面两个的和
	然后初始值 f(0) = 1 f(1) = 1
*/

func climbStairs(n int) int {
	dp := make([]int, n+1)
	if n == 0 || n == 1 {
		return 1
	}
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

/*
	直接用两个变量
*/
func climbStars(n int) int {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
