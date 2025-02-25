package linked_list

/*
思路：
 1. 注意题目的条件（提示部分）。k 和链表数目没有直接关系，这说明 k 可能比链表长度还大
    如果 k > 链表长度，链表的旋转会循环，我们先算出长度，然后对 k 取余

2. 快慢指针找到要旋转的部分，接到头部
*/
func rotateRight(head *ListNode, k int) *ListNode {
	/*
	   基础情况：如果链表为空或只有一个节点，无论怎么旋转结果都一样
	   这是最基本的边界情况，先处理掉可以简化后续逻辑
	   推荐保持这个习惯，即使有时候主体逻辑也能兼容这个基本情况也最好写上
	   因为链表的递归解法一般也是这个终止条件，先写上，递归不容易漏
	*/
	if head == nil || head.Next == nil {
		return head
	}

	/*
	   计算有效旋转次数
	*/
	var listLen int
	cur := head
	for cur != nil {
		listLen++
		cur = cur.Next
	}
	k %= listLen
	// 如果旋转次数是链表长度的整数倍，相当于没有旋转，这个特例很容易遗忘，要记得补上
	if k == 0 {
		return head
	}

	/*
	   快慢指针定位分割点
	   - 要让链表右旋转k次，实际上是把后k个节点移到前面
	   - 快指针先走k步，这样快慢指针之间就相隔k个节点
	   - 然后同时移动，当快指针到达尾部时，慢指针正好指向新链表的尾部
	   - 此时 slow.Next 就是新链表的头节点
	*/
	slow, fast := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	res := slow.Next
	fast.Next = head
	slow.Next = nil
	return res
}
