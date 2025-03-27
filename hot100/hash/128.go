package hash

/*
https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-100-liked

给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

示例 2：
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
*/

func longestConsecutive(nums []int) int {
	n := len(nums)
	mp := make(map[int]bool, n)
	for _, v := range nums {
		mp[v] = true
	}

	var res int
	for k := range mp {
		if _, ok := mp[k-1]; ok {
			continue
		}

		cnt := 1
		for mp[k+1] {
			cnt++
			k++
		}

		res = max(res, cnt)
	}

	return res
}

/*
解题思路：
本问题要求找出数组中最长连续整数序列的长度（例如输入[100,4,200,1,3,2]时，最长连续序列是[1,2,3,4]，长度为4）。
核心挑战在于避免暴力解法的O(n²)时间复杂度（对每个元素都尝试向两边扩展）
能想到用哈希表？
哈希表（map）提供O(1)的元素存在性检查，能快速判断k+1、k-1是否在数组中，这是优化的基础。
进而，关键突破点在于只从连续序列的起点开始计数，从而避免重复计算。
为什么只从序列起点开始？
每个连续序列只被处理一次（从起点开始计数），避免了对2、3、4重复计数

算法步骤：
创建哈希表mp，将nums中所有元素存入（键为元素值，值为true），用于O(1)查找。
初始化结果变量res为0。
遍历哈希表中每个键k（即数组中的每个元素）：
	a. 检查k-1是否在mp中：
	若存在（说明k不是序列起点），跳过当前k。
	若不存在（说明k是序列起点），执行以下操作：
		i. 初始化计数器cnt=1（当前元素k自身）。
		ii. 循环检查k+1是否在mp中：
		- 若存在，cnt++，k++（向右移动到下一个连续数字）。
		- 重复直到k+1不存在于mp中。
		iii. 更新res = max(res, cnt)。
返回res（最长连续序列长度）。
*/
