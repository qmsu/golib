package golib

import(
	"testing"
	"fmt"
)

func TestReverse(t *testing.T) {
	var testStr = "abc"
	fmt.Println(Reverse(testStr))
}