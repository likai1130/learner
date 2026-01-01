package _04

import (
	"fmt"
	"sync"
	"testing"
)

// 使用两个goroutine交替向控制台打印字母与数字，输出结果为：a1b2c3d4e5f6g7h8i9j10k11l12m13n14o15p16q17r18s19t20u21v22w23x24y25z26
func TestGoroutine08(t *testing.T) {
	str := "abcdefghijklmnopqrstuvwxyz"
	ch := make(chan struct{})
	ch2 := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			fmt.Print(i)
			ch2 <- struct{}{}
			<-ch
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < len(str); i++ {
			<-ch2
			fmt.Print(string(str[i]))
			ch <- struct{}{}
		}
	}()
	wg.Wait()
	fmt.Println()
}
