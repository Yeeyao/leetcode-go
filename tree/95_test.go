package tree

import "fmt"

/*
	对比 95 这次要求返回所有的结果树，需要注意是 BST
	DFS 递归，只需要返回所有的结果树的 head
	开始，选择任意一个作为 head 结束，所有的节点都被使用了 答案，将结束的时候的 head 保存到返回的 head slice
*/

/*
	这个是所有可能的版本,BST 则需要对部分结果进行剪枝就好了
*/
func generateTreesAll(n int) []*TreeNode {
	ret := make([]*TreeNode, 0)
	nodeList := make([][]*TreeNode, 0)
	for i := 1; i <= 1; i++ {
		head := &TreeNode{Val: i}
		usedI := make(map[int]struct{})
		usedI[i] = struct{}{}
		generateTreesHelper(n, usedI, []*TreeNode{head}, &nodeList)
	}
	fmt.Println(nodeList)
	// 之后就根据 nodeList 重新构建树
	return ret
}

/*
head 确定了，然后下面需要选择下一个节点，作为左边或者右边，同时不能选择已经出现过的数字
需要 map 记录已经使用的数字？每个都需要作为当前 head 的左子树或者右子树的节点，head 上面的默认已经构造好了
需要加上 BST 的限制，应该有很多的剪枝
为什么没有出现所有（遗漏了一些情况），因为这里循环下去，使用的 head 都是同一个，然后不同的递归分支全部都只是修改同一个结果
因此需要先将数值保存，然后重新构造结果的树，没有办法提前将返回的树进行组合。因此，应该是使用数组，然后下面有分支的时候需要复制一个原来的数组
每个子循环只需要关心自己的结果就好了
*/
func generateTreesHelper(n int, usedI map[int]struct{}, tempList []*TreeNode, nodeList *[][]*TreeNode) {
	// head 重新使用新的地址，但是上一个的指向将会丢失
	//fmt.Printf("fakeHead addr %p, fakeHead next val: %v\n", &fakeHead, fakeHead.Left.Val)
	//if fakeHead.Left.Left != nil {
	//	fmt.Printf("fll addr %p, val: %v\n", &fakeHead.Left.Left, fakeHead.Left.Left.Val)
	//}
	//if fakeHead.Left.Right != nil {
	//	fmt.Printf("flr addr %p, val: %v\n", &fakeHead.Left.Right, fakeHead.Left.Right.Val)
	//}
	//fmt.Printf("usedI addr %p, usedI: %v\n", &usedI, usedI)
	//fmt.Printf("head addr %p, head val: %v\n", &head, head.Val)
	//fmt.Printf("temp list addr: %p, temp list: %v\n", &tempList, tempList)

	// slice 的坑，需要 copy 一份，map 在更新的时候需要重置
	isAllUsed := true
	tempListLen := len(tempList)
	for i := 1; i <= n; i++ {
		setNewVal := false
		if _, ok := usedI[i]; !ok {
			setNewVal = true
			isAllUsed = false
			usedILeft := usedI
			usedILeft[i] = struct{}{}
			leftNode := &TreeNode{Val: i}
			leftTempList := append(tempList, []*TreeNode{leftNode, nil}...)
			generateTreesHelper(n, usedILeft, leftTempList, nodeList)
			rightTempList := tempList[:tempListLen]
			rightTempList = append(rightTempList, []*TreeNode{nil, leftNode}...)
			generateTreesHelper(n, usedILeft, rightTempList, nodeList)
		}
		if setNewVal {
			delete(usedI, i)
		}
	}
	if isAllUsed {
		*nodeList = append(*nodeList, tempList)
	}
}