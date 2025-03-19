package linked_list

/*
主要是能不能想到归并排序
合并方式和21题一样
*/
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}

	left, right := 0, n-1
	mid := (right-left)/2 + left
	l1 := mergeKLists(lists[:mid+1])
	l2 := mergeKLists(lists[mid+1:])
	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}
