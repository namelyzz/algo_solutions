package linked_list

/*
递归法
寻找到最小的可进入递归体的例子 1 -> 2 -> nil
此时，head 为 1
反转就是 2 指向 1，即 head.Next.Next = head
然后 1 指向 nil，即 head.Next = nil
这样我们就完成了这个最小例子的反转了
而对于更长的例子，其实本质上也是一样的，比如 1 -> 2 -> 3 -> nil
第一次递归反转（反转 2-3-nil 部分）会得到

	3
	⬇

1 -> 2 -> nil
之后的步骤也是一样的
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}

/*
迭代法
迭代法的精髓是有一条新的链表，它是空的，我们把当前链表的节点一个个搬过去
不要固执的陷入在同一条链表上操作的循环中
*/
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
