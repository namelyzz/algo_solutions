package backtracking

var mp map[byte][]byte = map[byte][]byte{
    '2': []byte{'a', 'b', 'c'},
    '3': []byte{'d', 'e', 'f'},
    '4': []byte{'g', 'h', 'i'},
    '5': []byte{'j', 'k', 'l'},
    '6': []byte{'m', 'n', 'o'},
    '7': []byte{'p', 'q', 'r', 's'},
    '8': []byte{'t', 'u', 'v'},
    '9': []byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
    n := len(digits)

    // 边界条件：如果输入为空字符串，直接返回空列表
    // 值得注意这个边界条件要处理，因为后续的代码，path 为空会生成空字符串，会返回 [""] 的 slice
    // 而这种情况下，应该生成的是 []
    if n == 0 {
        return []string{}
    }

    res := make([]string, 0, 1<<n)
    path := make([]byte, n)

    var dfs func(i int)
    dfs = func(i int) {
        if i == n {
            res = append(res, string(path))
            return
        }

        chars := mp[digits[i]]
        for _, c := range chars {
            path[i] = c
            dfs(i + 1)
        }
    }

    dfs(0)
    return res
}
