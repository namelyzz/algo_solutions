package linked_list

/*
快慢指针找到链表的中点
然后一分为二
将后半部分反转
然后逐个插入回第一部分
*/
func reorderList(head *ListNode) {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	list := slow.Next
	slow.Next = nil

	var pre *ListNode
	cur := list
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	cur = head
	for pre != nil {
		next1 := cur.Next
		next2 := pre.Next
		cur.Next = pre
		pre.Next = next1
		pre = next2
		cur = next1
	}
}
