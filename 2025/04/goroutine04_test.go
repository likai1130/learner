package _04

import (
	"fmt"
	"testing"
	"time"
)

// 两个协程交替打印奇数偶数
func TestGoroutine04(t *testing.T) {
	ch := make(chan struct{})
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- struct{}{}
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			<-ch
			if i%2 != 0 {
				fmt.Println(i)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
