/**
 *  datetime: 2017-05-25
 *  Go语言-结构体
 */
package main

import (
	"fmt"
)

/**
 *  声明一个person类型，包含两个字段，name 和 age
 */
type person struct {
	name string
	age  int
}

/**
 *	struct
 */
func older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

//struct 匿名字段

type human struct {
	name   string
	age    int
	weight int
}

//定义学生
//student 对 human 实现字段继承
type student struct {
	human      //匿名字段
	schoolName string
}

func main() {

	//使用person类
	var p person
	p.name = "jsser"
	p.age = 20
	fmt.Printf("The person's name is %s\n", p.name)
	//初始化1
	tom := person{age: 10, name: "tom"}
	//初始化2
	jack := person{"jack", 15}

	older, ageDiff := older(tom, jack)
	fmt.Printf("this old is %s,  diff %d\n", older.name, ageDiff)

	//初始化学生
	s1 := student{human{name: "jsser", age: 20, weight: 60}, "middle school"}
	fmt.Println("school Name is", s1.schoolName)
	fmt.Println("my name is ", s1.human.name)
	fmt.Println("my name is ", s1.name) //比较
	//修改年龄
	s1.age += 10
	fmt.Println("age is", s1.age)
}
