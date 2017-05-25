// tests
// interface类型为通用类型
// 将一个变量定义为interface类型，那么编译器根据变量的值来推动变量类型。
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var it interface{}
	it = 10 //不会报错 reflect.TypeOf(it)  int
	it = 10
	var it2 int
	it2 = 10
	_ = it2
	fmt.Println("it typeof ", reflect.TypeOf(it))         //int
	fmt.Println("it == i2  true or false ?  ", it == it2) // true
	it = "a"                                              //不会报错 reflect.TypeOf(it)  string
	it3 := "a"
	_ = it3
	fmt.Println("it typeof ", reflect.TypeOf(it))         //string
	fmt.Println("it == i2  true or false ?  ", it == it3) // true
}

//$ go run interface2_main.go

//it typeof  int
//it == i2  true or false ?   true
//it typeof  string
//it == i2  true or false ?   true
