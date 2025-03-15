package tree

/*
思路：
问题的核心是交换每个节点的左右子树，这个问题的子问题也是一致的，因此可以很快想到使用递归
解题过程就是：
如果该节点有左子树，翻转后应该成为右子树
如果该节点有右子树，翻转后应该成为左子树
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left := invertTree(root.Right)
	right := invertTree(root.Left)
	root.Left = left
	root.Right = right
	return root
}
