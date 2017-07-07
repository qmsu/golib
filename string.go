package golib

//字符串处理

// 字符串反转，支持中文
func Reverse(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

// 把字符串打散为数组
// separator	必需。规定在哪里分割字符串
// string	必需。要分割的字符串
// limit
// 可选。规定所返回的数组元素的数目
// 可能的值：
// 大于 0 - 返回包含最多 limit 个元素的数组
// 小于 0 - 返回包含除了最后的 -limit 个元素以外的所有元素的数组
// 0 - 返回包含一个元素的数组
// func Explode(s string, separator string, limit int) []interface{} {
// }

// 函数返回由数组元素组合成的字符串
// separator	可选。规定数组元素之间放置的内容。默认是 ""（空字符串）
// array	必需。要组合为字符串的数组
// func Implode(separator string, arr []interface{}) string {

// }
