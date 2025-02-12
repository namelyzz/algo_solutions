package backtracking

import "slices"

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	path := make([]int, 0, k)

	var dfs func(int, int)
	// start：当前递归中数字选择的起始位置（用于去重，确保数字递增）
	// sum：当前路径中所有数字的和（用于快速判断是否符合目标）
	dfs = func(start, sum int) {
		// 递归终止条件：路径长度为 k 且当前和等于 n，同时满足这两个条件才是有效组合
		if len(path) == k && sum == n {
			res = append(res, slices.Clone(path))
			return
		}

		// i 从 start 开始：保证数字严格递增，避免重复组合（如 [1,2] 和 [2,1] 只生成前者）
		// 题目要求：范围限制在 1-9，到时候 dfs 的 start 从 1 开始
		for i := start; i <= 9; i++ {
			// 剪枝优化：若当前和加上 i 已超过目标 n，后续数字更大，必然超，直接终止循环
			if sum+i > n {
				break
			}

			path = append(path, i)
			dfs(i+1, sum+i) // 递归深入：下一个数字必须从 i+1 开始（避免重复使用同一数字），同时累加和
			path = path[:len(path)-1]
		}
	}

	// 从数字 1 开始递归，初始和为 0
	dfs(1, 0)
	return res
}
