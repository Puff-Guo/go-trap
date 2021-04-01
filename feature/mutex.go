package feature

import (
	"fmt"
	"sync"
)

type MyMutex struct {
	count int
	sync.Mutex
}

//锁赋值:会把锁状态一起拷贝
func MutexAssign() {
	var mu MyMutex
	mu.Lock()
	var mu1 = mu
	mu.count++
	mu.Unlock()
	mu1.Lock()
	mu1.count++
	mu1.Unlock()
	fmt.Println(mu.count, mu1.count)
}
