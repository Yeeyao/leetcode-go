package stack

import "container/heap"

/*
	给定 points 数组表示点 points[i] = [xi, yi]，数值是点的坐标，给定整数 k，给出离原点最近的 k 个点
	同样使用堆处理
*/
func kClosest(points [][]int, k int) [][]int {
	h := PointDistanceHeap{distanceList: []*PointsWithDistance{}}
	for _, p := range points {
		heap.Push(&h, &PointsWithDistance{
			x:        p[0],
			y:        p[1],
			distance: int64(p[0]*p[0] + p[1]*p[1]),
		})
	}
	var res [][]int
	for k > 0 {
		top := heap.Pop(&h).(*PointsWithDistance)
		res = append(res, []int{top.x, top.y})
		k--
	}
	return res
}

type PointsWithDistance struct {
	x int
	y int
	// 这里不需要记录实际的距离吧，只需要记录平方和就好了
	distance int64
}

type PointDistanceHeap struct {
	distanceList []*PointsWithDistance
}

func (h *PointDistanceHeap) Len() int {
	return len(h.distanceList)
}

func (h *PointDistanceHeap) Less(i, j int) bool {
	return h.distanceList[i].distance < h.distanceList[j].distance
}

func (h *PointDistanceHeap) Swap(i, j int) {
	temp := h.distanceList[i]
	h.distanceList[i] = h.distanceList[j]
	h.distanceList[j] = temp
}

func (h *PointDistanceHeap) Push(i interface{}) {
	h.distanceList = append(h.distanceList, i.(*PointsWithDistance))
}

func (h *PointDistanceHeap) Pop() interface{} {
	top := h.distanceList[len(h.distanceList)-1]
	h.distanceList = h.distanceList[:len(h.distanceList)-1]
	return top
}
