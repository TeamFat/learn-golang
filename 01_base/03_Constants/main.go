/**
 *  datetime: 2017-05-21
 *  Go语言-常量
 *  Go语言的常量值只能包含：字符串值、布尔值、数值类型(int,float,complex等)
 *  常量的值还可以使用 len, cap, unsafe.Sizeof 等编译器内可确定结果的函数返回值
 *  常量命名规范：官方推荐使用驼峰方式命名(以大写开头的常量为导出常量，小写开头的为包常量)
 */
package main

import "fmt"
import "unsafe"

// 声明ip常量
const ip string = "192.168.1.0"

// 声明ip2常量，并且常量类型通过值推到.
const ip2 = "192.168.1.1"

const d = 1.3

//声明一组常量
const (
	c1, c2 = 10, 20
	c4     = false
)

//声明组常量
const (
	c5 = 100
	c6 //c6的值为100。(如果常量不提供类型和初始化值，那么该常量的值和类型视作和上一个常量相同)
	c7 = 102
)

// 常量的值还可以使用 len, cap, unsafe.Sizeof 等编译器内可确定结果的函数返回值
var arr1 = [3]int{1, 2, 3} //声明一个长度为3的整形数组。(固定长度)

const (
	c8  = "c8 constant"
	c9  = len(c8)
	c10 = unsafe.Sizeof(c8)
	c11 = cap(arr1)
)

func main() {
	//声明c5变量，和常量名称有冲突，编译器不会有报错，这时c5的常量值被c5变量值覆盖。
	c5 := 10

	//在函数内部声明常量, 其作用域也是在函数内部，不能再其他函数内使用
	const c13 = 10000

	fmt.Println(c5)
	test() //打印常量c5

	//len cap unsafe.Sizeof

	fmt.Println(c9, c10, c11)

	//常量名称区分大小写
	// fmt.Println(C1)

}

func test() {
	//使用常量c5, 但是不能调用到c13常量
	fmt.Println(c5)
}

//$ go run main.go
/*
	10
	100
	11 16 3
*/
