package _04

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 限制10个goroutine执行，每执行完一个，就放另外一个进来
func TestGoroutine09(t *testing.T) {
	ch := make(chan struct{}, 10)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine: %d \n", i)
			time.Sleep(time.Second)
			<-ch
		}(i)
	}
	wg.Wait()
}
