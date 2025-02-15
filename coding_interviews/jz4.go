package coding_interviews

import "sort"

/*
JZ4 二维数组中的查找

在一个二维数组array中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
[
[1,2,8,9],
[2,4,9,12],
[4,7,10,13],
[6,8,11,15]
]
给定 target = 7，返回 true。
给定 target = 3，返回 false。
*/

/*
思路：题目中关键信息——每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序
- target < 最左元素 —— 之后的元素都比 target 大，不可能找到
- target > 最右元素 —— 往下一层寻找
- 剩余最后一种情况，target > 最左元素 && target < 最右元素，target 在这一行
	- 又因为元素都是有序的，可以用二分查找，以 0(logn) 的复杂度找到目标元素
	- 用二分查找找到索引后，还得记得判断一下是否等于 target，因为 target 可能不存在与 array 中
*/

func Find(target int, array [][]int) bool {
	n := len(array[0])
	if n == 0 {
		return false
	}

	for _, nums := range array {
		if nums[0] > target {
			return false
		} else if nums[n-1] < target {
			continue
		} else {
			idx := sort.SearchInts(nums, target)
			if idx < n && nums[idx] == target {
				return true
			}
		}
	}
	return false
}
