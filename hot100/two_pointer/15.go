package two_pointer

import "slices"

/*
三数之和
https://leetcode.cn/problems/3sum/description/?envType=study-plan-v2&envId=top-100-liked

给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。

请你返回所有和为 0 且不重复的三元组。

示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。

示例 2：
输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。

示例 3：
输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。
*/

func threeSum(nums []int) (res [][]int) {
	slices.Sort(nums)

	for i, n := 0, len(nums); i < n-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if nums[i]+nums[n-2]+nums[n-1] < 0 {
			continue
		}

		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}

/*
解题思路
排序：将数组排序后，相同元素相邻，便于快速跳过重复；
双指针：对每个固定元素i，将问题转化为在子数组[i+1, n-1]中找两数之和为 -nums[i] 的问题；
剪枝优化：利用排序后的有序性提前终止无效搜索，减少计算量。

为什么排序是必须的？
- 排序后，相同元素连续出现（如[-1,-1,0,1,2]），可通过比较相邻元素快速跳过重复值
- 方便我们使用双指针思路
为什么用双指针思路？
对于固定i，目标是找j和k（j>i, k>j）使得nums[j] + nums[k] = -nums[i]。
排序后，子数组[i+1, n-1]是升序的，因此：
	若当前和 < -nums[i]，需增大和 → j++（向右移动j，取更大值）；
	若当前和 > -nums[i]，需减小和 → k--（向左移动k，取更小值）；
	若相等，记录结果并移动j和k跳过重复值。
剪枝优化
	条件：nums[i] + nums[i+1] + nums[i+2] > 0
	做法：提前 break
	说明排序后，nums[i] ≤ nums[i+1] ≤ nums[i+2]，若三者和已>0，则后续所有组合（i更大）的和必然>0（因为所有数更大）。
	条件：nums[i] + nums[n-2] + nums[n-1] < 0
	做法：提前 continue
	排序后，nums[i] ≤ nums[n-2] ≤ nums[n-1]，若最小可能的和（nums[i] + 最后两个数）仍<0，则当前i太小，无法通过增大j/k使和≥0。

算法步骤：
排序数组：将nums按升序排列（确保后续逻辑正确）。
遍历每个可能的i（第一个数）：
	a. 跳过重复i：若i>0且nums[i] == nums[i-1]，跳过当前i（避免重复三元组）。
	b. 剪枝1：若nums[i] + nums[i+1] + nums[i+2] > 0，break（后续和必然>0）。
	c. 剪枝2：若nums[i] + nums[n-2] + nums[n-1] < 0，continue（当前i太小，和不可能为0）。
双指针j和k初始化：j = i+1, k = n-1。
在j < k范围内循环：
	a. 计算当前和：sum = nums[i] + nums[j] + nums[k]。
	b. 若sum == 0：
		添加三元组[nums[i], nums[j], nums[k]]到结果。
		移动j：j++直到nums[j] != nums[j-1]（跳过重复j值）。
		移动k：k--直到nums[k] != nums[k+1]（跳过重复k值）。
	c. 若sum < 0：j++（需增大和，向右移动j）。
	d. 若sum > 0：k--（需减小和，向左移动k）。
返回结果：所有唯一三元组组成的二维切片。
*/
