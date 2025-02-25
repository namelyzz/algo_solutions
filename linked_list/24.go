package linked_list

/*
考察链表的基本功吧，没什么特别难的地方
定义好 pre，cur，next，草稿纸上涂涂画画，大概就明白怎么两两调换节点了
注意边界的问题，就是 pre.Next != nil && pre.Next.Next != nil 这一句
条件1：pre.Next != nil

	如果 pre.Next == nil，说明没有第一个节点可以交换
	这是最基本的前提条件

条件2：pre.Next.Next != nil

	如果 pre.Next.Next == nil，说明有第一个节点但没有第二个节点
	无法完成"两两"交换，只能保持原样
*/
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		cur, next := pre.Next, pre.Next.Next
		cur.Next = next.Next
		next.Next = cur
		pre.Next = next
		pre = cur
	}
	return dummy.Next
}
