package linked_list

/*
思路：
暴力解法，一次遍历找倒数第 n 个节点，第二次遍历删除它。
但我们希望能在 一次遍历 中解决这个问题。
为了实现这一目标，我们可以利用两个指针——快慢指针

核心步骤分析：
我们可以利用两个指针，快指针（fast）和慢指针（slow），其中 fast 指针先走 n 步，然后两者一起走，直到 fast 指针到达链表的末尾。
由于 fast 比 slow 先走了 n 步，当 fast 到达链表末尾时，slow 就指向了倒数第 n 个节点的前一个节点，这样我们就能轻松删除目标节点。
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
