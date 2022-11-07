package tree

import "fmt"

/*
	对比 95 这次要求返回所有的结果树，需要注意是 BST
	DFS 递归，只需要返回所有的结果树的 head
	开始，选择任意一个作为 head 结束，所有的节点都被使用了 答案，将结束的时候的 head 保存到返回的 head slice
*/

/*
	在所有可能的版本上进行剪枝，BST 的条件，看起来思路是穷举法然后提前剪枝，需要确定怎么穷举法
	这个是所有可能的版本,BST 则需要对部分结果进行剪枝就好了，递归之前需要从 head 开始判断？不用吧， head 之前的都是满足的
	需要怎么确定当前的节点的父节点是哪个呢？
*/
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTreesHelper(1, n)
}

/*
	类似 96 的思路试试，先选择一个 root，数值是 i 然后左边的子树从 1,i-1 右边是 i+1,n。
	[ref](https://leetcode.cn/problems/unique-binary-search-trees-ii/solution/bu-tong-de-er-cha-sou-suo-shu-ii-by-leetcode-solut/)
*/
func generateTreesHelper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	var allTrees []*TreeNode
	for i := start; i < end+1; i++ {
		// 分别构造左右子树集合，本来就是使用小于 i 的数值和大于 i 的数值分别构建左右子树，所以左右子树是一定满足 bfs 的
		// 因此下面直接笛卡尔积就完事了
		leftTrees := generateTreesHelper(start, i-1)
		rightTrees := generateTreesHelper(i+1, end)
		// 最后进行笛卡尔积
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				// 子树已经是满足 BST 条件了，只需要判断两个子树的 root 和新的 root
				allTrees = append(allTrees, &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				})
			}
		}
	}
	return allTrees
}

/*
	TODO: 构造所有可能的树呢？思路是什么？
*/
func generateTreesAll(n int) []*TreeNode {
	nodeList := make([][]*TreeNode, 0)
	for i := 1; i <= 1; i++ {
		head := &TreeNode{Val: i}
		usedI := make(map[int]struct{})
		usedI[i] = struct{}{}
		generateTreesAllHelper(n, usedI, []*TreeNode{head}, &nodeList)
	}
	return regenerateFromNodeList(nodeList)
}

/*
head 确定了，然后下面需要选择下一个节点，作为左边或者右边，同时不能选择已经出现过的数字
需要 map 记录已经使用的数字？每个都需要作为当前 head 的左子树或者右子树的节点，head 上面的默认已经构造好了
需要加上 BST 的限制，应该有很多的剪枝
为什么没有出现所有（遗漏了一些情况），因为这里循环下去，使用的 head 都是同一个，然后不同的递归分支全部都只是修改同一个结果
因此需要先将数值保存，然后重新构造结果的树，没有办法提前将返回的树进行组合。因此，应该是使用数组，然后下面有分支的时候需要复制一个原来的数组
每个子循环只需要关心自己的结果就好了
*/
func generateTreesAllHelper(n int, usedI map[int]struct{}, tempList []*TreeNode, nodeList *[][]*TreeNode) {
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

	// slice 的坑，需要 copy 一份，map 在更新的时候需要重置，其实还是递归的数据重复判断部分不熟悉
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
		// 本轮使用了某个元素，需要重置成未使用状态
		if setNewVal {
			delete(usedI, i)
		}
	}
	if isAllUsed {
		*nodeList = append(*nodeList, tempList)
	}
}

func regenerateFromNodeList(nodeList [][]*TreeNode) []*TreeNode {
	ret := make([]*TreeNode, 0)
	for _, v := range nodeList {
		head := regenerateOne(v)
		ret = append(ret, head)
	}
	return ret
}

func regenerateOne(nodeList []*TreeNode) *TreeNode {
	fmt.Printf("nodeList: %v", nodeList)
	head := &TreeNode{}
	fakeHead := TreeNode{Left: head}
	head.Val = nodeList[0].Val
	for i := 1; i < len(nodeList); i = i + 2 {
		head.Left = nodeList[i]
		head.Right = nodeList[i+1]
		if head.Left != nil {
			head = head.Left
		} else {
			head = head.Right
		}
	}
	printTree(fakeHead.Left)
	return fakeHead.Left
}

func printTree(root *TreeNode) {
	var retTree []*TreeNode
	if root != nil {
		retTree = append(retTree, root)
	}
	road := ""
	for len(retTree) > 0 {
		head := retTree[0]
		road += fmt.Sprintf(" %v ", head.Val)
		retTree = retTree[1:]
		if head.Left != nil {
			retTree = append(retTree, head.Left)
		} else {
			road += fmt.Sprintf(" lnil ")
		}
		if head.Right != nil {
			retTree = append(retTree, head.Right)
		} else {
			road += fmt.Sprintf(" rnil ")
		}
	}
	fmt.Println(road)
}
