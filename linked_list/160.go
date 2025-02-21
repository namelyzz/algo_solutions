package linked_list

/*
l1 一直走，走到尽头就转入 headB
l2 一直走，走到尽头就转入 headA
如果有交点，最终会在交点相遇
如果没有交点，则最终都会走到 nil 节点（其实 nil 也算是两条链表的交点）

为什么有交点，就一定会在交点相遇呢（nil 也算两天链表的交点）？
核心原理就是这种走法，最终路径长度是相等的，即“对齐”两条链表的长度差，从而保证如果有交点，两指针最终会在交点相遇。
假设：
- 链表A独有部分长度为 a
- 链表B独有部分长度为 b
- 公共部分长度为 c
指针l1的路径：
- 先走完A的独有部分：a
- 切换到B，走完B的独有部分：b
- 到达交点：总共走了 a + b 步
指针l2的路径：
- 先走完B的独有部分：b
- 切换到A，走完A的独有部分：a
- 到达交点：总共走了 b + a 步
关键点——两个指针都会走过：
- 对方链表的独有部分
- 自己链表的独有部分
- 最终在交点相遇
因此，让两个指针走相同的总路径长度，从而保证有交点时必然相遇
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headB
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headA
		}
	}
	return l1
}
