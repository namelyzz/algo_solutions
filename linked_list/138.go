package linked_list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
思路：
深拷贝的意思是：创建一个新的链表，新链表的每个节点都是新创建的，而不是原链表的引用。
这道题的关键难点在于：
链表中除了普通的 next 指针，还有一个 random 指针可能指向链表中的任意节点（包括 null）。
如果简单遍历复制，新链表的 random 指针会指向原链表的节点，而不是新链表的对应节点。

使用哈希表建立原节点到新节点的映射关系
为什么想到用哈希表？
1. 单独深拷贝出来一个节点，此时我们还没有设置它的 Next 和 Random，而 Next 和 Random 也是深拷贝出来的节点
2. 换句话说，我们深拷贝出来的节点，它们直接的连接关系还未构建
3. 想知道它们之间的连接关系，就需要依照原链表将它们重新构建
4. 所以，我们需要在复制过程中保持原节点与新节点的对应关系。哈希表就是一个很好的选择，帮我们快速找到新节点对应的原节点，从而得到它的 Next 和 Random 节点

具体步骤：
1. 第一次遍历创建所有新节点并建立映射
2. 第二次遍历通过映射关系设置新节点的next和random指针
*/

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	nodeMp := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		nodeMp[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		newNode := nodeMp[cur]

		if cur.Next != nil {
			newNode.Next = nodeMp[cur.Next]
		}
		if cur.Random != nil {
			newNode.Random = nodeMp[cur.Random]
		}

		cur = cur.Next
	}

	return nodeMp[head]
}
