package golib

// 生成随机数,三种模式：
// 1.纯数字
// 2.数字加字母
// 3.纯字母

import (
	"math/rand"
	"time"
)

// 返回一个取值范围在[0,n)的伪随机int值，如果n<=0 return -1
// seed 给定的种子,default time.Now().UnixNano()
func RandNum(n int, seed int64) int {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	s := rand.NewSource(seed)
	if n <= 0 {
		return -1
	}
	r := rand.New(s)
	return r.Intn(n)
}

// 返回一个[min,max]之间的随机数
func RandNumMinMax(min int, max int) int {
	if min == max {
		return max
	}
	return rand.Intn(max-min) + min
}
