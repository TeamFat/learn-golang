/**
 *  datetime: 2017-05-23
 *  Go语言 - 方法
 *  方法是函数的一种，他是某个数据类型关联在一起的函数
 */
package main

import "fmt"

//User 定义User结构
type User struct {
	name string
	age  int
}

//方法声明
//方法只是在func和函数名称之间加了一个圆括号包裹的接受者声明。(User)， 表示getName方法是和User数据类型关联
func (user *User) getName() (name string) {
	name = user.name
	return
}

func (user *User) getAge() (age int) {
	age = user.age
	return
}

// myint类型为值类型
type myint int

//值方法
func (i myint) add(other int) myint {
	i = i + myint(other)
	return i
}

func main() {
	//声明一个User
	var u1 User
	u1.name = "jsser"
	u1.age = 19
	fmt.Println("u1.getName = ", u1.getName())
	fmt.Println("u1.getAge = ", u1.getAge())

	//myint

}

// $go run method_func.go

//
//u1.getName =  jsser
//u1.getAge =  19
//
