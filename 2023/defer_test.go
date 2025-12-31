package _023

import (
	"fmt"
	"testing"
)

// chapter4/sources/deferred_func_8.go

func foo1() {
	sl := []int{1, 2, 3}
	defer func(a []int) {
		fmt.Println(a)
	}(sl)

	sl = []int{3, 2, 1}
	_ = sl
}

func foo2() {
	sl := []int{1, 2, 3}
	defer func(p *[]int) {
		fmt.Println(*p)
	}(&sl)

	sl = []int{3, 2, 1}
	_ = sl
}

func TestDefer(t *testing.T) {
	foo1()
	foo2()
}
