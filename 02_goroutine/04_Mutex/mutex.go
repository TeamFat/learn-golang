/**
 * datetime: 2017-06-03
 * Sync.Mutex
 * 互斥锁
 * 在并发程序中，通过lock 和 unlock方法，保证同一个时刻只有一个goroutine访问一段代码。
 *
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

//Inc 递增计数
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	//++
	c.v[key]++
	c.mux.Unlock()
}

//Value 返回当前递增后的值
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()

	// 使用defer语句保证互斥锁会被解锁
	defer c.mux.Unlock()

	return c.v[key]
}

func main() {

	//使用SafeCounter
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("key")
	}
	time.Sleep(time.Second)
	fmt.Printf("c.Value(\"%s\") = %d\n", "key", c.Value("key"))
}

//$ go run mutex.go
//c.Value("key") = 1000
