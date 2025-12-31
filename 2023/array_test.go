package _023

import (
	"fmt"
	"testing"
)

/*
*
测试数组值拷贝
*/
func TestValueCopy(t *testing.T) {
	a := [3]int{1, 2, 3}
	fmt.Printf("a:%v, %p\n", a, &a)

	b := a
	fmt.Printf("b:%v, %p\n", b, &b)

	CopyArray(a)
}

func CopyArray(c [3]int) {
	fmt.Printf("c:%v, %p\n", c, &c)
}

/*
*
测试切片也是值拷贝
*/
func TestSliceCopy(t *testing.T) {
	a := []int{1, 2, 3}
	fmt.Printf("a: %p\n", a)

	b := a
	fmt.Printf("b: %p\n", b)
	CopySlice(a)
	fmt.Printf("a:%v %p\n", a, a)
	fmt.Printf("b: %p\n", b)
}

func CopySlice(c []int) {
	c = append(c, 333)
	fmt.Printf("c: %p\n", c)
}

func TestCopy(t *testing.T) {
	ints := make([]int, 0, 2)
	ints = append(ints, 10)
	ints = append(ints, 20)

	var a = []int{1, 2, 3}
	c := a
	fmt.Printf("ints: %v, %p\n", ints, ints)

	fmt.Printf("a: %v, %p\n", a, a)
	fmt.Printf("c: %v, %p\n", c, c)
	fmt.Println(copy(ints, a))

	fmt.Printf("copy ints: %v, %p\n", ints, ints)
	fmt.Printf("copy a: %v, %p\n", a, a)
	fmt.Printf("copy c: %v, %p\n", c, c)
}

func TestInitSlice(t *testing.T) {
	var a []int
	b := []int{}
	c := new([]int)
	d := make([]int, 0, 0)
	e := make([]int, 0)
	fmt.Printf("init a : %v, %p\n", a, &a)
	fmt.Printf("init b : %v, %p\n", b, &b)
	fmt.Printf("init c : %v, %p\n", c, &c)
	fmt.Printf("init d : %v, %p\n", d, &d)
	fmt.Printf("init e : %v, %p\n", e, &e)

	if a == nil {
		fmt.Printf("init new a : %v, %p\n", a, a)
		//fmt.Println("a is nil")
	}
}

func TestMap(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"d": 1,
		"c": 1,
		"b": 1,
	}

	fmt.Println(len(m))
}

func TestNewSlice(t *testing.T) {
	// 零切片
	slice2 := make([]int, 5)  // 0 0 0 0 0
	slice2[0] = 1             //[1 0 0 0 0]
	slice3 := make([]*int, 5) // nil nil nil nil nil
	a := 1
	slice3[0] = &a //[0xc00009c1f8 <nil> <nil> <nil> <nil>]

	//nil切片
	var slice4 []int //[]
	//slice4[0] = 1            //panic: runtime error: index out of range [0] with length 0
	var slice5 = *new([]int) //[]
	//	slice5[0] = 1            //panic: runtime error: index out of range [0] with length 0

	//空切片
	var slice6 = []int{}
	//	slice6[0] = 1 //panic: runtime error: index out of range [0] with length 0
	var slice7 = make([]int, 0)
	//	slice7[0] = 1 //panic: runtime error: index out of range [0] with length 0

	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)
	fmt.Println(slice7)
}

type user struct {
	name string
	age  uint64
}

func TestRangeSlice(t *testing.T) {
	u := []user{
		{"asong", 23},
		{"song", 19},
		{"asong2020", 18},
	}
	for i, _ := range u {
		if u[i].age != 18 {
			u[i].age = 20
		}
	}
	fmt.Println(u)
}
