package tree

/*
本质上来说就是 Nil 节点后不能再出现非空节点了
*/

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	hasNil := false

	for len(queue) > 0 {
		var tmp []*TreeNode
		for _, node := range queue {
			if node == nil {
				hasNil = true
			} else {
				if hasNil {
					return false
				}

				tmp = append(tmp, node.Left)
				tmp = append(tmp, node.Right)
			}
		}
		queue = tmp
	}

	return true
}
