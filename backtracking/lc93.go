package backtracking

import (
	"strconv"
	"strings"
)

// restoreIpAddresses 这道题目，考的频率挺高的，而且本身要处理的情况也多
// 本身还考察了面试者对字符串的处理，多种条件很容易遗漏或者犯错，需要多多复习
func restoreIpAddresses(s string) []string {
	n := len(s)
	res := make([]string, 0, 1<<n)
	// 已分割出来的 ip 段，题目要求必须刚好 4 段
	path := make([]string, 0, 4)

	var dfs func(int)
	dfs = func(start int) {
		// 递归终止条件：已经处理到边界了，并且 ip 段刚好为 4 个，可以把答案加入 res 了
		if start == n && len(path) == 4 {
			res = append(res, strings.Join(path, "."))
			return
		}

		// 非法条件1：还未到达边界，已经处理出来 4 个 ip 段了，这条路径是错误的
		if start < n && len(path) == 4 {
			return
		}

		// 一个 ip 段长度最长为 3, 我们从当前位置(start)开始，逐个构建可能的 ip 段
		for i := 0; i < 3; i++ {
			// 非法条件2：前导零，说明这条路径错误的
			if i != 0 && s[start] == '0' {
				return
			}

			// 非法条件3：越界
			if start+i >= n {
				return
			}

			str := s[start : start+i+1]
			// 非法条件4：将子串转换为整数，若大于 255，超出 IP 段范围（0-255），无效
			if val, _ := strconv.Atoi(str); val > 255 {
				return
			}

			path = append(path, str)
			dfs(start + i + 1) // 递归处理下一段，起始索引为当前段的结束索引 +1
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return res
}
