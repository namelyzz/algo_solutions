package hash

/*
https://leetcode.cn/problems/two-sum/description/?envType=study-plan-v2&envId=top-100-liked

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案。



示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]


提示：

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案
*/

func twoSum(nums []int, target int) []int {
	n := len(nums)
	mp := make(map[int]int, n)
	for i, v := range nums {
		if idx, ok := mp[target-v]; ok {
			return []int{i, idx}
		}
		mp[v] = i
	}
	return []int{-1, -1}
}

/*
解题思路：
采用哈希表（map）实现 O(n) 时间复杂度的解法，避免暴力枚举的 O(n²) 复杂度。
核心思想：遍历数组时，对每个元素 v，检查 target - v 是否已在哈希表中。
若存在，则当前元素与哈希表中存储的元素构成解；若不存在，则将当前元素存入哈希表。

算法步骤:
初始化哈希表，用于存储元素值与对应索引
遍历数组中的每个元素 v 及其索引 i：
a. 计算目标差值 diff = target - v
b. 检查 diff 是否存在于哈希表中：
若存在，返回 [i, 哈希表中 diff 对应的索引]
若不存在，将当前元素 v 和其索引 i 存入哈希表
若遍历结束未找到解，返回 [-1, -1]
*/
