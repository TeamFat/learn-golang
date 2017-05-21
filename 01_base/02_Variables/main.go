/**
 *  datetime: 2017-05-21
 *  Go语言-变量
 *  Go语言为静态语言，不能在运行时改变变量类型
 */
package main

import (
	"fmt"
)

/**
 *  1. 在函数外部声明的变量是全局变量
 *  2. 全局变量未被使用，编译器不会报错
 *
 */
var g1 = 10
var v3 = "this is v3"
var v4 int

func main() {

	//Go语言使用 var 关键字声明变量，并且自动初始化为零值。如果提供初始化值，可省略变量类型，编译器会自动推断变量类型。
	//如果声明局部变量未被使用，那么程序会编译不通过。可通过 _ 来回收未使用的局部变量。
	//变量名称区分大小写
	var v0 int
	_ = v0
	var v1 int = 100 //这样声明变量，编译器会给予一个这样提示：it will be inferred from the right-hand side。大概意思是变量可通过右侧值推断出变量类型。
	_ = v1

	//在函数内部可以使用:=来声明变量。
	v3 := "this is v3 plus"
	_ = v3

	fmt.Println(v3)

	// 局部变量优先级高于全局变量，也就是说 局部变量会修改全局变量的值
	g1 := 100
	fmt.Println(g1)

	//可一次定义多个变量
	var v4, v5, v6 = "golang", "hello", "world"
	var v7, v8 int

	var (
		v9, v10 int
	)
	var (
		v11 int
		v12 int
	)
	_, _, _, _, _, _ = v7, v8, v9, v10, v11, v12

	fmt.Println(v4 + v5 + v6)

	//变量交换
	v13, v14 := 10, 20
	fmt.Printf("v13 = %d, v14 = %d\n", v13, v14)
	v13, v14 = v14, v13
	fmt.Printf("v13 = %d, v14 = %d\n", v13, v14)

	// 在Go中，同一个作用域下，不允许定义两个同名变量
	var v15 = 10
	// var v15 = 100 //或者 v15 := 10000
	fmt.Println(v15)
}

//$ go run main.go
