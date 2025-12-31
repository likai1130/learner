package _023

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
	time.Sleep(3 * time.Second)
	a := <-c
	fmt.Println(a)
}
