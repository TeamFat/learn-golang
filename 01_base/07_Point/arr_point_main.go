/**
 *	datetime: 2017-05-27
 *  golang 数组指针
 *
 */
package main

import (
	"fmt"
	"reflect"
)

//Max Constant Value
const Max = 2

func main() {

	slice1 := []int{10, 20, 30, 40}
	fmt.Println("slice1", slice1)
	//定义一个指针数组
	var ptr [Max]*int
	fmt.Println(reflect.TypeOf(ptr))
	for i := 0; i < Max; i++ {
		ptr[i] = &slice1[i] //将元素值的内存地址赋值给指针数组
	}
	fmt.Println(ptr)

	//循环指针数组
	for i := 0; i < Max; i++ {
		fmt.Printf("ptr[%d] = %d\n", i, *ptr[i])
	}

}

//$ go run arr_point_main.go

// slice1 [10 20 30 40]
// [2]*int
// [0xc42006c0a0 0xc42006c0a8]
// ptr[0] = 10
// ptr[1] = 20
