package two_pointer

func trap_solution1(height []int) int {
	n := len(height)

	left := make([]int, n)
	left[0] = height[0]
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i])
	}

	right := make([]int, n)
	right[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	var res int
	for i := range height {
		res += min(left[i], right[i]) - height[i]
	}
	return res
}

/*
相比于 solution1，solution2 也是一个双指针的解法，但是时间复杂度为 O(n)，空间复杂度为 O(1)。
不依靠 left 和 right 两个数组，而是通过 leftMax 和 rightMax 两个变量来记录当前指针位置的左边和右边的最大高度。
*/
func trap_solution2(height []int) int {
	n := len(height)
	left, right := 0, n-1
	var res, leftMax, rightMax int
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		if height[left] <= height[right] {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return res
}

/*
为什么移动较短边？
对于任意位置，接水量由左右两侧最大高度的较小值决定（即min(leftMax, rightMax) - height[i]）。
当height[left] <= height[right]时，left边是限制因素（因为leftMax ≤ rightMax，证明见下文），故当前left位置的接水量 = leftMax - height[left]。
当height[left] > height[right]时，right边是限制因素（因为rightMax ≤ leftMax），故当前right位置的接水量 = rightMax - height[right]。
数学证明：
	若height[left] <= height[right]，则leftMax <= rightMax（反证：若leftMax > rightMax，则存在索引i < left使得height[i] = leftMax > rightMax，但rightMax是[right, n-1]的最大值，而i < left < right，当右指针遍历到i时，rightMax应≥height[i]，矛盾）。
	同理，若height[left] > height[right]，则rightMax <= leftMax。

为什么不会遗漏水量？
每次移动指针时，当前指针位置的接水量被精确计算（因为限制高度已确定）。
例如输入[2,1,0,1,2]：
	初始：left=0 (2), right=4 (2) → height[0]=2 <= height[4]=2，res += max(0,2)-2=0，left=1。
	left=1 (1), right=4 (2) → height[1]=1 <= 2，res += max(2,1)-1=1，left=2。
	left=2 (0), right=4 (2) → res += 2-0=2，left=3。
	left=3 (1), right=4 (2) → res += 2-1=1，left=4 → 结束，res=4（正确）。
*/
