package linked_list

/*
1. 快慢指针找到链表中点
2. 合并两个有序链表
3. 归并排序的思路，使用递归
*/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMid(head)
	l1, l2 := head, mid.Next
	mid.Next = nil

	left := sortList(l1)
	right := sortList(l2)
	return merge(left, right)
}

func getMid(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	} else {
		l2.Next = merge(l1, l2.Next)
		return l2
	}
}
