package backtracking

import "slices"

func partition(s string) [][]string {
	n := len(s)
	res := make([][]string, 0, 1<<n)
	path := make([]string, 0, n)

	// 需要知道当前子串的初始位置和结束位置，然后再取出这个子串判断是否为回文
	// i: 当前处理到的字符索引
	// start: 当前正在构建的子串的起始索引
	var dfs func(int, int)
	dfs = func(i, start int) {
		// 递归终止条件：当 i 等于字符串长度 n 时，说明已处理完所有字符
		if i == n {
			res = append(res, slices.Clone(path))
			return
		}

		// 【不选】，不分割当前位置，继续向后扩展子串
		// 注意，右边界需要小于 n-1, 这样才能考虑拆分的情况
		if i < n-1 {
			dfs(i+1, start) // 起始索引不变，递归处理下一个字符
		}

		// 【选】，当前子串是回文吗？如果是，才继续递归，不是就没有递归的必要了
		subStr := s[start : i+1]
		if isPartition(subStr) {
			path = append(path, subStr)
			dfs(i+1, i+1)             // 当前子串已经确认并加入 path，从新位置 i+1 开始，start 也移动到这个位置
			path = path[:len(path)-1] // 回溯：移除最后加入的子串，尝试其他分割方式
		}
	}

	dfs(0, 0)
	return res
}

func isPartition(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}
	return true
}
