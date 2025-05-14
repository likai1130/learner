package _025

import (
	"fmt"
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	var mu MyMutex
	mu.Lock()
	var mu2 = mu
	mu.count++
	mu.Unlock()
	mu2.Lock()
	mu2.count++
	mu2.Unlock()
	fmt.Println(mu.count, mu2.count)
}

type MyMutex struct {
	count int
	sync.Mutex
}
