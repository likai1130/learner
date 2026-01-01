package _04

import (
	"fmt"
	"testing"
	"time"
)

// 1. 开启100个协程，顺序打印1-1000，保证协程号1的，打印尾数为1的数字,协程2打印尾数为2以此类推 （15min完成）
func TestGoroutine01(t *testing.T) {
	// 思路: 携程池，用完释放，
	sig := make(chan struct{})          // 信号量
	pool := make(map[int]chan int, 100) // 线程池

	// 初始化线程池
	for i := 1; i <= 100; i++ {
		pool[i] = make(chan int)
	}

	// 监听各自协程号的数据，接收者先准备
	for i := 1; i <= 100; i++ {
		go func(i int) {
			for {
				nums := <-pool[i]
				fmt.Printf("goroutine%d 打印: %d\n", i, nums)
				sig <- struct{}{} // 执行完阻塞
			}
		}(i)
	}

	// 发送者再开始
	for i := 1; i <= 1000; i++ {
		lastNums := i % 100
		if lastNums == 0 {
			lastNums = 100
		}
		pool[lastNums] <- i // 绑定协程
		<-sig
	}
	time.Sleep(10 * time.Second)
}

func TestGoroutine01_2(t *testing.T) {
	// 思路: 携程池，用完释放，
	sig := make(chan struct{})          // 信号量
	pool := make(map[int]chan int, 100) // 线程池

	// 初始化线程池
	for i := 1; i <= 100; i++ {
		pool[i] = make(chan int)
	}

	go func() {
		// 发送者再开始
		for i := 1; i <= 1000; i++ {
			lastNums := i % 100
			if lastNums == 0 {
				lastNums = 100
			}
			pool[lastNums] <- i // 绑定协程
			sig <- struct{}{}   // 执行完阻塞
		}
	}()

	// 监听各自协程号的数据，接收者先准备
	for i := 1; i <= 100; i++ {
		go func(i int) {
			for {
				nums := <-pool[i]
				fmt.Printf("goroutine%d 打印: %d\n", i, nums)
				<-sig
			}
		}(i)
	}

	time.Sleep(10 * time.Second)
}
