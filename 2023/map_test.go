package _023

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestNewMap(t *testing.T) {
	/*var intMap map[string]int                        // 声明map,没有分配内存的空映射
	fmt.Printf("intMap: %p %#v \n", &intMap, intMap) //intMap: 0xc000012058 map[string]int(nil)
	intMap["test"] = 1
	fmt.Println(intMap) // 因为map里面元素对象是nil，并未初始化，所以无法赋值*/

	nMap := new(map[string]int)               // 声明map，为map分配内存！返回的是指向map的指针
	fmt.Printf("nMap: %p %#v \n", nMap, nMap) // nMap: 0xc0000a4048 &map[string]int(nil)
	(*nMap)["test"] = 1                       // (*nMap)解引用，为map添加元素，因为new函数创建map返回的是指针类型*hmap，需要通过解引用访问指针指向的值。
	fmt.Println(nMap)                         // 依然是空指针，因为map里面元素对象是nil，并未初始化，所以无法赋值

	//结论是使用new为引用类型数据分配内存，初始化为nil，nil不能直接赋值
}

func TestMakeMap(t *testing.T) {
	m := make(map[string]int)
	fmt.Printf("m: %p %#v \n", m, m) //m: 0xc0000a4048 map[string]int{}
	m["test"] = 1
	fmt.Printf("m: %p %#v \n", m, m) //m: 0xc000012058 map[string]int{"test":1}

	// 结论：make不仅可以开辟一个内存，并且能给这个内存类型初始化零值
}

func TestNewMakeMap(t *testing.T) {
	var mv *map[string]int
	fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc0000a4048 (*map[string]int)(nil)

	mv = new(map[string]int)
	fmt.Printf("mv: %p %#v \n", &mv, mv) // mv: 0xc0000a4048 &map[string]int(nil)

	*mv = make(map[string]int)
	(*mv)["test"] = 1
	fmt.Printf("mv: %p %#v \n", &mv, mv) // mv: 0xc0000a4048 &map[string]int{"test":1}
}

func TestTopHash(t *testing.T) {
	m := make(map[string]int)
	m["apple"] = 1
	m["banana"] = 2
	m["cherry"] = 3

	// 获取 map 中的 tophash 值
	tophashApple := getTopHash(m, "apple")
	tophashBanana := getTopHash(m, "banana")
	tophashCherry := getTopHash(m, "cherry")

	fmt.Println("TopHash for apple:", tophashApple)
	fmt.Println("TopHash for banana:", tophashBanana)
	fmt.Println("TopHash for cherry:", tophashCherry)
}
func getTopHash(m map[string]int, key string) uint8 {
	// 通过类型转换将 tophash 值提取出来
	data := *(*[2]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&m)) + uintptr(len(m)) + uintptr(len(key)) + 2*unsafe.Sizeof(0)))
	return data[0]
}
