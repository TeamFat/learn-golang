/**
 *  datetime: 2017-05-23
 *  Go语言 - if else
 *  1. 可省略条件表达式括号
 *  2. 支持初始化语句，可自定义代码块局部变量
 *  3. 代码块左大括号必须在条件表达式尾部
 */
package main

import (
	"fmt"
)

func main() {

	v1 := 0
	if v1 > 1 {
		fmt.Println("v1 > 1")
	} else if v1 == 0 {
		fmt.Println("v1 == 1")
	} else {
		fmt.Println("v1 < 1")
	}

	//支持初始化语句

	if v2 := 0; v2 < 10 {
		fmt.Println("v2 < 10")
	}

	// go 不支持三元操作符

}

//$ go run if.go

//v1 == 1
//v2 < 10
