package stack

/*
	给定建筑物高度数组以及 n 个 bricks 以及 n 个 ladders。需要从 0 号建筑物开始不断向后移动，从建筑物 i 到 i + 1 移动时：
	- 如果建筑物 i 高度大于等于 i + 1 则不需要梯子和砖块
	- 如果建筑物 i 高度小于 i + 1，可以使用一架梯子和 h[i+1]-h[i] 个砖块
	- 使用最佳的方式使用梯子和砖块，返回可以到达的最远的建筑物的下标

	这里的关键是怎么使用砖块，怎么感觉有点动态规划的味道
*/

func furthestBuilding(heights []int, bricks int, ladders int) int {

}
