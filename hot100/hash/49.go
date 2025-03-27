package hash

import (
	"maps"
	"slices"
)

/*
https://leetcode.cn/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-100-liked

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

示例 1:
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

解释：
在 strs 中没有字符串可以通过重新排列来形成 "bat"。
字符串 "nat" 和 "tan" 是字母异位词，因为它们可以重新排列以形成彼此。
字符串 "ate" ，"eat" 和 "tea" 是字母异位词，因为它们可以重新排列以形成彼此。

示例 2:
输入: strs = [""]
输出: [[""]]

示例 3:
输入: strs = ["a"]
输出: [["a"]]

提示：
1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母
*/

func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	mp := make(map[string][]string, n)
	for _, str := range strs {
		tmp := []byte(str)
		slices.Sort(tmp)
		k := string(tmp)
		mp[k] = append(mp[k], str)
	}
	return slices.Collect(maps.Values(mp))
}

/*
解题思路：
利用排序后的字符串作为哈希表键值，将异位词（字符相同但顺序不同）分组。
对每个字符串排序后生成唯一键，相同键的字符串归入同一组，
实现O(n*k log k)时间复杂度的分组（k为字符串平均长度）。

算法步骤：
初始化哈希表，键为排序后的字符串，值为字符串切片。
遍历输入字符串数组：
a. 将当前字符串转为字节切片并排序。
b. 将排序后的字节切片转回字符串作为键。
c. 将原字符串添加到哈希表中该键对应的切片。
收集哈希表所有值（即分组结果），返回二维字符串切片。
*/
