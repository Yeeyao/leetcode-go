package link_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
从一个单链表中删除一个节点，参数是将要删除的节点，同时该节点不是列表的尾节点，每个节点的数值都是唯一的
这里强调不是尾节点的目的是什么呢？
这里需要删除当前的节点，当前节点的上一个节点需要指向当前节点的下一个节点，同时，当前节点的下一个节点需要置 nil
但是这里传入的是当前的节点，如果获得上一个节点呢?（单链表就没有方法可以获得上一个节点吧？基于地址的也不行，这就要求编译器或者实现中内存限制）

这里直接修改数值，但是这样已经不是一般意义的删除了吧？这样根本就不是删除吧。。。
这里最后一个节点需要干掉，所以最后需要遍历到最后一个节点的前一个节点
*/
func deleteNode(node *ListNode) {
	for node.Next.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
	}
	node.Val = node.Next.Val
	node.Next = nil
}
