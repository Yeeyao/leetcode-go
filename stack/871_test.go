package stack

import "container/heap"

/*
	最低加油次数
	汽车从起点出发驶向目的地，该目的地位于出发位置东面 target 英里处。
	沿途有加油站 station[i][j] 表示加油站离起点 i 英里，加油站有 j 升燃料。
	假设汽车油箱大小是无限的，出发前有然后 startFuel 升，经过加油站会将加上加油站所有的油
	判断为了到达目的地，需要最少加油几次，如果无法达到目的地就返回 -1
	到达加油站剩余 0 升还可以加油，到达目的地剩余 0 升也算到达了

	为了最少加油次数，则不需要加油就不加，什么时候必须加油则是如果不加油就不能到达下一个目的地的时候
	这里如果必须加油，在哪里加油呢？就在含有最大汽油的油站加油，这样转换为动态求极值的场景，适合使用堆

	这里判断油量是判断到达下一个加油站的距离使用的油量
	算法描述
		每经过一个加油站，将它的油量加到堆里，经过才会入堆
		往前开，只要油量大于 0 就继续开
		如果油量小于 0 从堆中获取最大的油量加到油箱，如果还是小于 0 就继续找下一个
		如果加完油大于 0 继续开，重复步骤。否则返回 -1 表示无法到达目的地
	times: 1
*/

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	// 这里将目标也放入加油站
	gs := &gasStation{
		stationInfoList: []*stationInfo{},
	}
	// 将终点也作为其中一个加油站
	stations = append(stations, []int{target, 0})
	addTimes := 0
	leftFuel := startFuel
	// 记录当前已经走过的距离
	runDistance := 0
	for _, v := range stations {
		distance, fuel := v[0], v[1]
		// 本次需要走的距离 = 加油站的距离 - 已经走过的距离
		leftFuel -= distance - runDistance
		// 不能走到这个加油站则继续加油，这里等于 0 不需要加油，因为能走到这个加油站
		for leftFuel < 0 && gs.Len() > 0 {
			addGs := heap.Pop(gs).(*stationInfo)
			leftFuel += addGs.gasVolume
			addTimes++
		}
		// 继续加油也不能走完就返回 -1
		if leftFuel < 0 {
			return -1
		}
		// 经过这个加油站，加入到 heap
		heap.Push(gs, &stationInfo{
			distance:  distance,
			gasVolume: fuel,
		})
		// 已经走过的距离等于这个加油站的距离
		runDistance = distance
	}
	return addTimes
}

type gasStation struct {
	stationInfoList []*stationInfo
}

type stationInfo struct {
	distance  int
	gasVolume int
}

func (sp *gasStation) Less(i, j int) bool {
	return sp.stationInfoList[i].gasVolume < sp.stationInfoList[j].gasVolume
}

func (sp *gasStation) Swap(i, j int) {
	temp := sp.stationInfoList[j]
	sp.stationInfoList[j] = sp.stationInfoList[i]
	sp.stationInfoList[i] = temp
}

func (sp *gasStation) Len() int {
	return len(sp.stationInfoList)
}

func (sp *gasStation) Push(v interface{}) {
	sp.stationInfoList = append(sp.stationInfoList, v.(*stationInfo))
}

func (sp *gasStation) Pop() interface{} {
	top := sp.stationInfoList[len(sp.stationInfoList)-1]
	sp.stationInfoList = sp.stationInfoList[:len(sp.stationInfoList)-1]
	return top
}
