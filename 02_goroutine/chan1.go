/**
 * datetime: 2017-05-31
 * channel是有类型的管道, 可以使用channel操作符 <- 对其发送或者接收值
 * 1. ch <- v 将v发送到channel ch
 * 2. v := <- ch  从ch接收，并且赋值给v
 * 3. 使用make 创建channel
 */
package main

import (
	"fmt"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum //将sum值发送入到c中
}

func main() {

	a := []int{2, 3, 4, 5, 6, 1, 4, 5}
	//3. 使用make 创建channel
	c := make(chan int, 1)
	go sum(a, c)
	go sum(a[:2], c)
	x := <-c //从c接收，赋值给x
	y := <-c
	fmt.Println(x, y, x+y)
	// fmt.Println(x, y)
}

//$ go run
//5 30 35
