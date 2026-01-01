package _04

import (
	"fmt"
	"sync"
	"testing"
)

// 使用不超过10个goroutine不重复的打印slice中的100个元素
func TestGoroutine03(t *testing.T) {
	nums := make([]int, 0, 100)
	for i := 1; i <= 100; i++ {
		nums = append(nums, i)
	}
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 0; i < len(nums); i++ {
		ch <- nums[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := <-ch
			fmt.Println(result)
		}()
	}
	wg.Wait()
}
