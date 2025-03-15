package tree

/*
也是 102 的扩展，要特殊处理的地方就是奇数层的顺序是从右到左（从第0层开始）
想知道第几层（是奇数层还是偶数层），看 res 的元素数量即可
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		n := len(queue)
		m := len(res)

		path := make([]int, n)
		tmp := make([]*TreeNode, 0, n*2)
		for i := 0; i < n; i++ {
			cur := queue[i]
			if m%2 == 0 {
				path[i] = cur.Val
			} else {
				path[n-i-1] = cur.Val
			}

			if cur.Left != nil {
				tmp = append(tmp, cur.Left)
			}
			if cur.Right != nil {
				tmp = append(tmp, cur.Right)
			}
		}
		queue = tmp
		res = append(res, path)
	}
	return res
}
