package _04

import (
	"fmt"
	"testing"
)

// 3个goroutine交替打印abc 10 次 （15min完成）
// 思路: go1，go2，go3 通过三个信号量控制 sg1，sg2,sg3
//	go1读取信号量看是否能读 <- sg1 ,打印，完成后 写入sg2 <- struct{}{}
//  go2读取信号量看是否能读 <- 写入sg2 ,打印，完成后 写入sg3 <- struct{}{}
//  go2读取信号量看是否能读 <- 写入sg3 ,打印，完成后 写入sg1 <- struct{}{}
// for 10 表示10次
//   开始执行：启动第一个goroutine ,sg1 <- struct{}{},需要一个全局的状态控制器done

func TestGoroutine02(t *testing.T) {
	// 创建三个信号量通道，用于控制执行顺序
	sg1 := make(chan struct{}, 1) // 初始化时允许第一个goroutine执行
	sg2 := make(chan struct{}, 1)
	sg3 := make(chan struct{}, 1)

	// 用于控制执行次数的通道
	done := make(chan bool)

	// 启动goroutine 1 (打印 a)
	go func() {
		for i := 0; i < 10; i++ {
			<-sg1 // 等待信号
			fmt.Println("sg1: abc")
			sg2 <- struct{}{} // 通知goroutine 2
		}
	}()

	// 启动goroutine 2 (打印 b)
	go func() {
		for i := 0; i < 10; i++ {
			<-sg2 // 等待信号
			fmt.Println("sg2: abc")
			sg3 <- struct{}{} // 通知goroutine 3
		}
	}()

	// 启动goroutine 3 (打印 c)
	go func() {
		for i := 0; i < 10; i++ {
			<-sg3 // 等待信号
			fmt.Println("sg3: abc")
			fmt.Println() // 换行
			if i == 9 {   // 最后一次执行后不再发送信号给sg1
				close(done)
			} else {
				sg1 <- struct{}{} // 通知goroutine 1
			}
		}
	}()

	// 开始执行：启动第一个goroutine
	sg1 <- struct{}{}

	// 等待所有goroutine完成
	<-done
}
