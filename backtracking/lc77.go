package backtracking

import "slices"

func combine(n int, k int) [][]int {
	var res [][]int
	path := make([]int, 0, k)

	var dfs func(int)
	// start 参数表示在递归过程中，元素选择的起始位置
	// 这与循环主体中的条件共同构成了组合的剪枝
	dfs = func(start int) {
		// 递归终止条件：当当前路径的长度等于 k 时，说明找到了一个有效组合
		if len(path) == k {
			// 注意：必须克隆当前路径再加入结果集
			// 原因：切片是引用类型，直接 append(path) 会导致后续修改影响已存入的结果
			res = append(res, slices.Clone(path))
			return
		}

		// 遍历选择元素，自顶向下，构建组合
		// 保证每次选择的元素都比上一个大，避免出现 [2,1] 这种与 [1,2] 重复的组合
		// 本题解虽然没有显式的剪枝语句，但是实际上在这里，配合 start，已经实现了一种隐式的剪枝
		// 因为它确保了数字按顺序选择，只会当前元素及之后的元素，不会选之前已经组合过的元素，避免了重复组合
		// 灵神的题解，自底向上的解法：
		// https://leetcode.cn/problems/combinations/solutions/2071017/hui-su-bu-hui-xie-tao-lu-zai-ci-pythonja-65lh
		// 就有显示的剪枝，d = k - len(path), j >= d
		// 这个自顶向下的题解，如果写出显示的剪枝语句，就是 i <= n - (k - len(path)) + 1
		for i := start; i <= n; i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	// 从数字 1 开始启动递归
	dfs(1)
	return res
}
