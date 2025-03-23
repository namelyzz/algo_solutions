package two_pointer

/*
怎么算储水量？两个因素
短板高度：两条垂线中较矮的那条
水平距离：两条垂线在 x 轴上的距离（即两个索引的差值）
得到储水量公式：当前储水量 = min(左边垂线高度, 右边垂线高度) × (右边索引 - 左边索引)

有了上面的公式，求解思路易得：利用双指针缩小搜索范围，结合贪心策略保证每次移动都有机会找到更大的储水量。
每次计算当前双指针对应的储水量后，移动指向短板的指针
为什么要移动短板？我们用反证法
假设当前左指针是短板（height[left] < height[right]）
若移动长板（right--）：

	新的水平距离会减小（right-left 减 1），
	而新的短板要么还是 height[left]（若新的 right 高度 ≥ height[left]），
	要么是更短的 height[new right]（若新的 right 高度 < height[left]）。
	无论哪种情况，新的储水量一定小于当前储水量。

反之，若移动短板（left++）：

	虽然水平距离减小，但可能找到更高的左指针高度，使得新的短板高度提升，最终储水量可能增大。

因此，移动短板是唯一可能找到更大储水量的选择，这就是贪心策略的体现。
*/
func maxArea(height []int) int {
	res, left, right := 0, 0, len(height)-1
	for left < right {
		res = max(res, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
