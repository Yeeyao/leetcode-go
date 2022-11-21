package tree

import "sort"

/*
给定二叉树的根，计算该二叉树的 vertical order traversal
对位于 r, c 的节点，其左孩子位置是 r+1,c-1 右孩子位置是 r+1,c+1
这里的遍历是从最左边到最右边对每一列自顶向下的遍历顺序，相同的列或者行中有多个节点，遍历的时候需要按照它们的大小进行排序
返回二叉树的这个遍历的顺序
*/

/*
	要在二叉树中找到同一列的节点，然后还要将同一列的节点的数值进行升序排列
	这里看起来就是要中序遍历，但是同一列的要怎么判断呢，根据 root 的序号确定左右子树的节点的序号
	然后将同一列的数值的放到一起之后排序，横坐标是否有用呢？需要实时计算当前的 root 的坐标吧

	这里只有在两个节点的坐标都相同的时候才需要对它们的数值进行升序排列，如果坐标不同则小的出现在前面
	难道需要把坐标也保存下来进行排序吗？
	没有什么意思的 hard 题目，用来吓人
*/
type nodeWithCoordinate struct {
	val int
	row int
	col int
}

var colListMap map[int]*[]nodeWithCoordinate

func verticalTraversal(root *TreeNode) [][]int {
	colListMap = make(map[int]*[]nodeWithCoordinate)
	verticalTraversalHelper(root, 0, 0)
	var res [][]int
	var colList []int
	for k, v := range colListMap {
		colList = append(colList, k)
		colList := *v
		sort.Slice(colList, func(i, j int) bool {
			a, b := colList[i], colList[j]
			return a.col < b.col || a.col == b.col && a.row < b.row || a.col == b.col && a.row == b.row && a.val < b.val
		})

	}
	sort.Ints(colList)
	res = make([][]int, len(colList))
	for i := 0; i < len(colList); i++ {
		nodeWithCoordinateList, _ := colListMap[colList[i]]
		var nodeList []int
		for _, v := range *nodeWithCoordinateList {
			nodeList = append(nodeList, v.val)
		}
		res[i] = append(res[i], nodeList...)
	}
	return res
}

func verticalTraversalHelper(root *TreeNode, row, col int) {
	if root == nil {
		return
	}
	if found, ok := colListMap[col]; ok {
		*found = append(*found, nodeWithCoordinate{
			val: root.Val,
			row: row,
			col: col,
		})
	} else {
		colListMap[col] = &[]nodeWithCoordinate{{
			val: root.Val,
			row: row,
			col: col,
		}}
	}
	verticalTraversalHelper(root.Left, row+1, col-1)
	verticalTraversalHelper(root.Right, row+1, col+1)
}
