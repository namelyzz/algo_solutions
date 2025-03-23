package two_pointer

import (
	"math"
	"slices"
)

/*
相比于 15 题，这里我用一个结构体，是为了便于存储和比较三元组和与目标值之间的差异。
sum：代表当前三元组的和（nums[i] + nums[j] + nums[k]）。
diff：代表当前三元组和与目标值 target 之间的差值。diff = abs(sum - target)。
只要找到更小的 diff 就更新 sum 和 diff，最后返回 pair.sum 即可
*/
type pair struct {
	sum, diff int
}

func threeSumClosest(nums []int, target int) int {
	/*
	   排序是一个常见的技巧，用于简化问题
	   这里排序的目的是为了后续使用双指针方法来高效地查找最接近目标的三元组
	*/
	slices.Sort(nums)

	res := pair{0, math.MaxInt}
	n := len(nums)

	for i := 0; i < n-2; i++ {
		/*
		   和 15 题一样，剪枝方式一样，不一样的是要处理一下结果
		*/
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if s := nums[i] + nums[i+1] + nums[i+2]; s > target {
			diff := abs(s, target)
			if diff < res.diff {
				return s
			}
			return res.sum
		}

		if s := nums[i] + nums[n-2] + nums[n-1]; s < target {
			diff := abs(s, target)
			if diff < res.diff {
				res = pair{s, diff}
			}
			continue
		}

		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}

			diff := abs(sum, target)
			if diff < res.diff {
				res = pair{sum, diff}
			}

			if sum < target {
				for j++; j < k && nums[j-1] == nums[j]; j++ {
				}
			} else {
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				}
			}
		}
	}

	return res.sum
}

func abs(a, b int) int {
	val := a - b
	if val < 0 {
		return -val
	}
	return val
}
