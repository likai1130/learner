package _024

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	s := make([]int, 1)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	//fmt.Println(test())
}

func test() int {
	i := 0
	defer func(i int) {
		i = i + 1
		fmt.Printf("defer %d\n", i)
	}(i)
	fmt.Printf("return %d\n", i)
	return i
}
