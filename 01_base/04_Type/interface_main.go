/**
 *  datetime: 2017-05-24
 *  Go语言-interface(接口类型)
 *  接口类型用于定义一组方法行为, 并且方法只有方法声明，没有方法体。
 *  一个接口类型声明，可以嵌套其他接口类型。
 */
package main

import "fmt"

// 声明Talk接口类型 - 关键字 type, interface
type Talk interface {
	Hello (UserName string) string
	Talk(heard string) (saying string, end bool , err error)
}


type Chatbot interface {
	Name() string
	Talk //Talk类型接口
	Start() bool
	Shutdown() bool
}

// 自定义mytalk类型
type mytalk string;

//*mytalk实现Talk接口类型
func (talk *mytalk) Hello (UserName string) string {
	return UserName
}

func (talk *mytalk) Talk(heard string) (saying string, end bool,err error) {
	saying = heard
	return 
}

func (talk *mytalk) Name () string {
	return "name"
}


func (talk *mytalk) Start() bool {
	isStart := true
	return isStart
} 

func (tallk *mytalk) Shutdown() bool {
	down := true
	return down
}


func main () {

	// *myTalk 实现 Talk类型
	var talk Talk = new(mytalk) //返回地址
	_, ok := talk.(*mytalk)
	if ok {
		fmt.Println("is ok")
	}

	//断言测试
	var bot1 Chatbot = new(mytalk)
	_, ok2 := bot1.(*mytalk)
	if ok2 {
		fmt.Println("bot1 is ok")
	}
}

//$ go run interface_main.go

//
//is ok
//bot1 is ok
//