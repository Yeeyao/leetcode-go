package link_list

/*
148. Sort List 类似 234
Sort a linked list in O(n log n) time using constant space complexity.
链表版本的快排？归并排序
[ref](https://leetcode-cn.com/problems/sort-list/solution/sort-list-gui-bing-pai-xu-lian-biao-by-jyd/)
这里要怎么划分，没办法，只能从开头向后遍历找到中点
中点的查找使用双指针

先判断 head 以及 head.Next 如果任意为 nil 就返回 head
使用 fast 和 slow 指针找到中间节点
将 slow.Next = nil left, right 等于递归调用 head 以及 slow.Next
初始化两个新节点，其中一个用来进行归并处理
最后返回未使用的节点
*/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 找到中间节点 slow 这里因为 fast 在前面，所以只需要判断 fast
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 如果是奇数节点，slow 就是中间节点，否则是中间的前一个
	// 这里递归调用 head 和 slow.Next 将后者设置为 nil 进行截断
	temp := slow.Next
	slow.Next = nil
	// 因为最后需要返回，所以这里需要递归的结果
	left := sortList(head)
	right := sortList(temp)
	// 因为这里初始化的是空节点，所以元素加到空节点的 next 中
	newHead := &ListNode{}
	res := newHead
	for left != nil && right != nil {
		if left.Val < right.Val {
			newHead.Next = left
			left = left.Next
		} else {
			newHead.Next = right
			right = right.Next
		}
		newHead = newHead.Next
	}
	if left != nil {
		newHead.Next = left
	} else {
		newHead.Next = right
	}
	return res.Next
}
