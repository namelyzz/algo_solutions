package linked_list

/*
比较考验链表基本功的做法，也比较符合直觉，就是找到要反转的开头和结尾，然后直截了当的逐个换位
缺点是要在链表上标出来很多个位置，容易搞晕了
所以这种做法，在实际面试中我觉得可能比较好想到，但是做起来会比较困难
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	r := pre.Next
	for i := 0; i < right-left; i++ {
		r = r.Next
	}
	end := r.Next

	for pre.Next != r {
		cur := pre.Next
		next := cur.Next
		cur.Next = end
		pre.Next = next
		r.Next = cur
		end = cur
	}

	return dummy.Next
}

/*
第二种做法，同样也需要找到反转部分的头部和尾部，然后结合之前的反转链表的解法
1. 定位反转起始位置的前一个节点（pre）
2. 反转指定区间内的节点(之前的迭代法)
3. 重新连接反转后的部分与原始链表
下面，假设链表：1 → 2 → 3 → 4 → 5，left = 2，right = 4 开始注释
*/
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	/*
		这一部分为定位 pre 节点 —— 反转区间的前一个节点
		不难找到 pre = 节点 1
	*/
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	/*
		反转区间内的节点，用迭代法解
		第一次循环：list = 2 → nil，cur = 3
		第二次循环：list = 3 → 2 → nil，cur = 4
		第三次循环：list = 4 → 3 → 2 → nil，cur = 5
		注意，这里最终 cur 会移动到 5，也就是反转结束后，cur 会指向反转区间之后的一个节点
	*/
	var list *ListNode
	cur := pre.Next
	for i := 0; i < right-left+1; i++ {
		next := cur.Next
		cur.Next = list
		list = cur
		cur = next
	}

	/*
		重新连接链表
		1. 原反转区间的头节点指向剩余部分，
			也就是 节点2(原来的头节点，即 pre.Next) -> 节点5(cur)
		2. pre 接入反转后的新头节点，也就是 节点1 → 节点4
	*/
	pre.Next.Next = cur
	pre.Next = list
	return dummy.Next
}
