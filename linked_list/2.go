package linked_list

/*
思路：
1。看完题目，注意到链表表示的数字是反过来的，如 2 -> 4 -> 3 表示数字 342
2. 因此可以从链表头逐个相加
3. 需要注意的是进位的处理，增加变量来存进位信息
4. 只要有一条链表还没遍历完，就继续遍历下去，已经遍历完的链表视为 0。因为还有进位信息要处理，可能要一直计算，例子如 9 -> 9 -> 1 和 9 -> 9
5. 最后，如果两个链表都遍历完了，还需要看是否还有进位信息，如 99 + 99，我们还需要处理最后的进位
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	var carry, v1, v2 int // carry: 进位, v1/v2: 当前位的值

	// 遍历两个链表，直到都到达末尾
	for l1 != nil || l2 != nil {
		// 处理l1的当前位，如果l1已遍历完则值为0
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
			l1 = l1.Next
		}

		// 处理l2的当前位，如果l2已遍历完则值为0
		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		val := sum % 10
		carry = sum / 10
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}

	// 处理最后的进位（如果还有进位，需要额外创建一个节点）
	if carry == 1 {
		cur.Next = &ListNode{Val: 1}
	}
	return dummy.Next
}
