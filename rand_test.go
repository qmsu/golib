package golib

import (
	"fmt"
	"testing"
)

func TestRandNum(t *testing.T) {
	n := 10000
	r := RandNum(n, 0)
	fmt.Println("r is:", r)
}

func TestRandNumMinMax(t *testing.T) {
	min := 1000
	max := 9999
	for i := 0; i < max; i++ {
		r := RandNumMinMax(min, max)
		fmt.Println("r is: ", r)
	}

}
