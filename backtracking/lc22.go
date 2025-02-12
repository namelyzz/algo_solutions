package backtracking

/*
思路还是很好理解的，问题的本质就是生成 n 对有效括号，核心是保证两个约束：
- 总左括号数 = 总右括号数 = n
- 任何前缀中，左括号数 ≥ 右括号数（避免出现无法闭合的右括号）
那么回溯的设计思路就是：
- 状态表示：用left和right记录当前已使用的左右括号数量
- 分支选择：每次递归有两个可能选择（添加左括号或右括号），但最终组合受有效性约束限制，从这里出发，考虑什么时候添加左括号或右括号
- 终止条件：当右括号数达到 n 时，必然已生成有效组合
注意路径的构建，使用固定长度的 byte 切片path存储当前序列，通过left+right计算当前位置（已用括号总数 = 左 + 右）
用字符串拼接存在频繁的内存分配，效率低。
而且从刷题经验来看，当最终结果的长度是固定且可预知的, path 存储 byte 时, 往往用 make([]byte, cap) 而不是 make([]byte, 0, cap)
回溯的核心是 “尝试选择→递归深入→撤销选择”，而byte切片的索引替换方式天然适配这种逻辑：
- 固定长度的切片中，每个位置的索引可以通过当前递归的状态变量计算得出
- 递归深入时，直接在对应索引位置写入答案，回溯时，无需像动态切片那样执行path = path[:len(path)-1]的截断操作
- 这两种操作模式，容易搞混，一时没转过脑经写太快了会混用然后出错，这里做个小提醒
*/
func generateParenthesis(n int) []string {
	res := []string{}
	path := make([]byte, n*2)

	// 思路很好理解，就是选择填入左括号还是右括号
	// left：当前已使用的左括号数量
	// right：当前已使用的右括号数量
	var dfs func(int, int)
	dfs = func(left, right int) {
		// 递归终止条件：右括号数量达到n（此时左括号必然也为n，因有效性约束）
		if right == n {
			res = append(res, string(path))
			return
		}

		// 尝试添加左括号的条件：
		// 1. left < n：左括号数量未达上限（最多n个）
		// 2. left >= right：保证添加后仍满足"任意前缀左括号>=右括号"（有效性约束，不过这个条件去掉也没问题，
		// 因为代码顺序，left 比 right 先判断，因此 left 天然 >= right ）
		if left < n && left >= right {
			// 在当前位置（已用括号数=left+right）放置左括号
			path[left+right] = '('
			dfs(left+1, right)
		}

		// 尝试添加右括号的条件：
		// right < left：右括号数量小于左括号，保证添加后仍满足有效性（避免出现"))("等无效情况）
		if right < left {
			path[left+right] = ')'
			dfs(left, right+1)
		}
	}

	dfs(0, 0)
	return res
}
