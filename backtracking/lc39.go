package backtracking

import "slices"

func combinationSum(candidates []int, target int) [][]int {
	// 排序预处理。为什么需要这一步？
	// 1. 我们后续步骤需要剪枝，数组排好序了，我们知道当前元素累加后超过 target 就可以不用再遍历后面的元素了
	// 2. 答案不能出现重复组合，也就是[3,2]这类与[2,3]重复的组合, 排序后就能保证组合的顺序性了
	slices.Sort(candidates)
	n := len(candidates)
	var res [][]int
	path := make([]int, 0, n)

	// start: 当前递归中选择元素的起始索引
	// 注意，本题允许重复选当前元素，所以 dfs 可以传当前索引 i 进去，如果题目要求不能重复选同一个元素，则 dfs 传入 i+1
	// sum 用来存储当前的累加和
	var dfs func(int, int)
	dfs = func(start, sum int) {
		// 递归终止条件：当累加和达到 target 时说明找到了一组可行的组合
		if sum == target {
			res = append(res, slices.Clone(path))
			return
		}

		// 遍历元素，从 start 开始，避免选择之前已经选过的元素
		// 基础剪枝：sum <= target，若当前总和已经超过了目标和，后续累加只会越来越大，所以无需再遍历了
		// 当然，提取出来这个条件，以一个 if 分支放在 for 循环里面，break 退出，这个写法更直观
		for i := start; i < n && sum <= target; i++ {
			path = append(path, candidates[i])
			// 注意，这里传入的是 i，而非 i+1！
			// 是“元素可重复选取”的核心实现
			// 然后 sum 开始累加
			dfs(i, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0)
	return res
}

/*
另一种写法，思路一样的，只是不断扣减，能刚刚好为 0 则加入 res
func combinationSum(candidates []int, target int) [][]int {
    slices.Sort(candidates)
    n := len(candidates)
    var res [][]int
    path := []int{}

    var dfs func(int, int)
    dfs = func(i, val int) {
        if val == 0 {
            res = append(res, slices.Clone(path))
            return
        }

        for j := i; j < n && candidates[j] <= val; j++ {
            path = append(path, candidates[j])
            dfs(j, val - candidates[j])
            path = path[:len(path)-1]
        }
    }

    dfs(0, target)
    return res
}
*/
