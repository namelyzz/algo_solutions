package tree

func isSymmetric(root *TreeNode) bool {
	// 判断是否轴对称，只有一个根节点就一定是对称的了，所以从他的左右子树开始判断
	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(p, q *TreeNode) bool {
		// 如果其中有一棵子树已经空了
		// 除非两棵子树都为空，否者不对称
		if p == nil || q == nil {
			return p == q
		}

		// 两棵子树的值不同，那么不对称
		if p.Val != q.Val {
			return false
		}

		// 根据轴对称对比（就是镜像）
		// 左子树的左子树对比右子树的右子树
		// 左子树的右子树对比右子树的左子树
		return dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
	}

	// 判断根节点的左右子树
	return dfs(root.Left, root.Right)
}
