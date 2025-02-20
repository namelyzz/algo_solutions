package linked_list

/*
快慢指针找到中点
反转中点后面的部分，然后判断两个链表是否一致，一致就是回文
*/
func isPalindrome(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    list := reverse(slow)
    for head != slow {
        if head.Val != list.Val {
            return false
        }

        head = head.Next
        list = list.Next
    }
    return true
}

func reverse(head *ListNode) *ListNode {
    var pre *ListNode
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    return pre
}
