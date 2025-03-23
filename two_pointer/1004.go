package two_pointer

/*
题目可等价为：寻找数组中最长的子数组，其包含的 0 的个数不超过 k
思路：利用滑动窗口维护一个 “有效区间”，区间内的 0 个数不超过 k
窗口调整规则：
- 当右指针 i 指向 1 时：窗口可直接扩张，无需消耗翻转名额，更新最大窗口长度；
- 当右指针 i 指向 0 且 k > 0 时：消耗一个翻转名额（k--），窗口继续扩张，更新最大窗口长度；
- 当右指针 i 指向 0 且 k = 0 时：窗口内 0 个数已达上限，
  - 需移动左指针 start 收缩窗口，直到释放出一个翻转名额（即左指针跳过一个 0）。
*/
func longestOnes(nums []int, k int) int {
    var start, res int
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            res = max(res, i-start+1)
            continue
        }

        if nums[i] == 0 && k > 0 {
            k--
            res = max(res, i-start+1)
            continue
        }

        for k == 0 {
            if nums[start] == 0 {
                start++
                break
            }
            start++
        }

    }
    return res
}
