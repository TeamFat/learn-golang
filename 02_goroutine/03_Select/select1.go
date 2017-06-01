/**
 * datetime: 2017-06-01
 * select 语句似的一个goroutine在多个通讯操作上等待
 * 1. select是阻塞的，直到某个条件可以执行，如果多个条件都满足，select会随机执行一个条件
 * 2. 为了非阻塞的发送或者接受，select可以使用default分支，当所有条件都不可执行时，default分支会被执行
 * 3. select代码形式和switch很相似，select的case里面只能是IO操作
 */
package main

import (
	"fmt"
)

func test() {

}

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 1
	ch2 <- 2
	select {
	case <-ch1:
		fmt.Println("pop ch1")
	case <-ch2:
		fmt.Println("pop ch2")
	default:
		fmt.Println("default")
	}
	//会随机执行
}

//go run select1.go
// pop cha1 or pop ch2
