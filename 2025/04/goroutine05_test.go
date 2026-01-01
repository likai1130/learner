package _04

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 单个channel实现01交替打印
func TestGoroutine05(t *testing.T) {
	msg := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			<-msg
			fmt.Println("a:0")
			msg <- struct{}{}
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-msg
			fmt.Println("b:1")
			msg <- struct{}{}
		}
	}()
	msg <- struct{}{}
	time.Sleep(1 * time.Second)
}

func TestGoroutine05X(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan struct{})
	//	ch <- struct{}{}
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("a:", 0)
			ch <- struct{}{}
			time.Sleep(1 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch
			fmt.Println("b:", 1)
			time.Sleep(1 * time.Millisecond)
		}
	}()
	wg.Wait()
}
