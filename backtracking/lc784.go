package backtracking

import "unicode"

func letterCasePermutation(s string) []string {
    n := len(s)
    res := make([]string, 0, 1<<n)
    path := make([]rune, n)

    var dfs func(int)
    dfs = func(i int) {
        if i == n {
            res = append(res, string(path))
            return
        }

        c := rune(s[i])

        // 【不选】，不选时，这个字符不做任何处理，填入 path
        path[i] = c
        dfs(i + 1)

        // 【选】 选择时：
        // 小写：要转为大写
        // 大写：要转为小写
        // 数字：不处理，相当于不选，在不选的路径上自动处理了
        if 'a' <= c && c <= 'z' {
            path[i] = unicode.ToUpper(c)
            dfs(i + 1)
        } else if 'A' <= c && c <= 'Z' {
            path[i] = unicode.ToLower(c)
            dfs(i + 1)
        }
    }
    
    dfs(0)
    return res
}
