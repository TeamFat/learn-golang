/**
 *	datetime: 2017-05-27
 *  指针
 *  支持指针类型 *T， 指针的指针 **T， 以及包含包名前缀的*<package>.T
 *  & 取址符， 取变量内存地址
 *  * 取值符， 取变量内存地址对应的目标对象
 *  指针默认值nil, 没有NULL常量
 *  不支持指针运算，不支持"->"运算符，直接用"." 访问目标成员
 *
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

	type data struct{ user string }

	d1 := data{user: "jsser"}
	var p4 *data
	p4 = &d1
	fmt.Printf("p4=%p, p4.user=%v\n", p4, p4.user)

}

//$ go run main.go

// a 的变量地址：c42000e238
// p1变量内存地址c42000c028
// p1变量的值101
// 指向int的空指针值为0
// 空指针
// p4=0xc42000e280, p4.user=jsser
