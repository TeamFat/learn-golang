/**
 *	datetime: 2017-05-27
 *  指针
 *  & 取址符， 取变量内存地址
 *  * 取值符， 取变量内存地址对应的值
 *
 */
package main

import (
	"fmt"
)

func main() {

	//指针使用流程
	// 1. 定义指针变量。2. 为指针变量赋值。3. 访问指针变量中指向地址的值

	var i int
	// 通过&符号，取得i变量的内存地址
	fmt.Printf("a 的变量地址：%x\n", &i)

	i = 100

	//声明指针 - 声明一个 指向int类型的指针
	var p1 *int

	p1 = &i //将i的地址赋予p1

	i = 101 //i的地址的值改变

	//为指针变量赋值
	fmt.Printf("p1变量内存地址%x\n", &p1)
	fmt.Printf("p1变量的值%d\n", *p1)

	//空指针 - nil
	var p3 *int
	fmt.Printf("指向int的空指针值为%x\n", p3)
	//空指针判断
	if p3 == nil {
		fmt.Println("空指针")
	} else {
		fmt.Println("非空指针")
	}

}

//$ go run main.go

// a 的变量地址：c42000e238
// p1变量内存地址c42000c028
// p1变量的值101
// 指向int的空指针值为0
// 空指针
