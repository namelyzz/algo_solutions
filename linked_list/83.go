package linked_list

/*
思路：
这道题要求删除排序链表中的重复元素，使得每个元素在链表中只出现一次。
由于链表是已排序的，重复元素一定是相邻的，因此可以通过一次遍历来删除重复节点。
只需要遍历链表，并在当前节点和下一个节点的值相等时，删除下一个节点。相比于82题要重复就全删，83题要简单得多
*/
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    dummy := &ListNode{Next: head}
    cur := head
    for cur.Next != nil {
        if cur.Val == cur.Next.Val {
            cur.Next = cur.Next.Next
            continue
        }
        cur = cur.Next
    }
    return dummy.Next
}
