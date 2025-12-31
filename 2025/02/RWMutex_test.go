package _2

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRWMutex(t *testing.T) {
	go A()
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println(count)
}

var mu sync.RWMutex
var count int

func A() {
	mu.RLock()
	defer mu.RUnlock()
	B()
}
func B() {
	time.Sleep(5 * time.Second)
	C()
}
func C() {
	mu.RLock()
	defer mu.RUnlock()
}
