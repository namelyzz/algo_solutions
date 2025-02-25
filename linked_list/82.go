package linked_list

/*
思路：
 1. 使用哑节点，可以快速返回到链表头部（因为 head 也可能因为重复被删）
 2. 双指针思路，pre 指向当前已处理的链表的尾部，cur 用于遍历剩余的节点
 3. 题目细节：因为链表是排序的，所以重复元素都是连续的
 4. 两种情况：
    a. cur 和它的下一个节点 val 不同，无事发生，正常遍历下去，移动 pre 和 cur
    b. cur 和它的下一个节点 val 不同，需要一直找到重复的元素，然后把这些重复的元素全部删掉
 5. 针对 4.b 理想的情况就是找到最后一个重复数，然后 pre 指向重复数的下一个节点。注意！此时 pre 不用移动
 6. pre 为什么不用移动？
    因为 pre 及其之前的节点依然是以处理的，而当前已处理的因为都是重复数，全都删除了。
    此时 pre.Next 依然是未处理的节点

这里的要考虑以上两种情况，pre 的移动方式不同
*/
func deleteDuplicates(head *ListNode) *ListNode {
	/*
	   初始化部分：
	   创建一个哑节点指向头节点，这样可以统一处理头节点可能被删除的情况
	   pre指针初始指向哑节点，用于记录当前已处理链表的最后一个节点，在开始处理时，我们还没有确认任何节点是无重复的，pre 指向 dummy
	*/
	dummy := &ListNode{Next: head}
	pre := dummy

	/*
	   主体逻辑部分：双指针策略划分已处理和未处理区域，在未处理区域检查重复
	   1. 重复：跳过所有重复节点
	   2. 未重复：正常移动 pre 指针
	*/
	for pre != nil && pre.Next != nil {
		// cur指向当前需要检查的节点
		cur := pre.Next

		// 检查当前节点是否存在重复，即是否与下一个节点值相同
		if cur.Next != nil && cur.Val == cur.Next.Val {
			// 明确有重复，开始循环检查
			// 不断向后移动cur，直到遇到不同值的节点或链表末尾
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}

			// 将pre的Next指向重复节点之后的不同节点，这样就直接跳过了所有重复的节点
			// pre 不用后移，因为重复节点已删除，后续的节点依然是未处理节点
			pre.Next = cur.Next
		} else {
			// 如果当前节点没有重复，正常处理：
			// 1. 确保pre.Next指向当前节点（这步其实可以省略，因为本来就是指向cur）
			// 2. 将pre指针移动到当前节点，继续处理下一个节点
			pre = cur
		}
	}
	return dummy.Next
}
