package backtracking

func isAdditiveNumber(num string) bool {
	n := len(num)
	// 题目要求：一个有效的 累加序列 必须 至少 包含 3 个数。
	// 2 个及以下长度不满足题意，直接返回 false 即可
	if n <= 2 {
		return false
	}

	// 最终结果 res 只需要一个 bool 变量存储，所以不用定义
	// 存储当前构建的累加数序列
	path := make([]int, 0, n)

	// 通过 dfs 尝试构建累加数序列
	// 解析一个子串，看是否能加入累加数序列并保证后续结果符合题意
	// left: 当前正在解析的数字字符串起始索引
	// right: 当前正在解析的数字字符串结束索引
	// sum: 当前这个数字字符串的值
	var dfs func(int, int, int) bool
	dfs = func(left, right, sum int) bool {
		// 递归终止条件：已解析完整个字符串
		if left == n {
			// 若序列长度至少为3，则是有效的累加数
			// 首先，累加数的定义本身就决定了序列长度至少为 3
			// 例如，"112358" 可以分割为 [1,1,2,3,5,8]，长度为 6（≥3）
			// 我们本质上就是计算出来一个可行的累加数序列
			// 所以，当已经遍历完整个字符串，而 path 长度大于 3 时即可判断已经有有效累加数序列了
			return len(path) >= 3
		}

		// 剪枝：前导零检查，当前数字长度>1且以'0'开头，无效
		if right-left+1 > 1 && num[left] == '0' {
			return false
		}

		sum = sum*10 + int(num[right]-'0')

		// 【不选】，延长当前数字字符串长度到下一位
		if right < n-1 {
			// 继续递归，看能否找到有效的累加数序列，如果最后找到了，返回 true
			if dfs(left, right+1, sum) {
				return true
			}
		}

		// 【选】，两种情况，满足其一既可以选：
		// 1. 序列中数字不足2个（说明此时还在构建前两个数字）
		// 2. 序列中已有至少2个数字，且最后两个数字之和等于当前sum，那说明刚好构成了有效累加数序列了
		if len(path) < 2 || path[len(path)-1]+path[len(path)-2] == sum {
			path = append(path, sum)
			// 当前子串已经加入 path，递归解析下一个数字，所以 left, right 从新子串重新出发
			if dfs(right+1, right+1, 0) {
				return true
			}
			path = path[:len(path)-1]
		}

		return false
	}

	return dfs(0, 0, 0)
}
