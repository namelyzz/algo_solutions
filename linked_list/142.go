package linked_list

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// 与 141 一样的思路，找到链表的环
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	// 如果没有环，直接 return nil
	if fast == nil || fast.Next == nil {
		return nil
	}

	// 到这一步，说明有环
	// slow 回到开头，与 fast 同时移动一步，最后相遇就是环的入口
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
