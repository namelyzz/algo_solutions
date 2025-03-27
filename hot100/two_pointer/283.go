package two_pointer

/*
https://leetcode.cn/problems/move-zeroes/description/?envType=study-plan-v2&envId=top-100-liked

移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:
输入: nums = [0]
输出: [0]
*/

func moveZeroes(nums []int) {
	slow := 0
	for n, fast := len(nums), 0; fast < n; fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}

/*
本问题要求将数组中的所有0移动到数组末尾，同时保持非零元素的相对顺序不变。
所以要实现原地操作（不使用额外空间）且时间复杂度为O(n)
- 慢指针（slow）：指向当前已处理的非零元素的下一个放置位置（即所有非零元素的末尾+1）。
- 快指针（fast）：遍历整个数组，扫描每个元素。
核心逻辑：当快指针遇到非零元素时，将其与慢指针位置交换，同时移动慢指针。
这样，慢指针之前的所有元素都是非零的（且顺序与原数组一致），
而慢指针之后的元素（包括当前快指针位置）会被后续非零元素逐步覆盖，最终0自然被推到末尾。
*/
