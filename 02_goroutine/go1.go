/**
 * datetime: 2017-05-30
 * Go提供原生的并发支持，类似协程，被称作goroutine机制
 * 在函数调用语句之前使用go关键字，就可以创建并发执行单元。
 * go并发执行，执行顺序是无序的
 * 当mian函数执行结束时，goroutine还未被运行
 *
 */
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go f1() //无需等待
	for i := 0; i < 5; i++ {
		//注意：将变量i 传入go函数，而不是在函数中直接使用i
		go func(i int) {
			fmt.Printf("i = %d\n", i)
			//使用如下两种方案，不能保证go函数优先main函数执行, 后面可以结合channel通信机制 来正确获取go函数的执行结果
			runtime.Gosched()
			time.Sleep(time.Microsecond)
		}(i)
	}
	fmt.Println("main end~")
}

func f1() {
	fmt.Println("this is f1 func")
}

//$ go run go1.go
// ... try to runing
