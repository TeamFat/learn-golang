/**
 *  datetime: 2017-05-26
 *  Go语言 - for
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// 最常见的for循环结构  - 切记这里的i的作用域只是在for循环内部
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}
	// fmt.Printf("i = %d\n", i) - undefined: i

	//for 模拟while （golang中无while循环结构）
	s1 := 0
	for s1 < 10 {
		s1++
	}
	fmt.Println("s1 = ", s1)

	//死循环
	for { //for true
		i := 0
		i++
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}

}

//$ go run for.go
