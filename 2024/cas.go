package _024

import (
	"sync/atomic"
	"testing"
)

func TestCAS(t *testing.T) {
	// 开启循环
	// 每次循环判断CAS有没有成功
	//设置兜底措施， 避免无限循环。1. 固定循环时间 2. 固定循环次数
	var newVal int64 = 1
	var oldVal int64 = 0
	for !atomic.CompareAndSwapInt64(oldVal, newVal) {

	}
}
