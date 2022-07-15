package main

import (
	"fmt"
	"sync"
)

// 缓冲信道第二种用法，替代锁的机制
//“不要通过共享内存来通信，而是通过通信来共享内存"
type counter struct {
	c chan int
	i int
}

func NewCounter() *counter {
	cter :=  &counter{
		c: make(chan int),
	}

	go func() {
		for  {
			cter.i ++
			cter.c <- cter.i
		}
	}()
	return cter
}

func (cter *counter) Increase() int {
	return <- cter.c
}

func main()  {
	cter := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i <10; i++ {
		wg.Add(1)
		go func(i int) {
			increase := cter.Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, increase)
			wg.Done()
		}(i)
	}
	wg.Wait()

}