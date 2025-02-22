package linked_list

/*
直观写法，比较符合一般思路
因为链表有序，依次对比，找到小的放到答案链表中
最后将还有剩余的链表接入答案链表中即可
这是一般很快能想到的思路，写起来也很快，不多做解释了
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}

/*
递归，如果理解了之后写起来很方便
当然了空间复杂度更高，因为递归需要使用到栈
思路：递归地比较两个链表头节点，选择较小的作为当前节点，然后递归合并剩余部分
*/
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	// 基本情况，也是递归终止条件：
	if list1 == nil {
		// list1为空，直接返回list2剩余部分
		return list2
	} else if list2 == nil {
		// list2为空，直接返回list1剩余部分
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
