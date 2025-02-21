package linked_list

/*
比较简单就用快慢指针
相遇了就是有环
fast 找到了 nil 节点则说明无环
*/
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
