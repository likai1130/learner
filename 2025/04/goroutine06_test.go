package _04

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// sync.Cond实现多生产者多消费者
func TestGoroutine06(t *testing.T) {
	// sync.Cond 是一个条件变量，用于协调多个 goroutine 之间的同步
	// 它通常用于以下场景：
	// 1. 多生产者-多消费者模式：当缓冲区满时，生产者等待；当缓冲区空时，消费者等待
	// 2. 事件通知：当某个条件满足时，通知等待的 goroutine
	// 3. 资源池管理：当资源可用时通知等待的使用者
	//
	// sync.Cond 需要配合一个锁（Mutex 或 RWMutex）使用
	// 主要方法：
	// - Wait(): 释放锁并等待信号
	// - Signal(): 唤醒一个等待的 goroutine
	// - Broadcast(): 唤醒所有等待的 goroutine
	/*cond := sync.Cond{L: &sync.Mutex{}}
	cond.Broadcast()
	cond.L.Lock()
	cond.Wait()
	cond.L.Unlock()*/
	ctx := context.Background()
	timeoutCtx, cancelFunc := context.WithTimeout(ctx, 3*time.Second)
	defer cancelFunc()
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(timeoutCtx, &wg, ch, i)
	}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go consumer(timeoutCtx, &wg, ch, i)
	}
	wg.Wait()
	close(ch)
	fmt.Println("done")
}

func producer(ctx context.Context, wg *sync.WaitGroup, out chan<- int, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Producer %d finished\n", id)
			return
		default:
			num := rand.Intn(500)
			// 如果 channel 满了，这里会自动阻塞，无需手动 wait
			select {
			case out <- num:
				fmt.Printf("Producer %d sent: %d\n", id, num)
			case <-ctx.Done():
				fmt.Printf("Producer %d interrupted\n", id)
				return
			}
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100))) // 模拟生产耗时
		}
	}
}

func consumer(ctx context.Context, wg *sync.WaitGroup, in <-chan int, id int) {
	defer wg.Done()
	for {
		select {
		case num, ok := <-in:
			if !ok {
				// channel 已关闭，退出
				fmt.Printf("Consumer %d: channel closed\n", id)
				return
			}
			fmt.Printf("Consumer %d received: %d\n", id, num)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(150))) // 模拟消费耗时
		case <-ctx.Done():
			// 超时，但可能 channel 还有数据？我们选择优雅退出
			// 注意：此时 channel 未关闭，可能还有数据未消费
			// 如果要确保消费完，需额外协调（见下方说明）
			fmt.Printf("Consumer %d timed out\n", id)
			return
		}
	}
}
