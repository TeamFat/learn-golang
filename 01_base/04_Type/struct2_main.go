/**
 *  datetime: 2017-05-25
 *  Go语言-结构体2
 */
package main

import "fmt"

type human struct {
	name  string
	age   int
	phone string
}

type student struct {
	human
	phone string //human中也有phone字段
}

func main() {
	// 重载字段，就近原则
	jack := student{human{name: "jack", age: 20, phone: "110"}, "119"}
	fmt.Println("jack phone uumber is ", jack.phone) //119
}
