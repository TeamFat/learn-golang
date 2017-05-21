/**
 * author: jsser
 * datetime: 2017-05-20
 * 01-hello.go
 * Go语言的第一个程序：hello world
 */

// 在go语言中，所有源码文件都必须是包的一部分，而main包为默认包
package main

// 引入fmt标准库
// fmt标准库提供基本的标准格式化功能
import "fmt"

// main包中必须包含main函数，才能够被正确运行
func main() {
	fmt.Print("hello world\n")
}

//$ go run hello.go
