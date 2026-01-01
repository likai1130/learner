package _04

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

// 使用go并发1000个控制，并设置超时1s
func TestGoroutine07(t *testing.T) {
	wg := sync.WaitGroup{}
	ctx := context.Background()
	task := make(chan int, 1000)
	ctxTime, cancelFunc := context.WithTimeout(ctx, 1*time.Second)
	defer cancelFunc()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		task <- i
		go func(ctxTime context.Context) {
			defer wg.Done()
			select {
			case <-ctxTime.Done():
				return
			default:
				taskNums := <-task
				log.Printf("goroutine:%d", taskNums)
			}
		}(ctxTime)
	}
	wg.Wait()
	close(task)
	fmt.Println("exec done")
}
