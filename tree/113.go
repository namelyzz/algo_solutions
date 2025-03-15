package tree

import "slices"

/*
这道题目还是很好判断使用回溯算法的

边界情况：
- 空树直接返回空结果

DFS 递归：
- 维护当前路径和当前路径节点值列表
- 在递归过程中累加节点值
- 记录经过的节点

终止条件：
- 到达叶子节点时检查路径和是否等于目标值
- 如果相等，保存当前路径的副本

回溯处理：
- 在递归返回前从路径中移除当前节点，让路径状态上一次的状态

更详细的思路过程，看代码注释。这里记录我第一次递归错误的思路
```go

	dfs = func(node *TreeNode, sum int) {
	    if node == nil && sum == targetSum {
	        res = append(res, slices.Clone(path))
	        return
	    }
	    if node == nil || sum > targetSum {
	        return
	    }

	    sum += node.Val
	    path = append(path, node.Val)
	    dfs(node.Left, sum)
	    dfs(node.Right, sum)
	    path = path[:len(path)-1]
	}

```

这个递归体虽然得到了正确答案，但是是答案中存在重复的路径
1. 叶子节点判断错误
当节点为叶子节点时，我们检查的条件是if node == nil && sum == targetSum，
这意味着我们会在节点的左子节点和右子节点都为nil时，分别检查两次
因为会分别递归调用左子和右子，它们都是nil。
在 node == nil 时才检查是否满足条件，这会导致同一个叶子节点被重复记录两次
2. 路径和计算错误
这是个隐性的错误，在递归到nil节点时检查 sum == targetSum，但此时sum并没有包含当前nil节点的值
看起来合理，但是是无效操作

如何去修改的：
1. 正确的叶子节点判断：只有当 node.Left == nil && node.Right == nil 时才是叶子节点
2. 正确的时机检查路径和：在确定是叶子节点且当前sum等于targetSum时才记录路径
*/
func pathSum(root *TreeNode, targetSum int) [][]int {
	// 边界情况：空树直接返回空结果
	if root == nil {
		return nil
	}

	var res [][]int
	var path []int

	// 定义DFS递归函数
	// node: 当前节点
	// sum: 从根节点到当前节点的路径和
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		// 基准情况：空节点直接返回
		if node == nil {
			return
		}

		// 处理当前节点：累加节点值并记录路径
		sum += node.Val
		path = append(path, node.Val)

		// 检查是否到达叶子节点且路径和等于目标值
		if node.Left == nil && node.Right == nil && sum == targetSum {
			res = append(res, slices.Clone(path))
		}

		dfs(node.Left, sum)
		dfs(node.Right, sum)
		path = path[:len(path)-1]
	}

	// 从根节点开始DFS遍历，初始路径和为0
	dfs(root, 0)
	return res
}
