package backtracking

import "slices"

func permute(nums []int) [][]int {
    n := len(nums)
    // 相较于子集型和组合型的回溯，排列型的回溯可以选择前面的元素
    // 所以，除了当前已经加入 path 的元素，我们还需要知道哪些元素可以被选
    // selects 的作用就是用来记录，当前已经那些元素被选择了（true 表示已经选在 path 里了）
    selects := make([]bool, n)
    var res [][]int
    path := make([]int, n)

    var dfs func(int)
    dfs = func(i int) {
        // 递归终止条件：当已经检查完所有元素了，这时已经生成了一组排列了
        if i == n {
            res = append(res, slices.Clone(path))
            return
        }

        // 遍历 selects 找到未选择的元素
        for idx, ok := range selects {
            // 当前的元素没被选入当前的 path，则加入，标记位改为 true
            // 回溯时，把标记为重新设置为 false，表示还没被选过
            if !ok {
                path[i] = nums[idx]
                selects[idx] = true
                dfs(i+1)
                selects[idx] = false
            }
        }
    }

    dfs(0)
    return res
}