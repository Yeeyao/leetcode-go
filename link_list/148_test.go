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

O(nlogn) 的时间复杂度和 O(1) 的空间复杂度。其中最适合链表的排序算法是归并排序。

自顶向下的方法，找到链表中点：
 - 使用快慢指针，一个走一步，一个走两步，快指针达到结尾的时候，慢指针的位置是链表的中点
 - 对两个子链表进行排序
 - 将两个子链表进行归并
*/

func sortList(head *ListNode) *ListNode {
	// fast 本身可能是 nil
	if head == nil || head.Next == nil {
		return head
	}
	// 这里死循环了 fast, slow := head, head
	// 先找到中间节点，这里 fast 指向下一个
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 这里需要从中间断开才能实现两端分别排序
	temp := slow.Next
	// head.Next 已经不会是 nil 然后 fast 的判断保证了 slow.Next 非 nil
	slow.Next = nil
	leftListHead := sortList(head)
	rightListHead := sortList(temp)
	// 归并 21
	retCur := &ListNode{}
	retHead := retCur
	for leftListHead != nil && rightListHead != nil {
		if leftListHead.Val < rightListHead.Val {
			retCur.Next = leftListHead
			leftListHead = leftListHead.Next
		} else {
			retCur.Next = rightListHead
			rightListHead = rightListHead.Next
		}
		retCur = retCur.Next
	}
	if leftListHead != nil {
		retCur.Next = leftListHead
	} else {
		retCur.Next = rightListHead
	}
	return retHead.Next
}
