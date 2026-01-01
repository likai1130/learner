package _04

import (
	"sync"
	"testing"
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
	cond := sync.Cond{L: &sync.Mutex{}}
	cond.Broadcast()
	cond.L.Lock()
	cond.Wait()
	cond.L.Unlock()
}
