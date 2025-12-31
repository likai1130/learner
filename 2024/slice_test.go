package _024

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestArray(t *testing.T) {
	var arr1 [5]int
	/*var arr2 [6]int
	var arr3 [5]string*/

	foo(arr1) // ok
	//foo(arr2) // 错误：[6]int与函数foo参数的类型[5]int不是同一数组类型
	//foo(arr3) // 错误：[5]string与函数foo参数的类型[5]int不是同一数组类型
}

func foo(arr [5]int) {}

func TestArrayMem(t *testing.T) {
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("数组长度：", len(arr))           // 6
	fmt.Println("数组大小：", unsafe.Sizeof(arr)) // 48
}
func TestSliceGrow(t *testing.T) {
}
