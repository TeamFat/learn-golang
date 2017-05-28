/**
 * datetime: 2017-05-28
 * Go类型转换
 * 不支持隐式类型转换，即便是从窄向宽转换也不行
 */
package main

import (
	"fmt"
)

func main() {

	var b byte = 100
	n := int(b) //显示类型转换
	fmt.Printf("n = %d\n", n)

	//不能将其他类型的值当做布尔值来判断

	a := 10
	if a { //此处会报错
		fmt.Println(10)
	}
}

//$ go run type_convert.go

//./type_convert.go:21: non-bool a (type int) used as if condition
