/**
 * datetime: 2017-05-31
 * channel是有类型的管道, 可以使用channel操作符 <- 对其发送或者接收值
 * 1. ch <- v 将v发送到channel ch
 * 2. v := <- ch  从ch接收，并且赋值给v
 * 3. 使用make 创建channel
 * 4. channel是可以带缓冲的，也就是make提供的第二个参数作为一个缓冲长度来初始化一个缓冲channel。 如果缓存长度参数为0，则是没有缓冲长度，发送给它的值立刻要被取走
 * 5. channel是典型的FIFO队列模型
 */
package main

import (
	"fmt"
)

func main() {

	//make 提供第二个参数作为chan的缓冲长度, 如果缓存长度参数为0，则是没有缓冲长度。发送给它的值立刻要被取走
	ch := make(chan int, 2)
	//发送数据到ch
	ch <- 1
	ch <- 2
	//ch <- 3 // 发送第三个数据到ch, 导致缓冲溢出，程序报错 //fatal error: all goroutines are asleep - deadlock!
	//顺序赋值
	fmt.Println(<-ch) //1
	fmt.Println(<-ch) //2

	ch <- 3
	ch <- 4
	fmt.Println(<-ch) //3
	fmt.Println(<-ch) //4
}

//$ go run
//5 30 35
