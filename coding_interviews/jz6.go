package coding_interviews

/*
JZ6 从尾到头打印链表

输入一个链表的头节点，按链表从尾到头的顺序返回每个节点的值（用数组返回）。
*/

/*
思路：
两次遍历
第一次遍历链表，求得链表长度，然后创建一个和链表长度一致的数组
第二次遍历链表，数组从尾到头记录链表的值
*/
func printListFromTailToHead(head *ListNode) []int {
	cur := head
	cnt := 0
	for cur != nil {
		cur = cur.Next
		cnt++
	}

	res := make([]int, cnt)
	for head != nil {
		res[cnt-1] = head.Val
		head = head.Next
		cnt--
	}

	return res
}
