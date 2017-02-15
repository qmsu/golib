package golib

/**
 * 类型转换
 */

import "strconv"

func StringToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

func StringToInt(str string) int {
	i,_ := strconv.Atoi(str)
	return i
}

func IntToString(i int) string  {
	str := strconv.Itoa(i)
	return str
}

func Int64ToString(i int64) string {
	str := strconv.FormatInt(i,10)
	return str
}