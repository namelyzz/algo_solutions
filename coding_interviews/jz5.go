package coding_interviews

import "strings"

/*
JZ5 替换空格

请实现一个函数，将一个字符串s中的每个空格替换成“%20“。
例如，当字符串为“We Are Happy“.则经过替换之后的字符串为“We%20Are%20Happy“。

注：我怎么记得以前是一道数组题目而不是字符串的...
*/
func replaceSpace(s string) string {
	return strings.Join(strings.Split(s, " "), "%20")
}
