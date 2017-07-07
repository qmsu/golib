package golib

import (
	"fmt"
	"testing"
	"time"
)

func TestInt64ToTime(t *testing.T) {
	i := int64(1484039584)
	st := Int64ToTime(i)
	fmt.Println("t is:", st)
}

func TestDaysOfWeek(t *testing.T) {
	tt := time.Date(2017, time.Month(7), 7, 0, 0, 0, 0, time.Now().Location())
	d, ok := DaysOfWeek(tt)
	if ok {
		fmt.Println(d)
	}
}

func TestLastNWeekTime(t *testing.T) {
	tt := time.Now()
	ttt := time.Date(2017, time.Month(1), tt.Day(), 0, 0, 0, 0, tt.Location())

	start, end, ok := LastNWeekTime(ttt, 1)
	if ok {
		fmt.Printf("start is:%v;\nend is:%v", start, end)
	}
}
