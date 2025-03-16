package tree

/*
思路：
先考虑一下最近公共祖先（LCA）有哪种情况：
1. 左右各一：p, q 分别位于左右子树中，那么 LCA 就是 root
2. 根即为 p：root 本身是 p，且 q 在 root 的子树中
3. 根即为 q: root 本身是 q，且 p 在 root 的子树中

怎么写递归函数？
写递归时，不要试图模拟压栈过程，先明确函数的定义
lowestCommonAncestor 函数本质上是：在以 root 为根的子树中，去寻找 p 和 q。
- 如果找到了 p 或 q 中的任意一个，就返回找到的那个节点。
- 如果两个都找到了（即 root 是 LCA），就返回 LCA。
- 如果都没找到，返回 nil。

首先处理最简单的情况，也就是递归什么时候该停下来，此时，我们只需要关注 root 节点
如果我们搜到了叶子节点的下面（空节点），说明这条路走不通，既没找到 p 也没找到 q。

其次，如果当前遍历到的节点 root 恰好就是我们要找的 p 或者 q，还需要继续向下找吗？
不需要了。比如我们找到了 p，根据 LCA 的定义，无论 q 在不在 p 的下面，p 都有可能是最近公共祖先
所以，一旦遇到目标节点，直接向上传递当前节点（告诉父节点“我找到了一个”）。

root 节点并非 p 或 q，我们不知道 p 和 q 在哪，那就分别去左子树和右子树里找。
根据左右子树的反馈，判断当前 root 的身份。

1. 左右都找到了
left 不为空（说明左边找到了一个），right 也不为空（说明右边找到了一个）。
这就意味着 p 和 q 分家了，一个在我的左边，一个在我的右边。
那我（当前 root）肯定就是那个分岔口，也就是最近公共祖先。直接返回自己。

2. 只有一边找到了
如果 left != nil 但 right == nil：说明 p 和 q 都在左边
同理，如果 left == nil 但 right != nil，说明都在右边
如果不满足“左右双全”，那就是谁有值返回谁。因为我们是自顶向下递归，会先遇到 LCA

时间复杂度：O(N)。最坏情况下我们需要遍历二叉树的所有节点（N 为节点总数）。
空间复杂度：O(N)。最坏情况下二叉树退化成链表，递归栈的深度为 N。平均情况下是 O(logN)。
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 递归终止情况：如果我们搜到了叶子节点的下面（空节点），说明这条路走不通，既没找到 p 也没找到 q。
	if root == nil {
		return nil
	}

	// 递归终止情况：找到其中一个节点
	if root == p || root == q {
		return root
	}

	// 我们不知道 p 和 q 在哪，那就分别去左子树和右子树里找。
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// left 不为空（说明左边找到了一个），right 也不为空（说明右边找到了一个）。
	// 这就意味着 p 和 q 分家了，一个在我的左边，一个在我的右边。 root 就是它们的 LCA
	if left != nil && right != nil {
		return root
	}

	// p, q在同一侧，非空的一侧就是 LCA
	if left != nil {
		return left
	}
	return right
}
