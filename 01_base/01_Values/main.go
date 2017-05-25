/**
 *  datetime: 2017-05-21
 *  Go语言字面量：Go语言中值的表示形式
 */
package main

import "fmt"

func main() {
	/**
	 *  1. 基础类型的字面量
	 */
	//字符串：支持打印所有UTF-8类型字符串
	fmt.Println("This is String")
	fmt.Println("中国")
	fmt.Println("中国" + "心") //两个字符串拼接使用 + 号
	//整数值
	fmt.Println(100)
	//浮点数值
	fmt.Println(10.323)
	fmt.Println(19E-3)
	//布尔值
	fmt.Println(true)
	fmt.Println(false)
	/**
	 *  2. 符合类型字面量
	 * 	符合类型字面量值后面学到在看：主要包括 struct、array、slice、map 类型的值
	 */

}

//$ go run main.go
