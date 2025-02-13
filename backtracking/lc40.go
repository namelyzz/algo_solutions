package backtracking

import "slices"

/*
我觉得这道题比起 39 难了一些，要多读一读题目，总结出题目的核心约束条件：
1. candidates 是包含了重复的元素
2. 每个元素在每个组合里只能使用一次，这和 39 题可以使用多次不同，也就是说这次 dfs 要传入 i+1
（以下和 39 的约束一样）
 3. 最终结果不能包含重复的组合，如 [1,2] 和 [1,2] 需去重
    虽然这道题的约束和 39 一样，但是要注意约束 1，40 题是有重复元素的
    我们当然需要在递归的过程中（for 循环主题内）去重，那么怎么去重呢？这是需要思考的问题

4.组合内元素无顺序。这倒是没什么影响，和 39 一样先排好序即可
*/
func combinationSum2(candidates []int, target int) [][]int {
	// 和 lc39 一样，先排好序，利用有序性剪枝
	// 还有一个好处，就是能让重复的元素相邻，为去重提供一个便利条件
	slices.Sort(candidates)
	n := len(candidates)
	var res [][]int
	path := []int{}

	var dfs func(int, int)
	dfs = func(i, val int) {
		// 递归终止条件：当剩余目标值为 0 时，说明当前 path 的和等于 target，是有效组合
		// 和 39 一样也可以用累加上去的方式计算终止条件，看想起什么用什么吧
		if val == 0 {
			res = append(res, slices.Clone(path))
			return
		}

		// 从起始索引 i 开始，避免重复使用之前的元素
		// 剪枝条件 candidates[j] <= val
		// 因为已经对数组排序了，如果后续的元素更大，必然无效，无需再遍历
		for j := i; j < n && candidates[j] <= val; j++ {
			// 本题最关键的地方，如果当前元素于前一个元素相同，并且不是当前
			// 怎么去重，条件解读：
			// 1. j > i: 确保是“同一层”（同一递归深度）的重复元素
			// 2. candidates[j] == candidates[j-1] 元素是相同的
			// 如果还是不明白这个去重逻辑，我在代码下面给出更具体的去重逻辑
			if j > i && candidates[j] == candidates[j-1] {
				continue
			}

			// 常规的回溯逻辑，记得 +1 即可
			path = append(path, candidates[j])
			dfs(j+1, val-candidates[j])
			path = path[:len(path)-1]
		}
	}

	dfs(0, target)
	return res
}

/*
去重逻辑，先从一个例子出发来看看
假设 candidates（已排好序） = [1, 1, 2, 5], target = 8

开始: dfs(0, 8)
├── 选择第一个1: path=[1], val=7
│   ├── 选择第二个1: path=[1,1], val=6
│   │   ├── 选择2: path=[1,1,2], val=4
│   │   │   └── 选择5: path=[1,1,2,5], val=-1 ✗
│   │   └── 选择5: path=[1,1,5], val=1 ✗
│   ├── 选择2: path=[1,2], val=5
│   │   └── 选择5: path=[1,2,5], val=0 ✓ [1,2,5]
│   └── 选择5: path=[1,5], val=2 ✗
├── 选择第二个1: 跳过! (因为 j>0 且 candidates[1]==candidates[0])
├── 选择2: path=[2], val=6
│   ├── 选择5: path=[2,5], val=1 ✗
│   └── 无更多选择
└── 选择5: path=[5], val=3 ✗

题解中，每个j值代表同一层级的不同选择，在回溯算法中，同一层也就是在同一递归深度、同一循环迭代中的选择

*/
