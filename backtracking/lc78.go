package backtracking

import "slices"

func subsets(nums []int) [][]int {
    n := len(nums)

    // 初始化结果集res，预分配容量为2^n（因为n个元素的子集总数是2^n）
    // 比较套路的回溯模板，res 表示最终结果，path 记录当前状态
    res := make([][]int, 0, 1<<n)
    path := make([]int, 0, n)

    // 定义深度优先搜索函数dfs，参数i表示当前处理的元素索引，从第i个元素开始，决定是否将其加入当前子集
    var dfs func(int)
    dfs = func(i int) {
        // 递归终止条件：当处理完所有元素（i等于数组长度n）
        if i == n {
            // 使用slices.Clone确保存入的是独立副本，避免后续修改影响结果
            res = append(res, slices.Clone(path))
            return
        }

        // 【不选择当前元素】的分支：直接递归处理下一个元素
        dfs(i + 1)

        // 【选择当前元素】的分支：
        // 1. 将当前元素加入路径
        path = append(path, nums[i])
        // 2. 递归处理下一个元素，此时path包含nums[i]
        dfs(i + 1)
        // 3. 回溯操作：移除路径中最后一个元素（即刚才加入的nums[i]），恢复path到选择当前元素之前的状态
        path = path[:len(path)-1]
    }

    dfs(0)
    return res
}
