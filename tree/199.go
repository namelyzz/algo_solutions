package tree

/*
就是 102 层序遍历的扩展，每次拿出最后一位放入答案即可。思路详见 102 题解，不多赘述
*/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	var res []int
	for len(queue) != 0 {
		n := len(queue)
		res = append(res, queue[n-1].Val)

		tmp := make([]*TreeNode, 0, n*2)
		for i := 0; i < n; i++ {
			cur := queue[i]
			if cur.Left != nil {
				tmp = append(tmp, cur.Left)
			}
			if cur.Right != nil {
				tmp = append(tmp, cur.Right)
			}
		}
		queue = tmp
	}
	return res
}
