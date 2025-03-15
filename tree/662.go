package tree

/*
思路：
先说说第一次做的错误思路：
首先看到题目，很容易想到层序遍历（102题），然后计算按题目要求，计算多少个节点
然后找到第一个非nil节点和最后一个非nil节点，计算个数
这样逐层计算，找到最大值

想法是美好的，但是结果是超出限制的
这个方法需要逐层计算出结果，遍历，时间复杂度高，此外，还存了一些不必要的 nil 节点，空间复杂度也不理想
那么有没有一种方法能够快速计算出个数呢？
如果说，我们只存有效节点，并且计算和存储有效节点的索引，这样只需要知道队列中第一个节点和最后一个节点的索引，就能轻松算出这一层的宽度了
*/

type pair struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	queue := []pair{{root, 0}}
	var res int
	for len(queue) != 0 {
		res = max(res, queue[len(queue)-1].index-queue[0].index+1)
		tmp := []pair{}
		for _, p := range queue {
			if p.node.Left != nil {
				tmp = append(tmp, pair{p.node.Left, p.index * 2})
			}
			if p.node.Right != nil {
				tmp = append(tmp, pair{p.node.Right, p.index*2 + 1})
			}
		}
		queue = tmp
	}
	return res
}
