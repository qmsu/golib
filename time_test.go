package golib

import (
	"testing"
	"fmt"
)

func TestInt64ToTime(t *testing.T) {
	i := int64(1484039584)
	st := Int64ToTime(i)
	fmt.Println("t is:",st)
}