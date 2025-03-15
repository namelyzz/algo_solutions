package tree

/*
同剑指offer26题

这道题要求判断树B是否是树A的子结构。
注意子结构与子树的区别：子结构不要求包含所有后代节点，只要B的结构和节点值与A的某一部分完全匹配即可。

思路：
1. 看 A, B 根节点是否相同
  - 相同：看 A.Left & B.Left 以及 A.Right & B.Right 是否相同
  - 不同，B 是否存在于 A 的左子树或右子树中

2. 因此，需要两种递归，对应上面两种情况
- 相同：
  - b 为空，则不用再比对，直接返回 true
  - a 为空，说明对不上，直接返回 false
  - a, b 都不为空，比较值是否相同
  - 根节点比对完毕后，比对左右子树

- 不同，则调用原函数 isSubStructure 从 A 的左右子树去找 B 的结构
*/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(a, b *TreeNode) bool {
		// 要注意顺序，先判断
		// 这个判断必须在a为nil之前，因为可能a还有节点但b已匹配完
		if b == nil {
			return true
		}
		if a == nil {
			return false
		}

		if a.Val != b.Val {
			return false
		}

		return dfs(a.Left, b.Left) && dfs(a.Right, b.Right)
	}

	return dfs(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
