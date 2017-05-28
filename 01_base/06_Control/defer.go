/**
 *  datetime: 2017-05-27
 *  defer 延迟调用。通常用于资源释放和错误处理
 */
package main

import (
	"fmt"
)

func main() {

	fmt.Println("before")
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	fmt.Println("after")
	//以上执行结果 - 当有多个defer时， 按照FILO次序进行(First in Last out) 顺序栈方式
	// before
	// after
	// defer2
	// defer1
}
