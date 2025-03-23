package two_pointer

/*
先理解题目要求：nums1 的前 m 个元素是有效的，后面有 m 个空位来存放 nums2 中的元素
思路是逆向合并（即从后往前填数据）
- 因为 nums1 已经有足够的空间，所以我们从 nums1 的最后一位开始填充合并后的元素。
-通过从后往前填充，可以避免覆盖 nums1 中的有效元素，因为我们只会修改那些空位置（即 nums1 后面 m 个空位置）。

可以使用双指针，p1 和 p2 分别指向 nums1 和 nums2 的末尾有效元素，比较哪个元素应该放到 nums1 末尾

需要注意处理剩余元素，如果 nums2 没有要处理的元素，nums1 本来就作为答案，可以不用处理；
如果 nums1 中的元素已经处理完（即 p1 < 0），那么剩下的元素必定在 nums2 中，因为 nums2 可能还有未处理的元素。
此时直接将 nums2 剩下的元素复制到 nums1 中。
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	cur, p1, p2 := len(nums1)-1, m-1, n-1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] < nums2[p2] {
			nums1[cur] = nums2[p2]
			p2--
		} else {
			nums1[cur] = nums1[p1]
			p1--
		}
		cur--
	}

	for p2 >= 0 {
		nums1[cur] = nums2[p2]
		p2--
		cur--
	}
}
