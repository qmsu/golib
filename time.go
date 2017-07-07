package golib

import (
	"time"
)

// 时间处理

// int64转时间格式
func Int64ToTime(i int64) time.Time {
	res := time.Unix(i, 0)
	return res
}

// 根据指定日期返回本周的第几天
func DaysOfWeek(t time.Time) (int, bool) {
	w := map[string]int{
		"Sunday":    0,
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
	}

	now_w := t.Weekday().String()
	if n, ok := w[now_w]; ok {
		return n, ok
	}
	return 0, false
}

// 返回指定日期前N周的时间段
// t 指定日期
// n 多少周
func LastNWeekTime(t time.Time, n int) (start time.Time, end time.Time, ok bool) {
	day, ok := DaysOfWeek(t)
	if !ok {
		return
	}
	num := day + 7*n
	start = t.AddDate(0, 0, -num)
	return start, t, true
}
