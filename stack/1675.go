package stack

/*
给你一个由 n 个正整数组成的数组 nums 。
你可以对数组的任意元素执行任意次数的两类操作：
可以执行的操作：元素是偶数则除以 2，元素是奇数则乘上 2
可以对数组元素进行操作然后求出其最小偏移量 数组的 偏移量 是数组中任意两个元素之间的 最大差值
看起来可以执行很多次，但是对元素的数值的影响是确定的。

元素一开始是偶数，则只能除以 2 直到变成奇数或者 1
元素一开始是奇数，则乘以 2 就结束了

这里不限制操作次数，但是总要找到一个目标数值吧
- 这里的目标数值不只是数组中的所有的元素数值本身，还包含了它们不断除以 2 的数值和不断乘以 2 的数值

这里暴力解法就是对每个元素
	根据奇偶计算可能的数值
	其他元素就根据相对目标数值的大小进行操作尽可能接近目标
	其中，因为所有的数值的目标数值可能出现重复，因此每次计算出来一个目标数值的偏移量就可以保存下来，避免重复计算
	同时，如果某个目标数值当前的偏移量比当前的最小偏移量的目标数值还大，则可以提前终止，这时记录的是终止的比偏移量小的数值

这里转换一下思路，则数组中每个元素都可以通过操作得到一个它可以转变的数组，因此这里的题目就变成了从 n 个数组每个数组选择一个元素组成新的数组，
使得新数组的偏移值最小，这个就有点类似 632 了
*/

func minimumDeviation(nums []int) int {

}
