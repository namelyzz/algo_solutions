package coding_interviews

/*
JZ3 数组中重复的数字

在一个长度为n的数组里的所有数字都在0到n-1的范围内。
数组中某些数字是重复的，但不知道有几个数字是重复的。也不知道每个数字重复几次。
请找出数组中任意一个重复的数字。
例如，如果输入长度为7的数组[2,3,1,0,2,5,3]，
那么对应的输出是2或者3。存在不合法的输入的话输出-1

输入：
[2,3,1,0,2,5,3]
返回值：
2
说明：
2或3都是对的
*/

/*
思路：
去重这种题目最先想到的就是用哈希表（或者 Set）
*/
func duplicate(numbers []int) int {
	n := len(numbers)
	mp := make(map[int]struct{}, n)
	for _, v := range numbers {
		if _, ok := mp[v]; ok {
			return v
		}
		mp[v] = struct{}{}
	}
	return -1
}
