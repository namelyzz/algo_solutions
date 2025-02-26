package linked_list

/*
也是考察基本功
*/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd, even := head, head.Next
	p1, p2 := odd, even
	for p2 != nil && p2.Next != nil {
		p1.Next = p2.Next
		p1 = p1.Next
		p2.Next = p1.Next
		p2 = p2.Next
	}

	p1.Next = even
	return odd
}
