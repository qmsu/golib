package golib

import "time"

// 时间处理

//int64转时间格式
func Int64ToTime(i int64) time.Time {
	res := time.Unix(i, 0)
	return res
}
