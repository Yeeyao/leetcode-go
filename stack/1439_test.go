package stack

/*
给你一个 m n 的矩阵 mat，以及一个整数 k ，矩阵中的每一行都以非递减的顺序排列。
你可以从每一行中选出 1 个元素形成一个数组。返回所有可能数组中的第 k 个 最小 数组和。

按照这里多路归并的思路，应该是每一行维护一个指向当前最小值的指针
一开始全部都指向第一个元素，得到的是第一个最小的数组和，然后将所有的指针的下一个进行比较，其中最小的就是需要移动的行的指针，作为第二个最小的数组和
但是，如果这里有多个指针的下一个是相同的数值，则会出现多种可能的结果，每个不同的指针表示一种可能的结果(因为会影响后续的指针移动)，不同的结果得到的数组和不一定相同。
这里提到每次分裂之后，极值发生变化，因此是一个动态求极值的题目，可以使用堆

堆配合元组的多路归并方法
*/

func kthSmallest(mat [][]int, k int) int {

}
