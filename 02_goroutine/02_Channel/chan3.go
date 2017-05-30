/**
 * datetime: 2017-05-31
 * channel是有类型的管道, 可以使用channel操作符 <- 对其发送或者接收值
 * 1. ch <- v 将v发送到channel ch
 * 2. v := <- ch  从ch接收，并且赋值给v
 * 3. 使用make 创建channel
 * 4. channel是可以带缓冲的，也就是make提供的第二个参数作为一个缓冲长度来初始化一个缓冲channel。 如果缓存长度参数为0，则是没有缓冲长度，发送给它的值立刻要被取走
 * 5. channel是典型的FIFO队列模型
 * 6. channel一般无需关闭他们，只有在需要告诉接受者没有更多的数据时，才需要手动关闭channel, 例如中断range
 * 7. channel 可以通过len获取当前通道元素，也可以通过cap获取通道容量
 */
package main

import (
	"fmt"
)

// 斐波那契
// 0,1,1,2,3,5 ....

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c { // range 会不断的从c中接收值，直到c关闭
		fmt.Println(i)
	}
}
