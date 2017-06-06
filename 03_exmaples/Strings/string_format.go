/**
 * datetime: 2017-06-06
 * string format   字符串格式化
 * fmt.Printf 方法打印
 * 参考网址1：http://www.cnblogs.com/golove/p/3284304.html
 * 参考网址2：https://gobyexample.com/string-formatting
 * 格式字符串由普通字符和占位符组成，例如：abc%+ #8.3[3]vdef"
 * 其中 abc 和 def 是普通字符，其它部分是占位符，占位符以 % 开头（注：%% 将被转义为一个普通的 % 符号，这个不算开头），以动词结尾，格式如下：
 * %[旗标][宽度][.精度][arg索引]动词，方括号中的内容可以省略。
 */
package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main() {

	/**
	 * v：默认格式，不同类型的默认格式如下：
	   布尔型：t
	　　整　型：d
	　　浮点型：g
	　　复数型：g
	　　字符串：s
	　　通　道：p
	　　指　针：p
	 * #v：默认格式，以符合Go语法方式输出，特殊类型语法格式如下：无符号整型：x
	 *
	*/
	p := point{1, 2}
	//整数型 d
	fmt.Printf("%v\n", 1)
	//布尔型 t
	fmt.Printf("%v\n", true)
	//浮点数型 g
	fmt.Printf("%v\n", 3.13)
	//字符串 s
	fmt.Printf("%v\n", "jsser")
	//指针类型 p
	p1 := &p
	fmt.Printf("%v\n", &p1)
	//通道类型 p
	c1 := make(chan int, 1)
	fmt.Printf("%v\n", c1)
	//结构体
	fmt.Printf("%v\n", p)

	//使用无符号类型
	fmt.Printf("%#v\n", uint16(11)) //0xb

	/**
	 * T: 输出变量或值的类型
	 */
	fmt.Printf("%T\n", true) //bool
	fmt.Printf("%T\n", 3)    // int
	fmt.Printf("%T\n", p)    // main.point

	/**
	 *  整型
	 *
	 *  b/o/d：输出2/8/10 进制格式
	 *  x/X：输出16进制格式(大写/小写)
	 *  c：输出数值所表示的Unicode字符
	 *  q：输出数值所表示的 Unicode字符（带单引号）。对于无法显示的字符，将输出其转义字符
	 *  U：输出 Unicode 码点（例如 U+1234，等同于字符串 "U+%04X" 的显示结果）
	 */

	//转换2进制
	fmt.Printf("%b\n", 10)
	//转换8进制
	fmt.Printf("%o\n", 10)
	//转换16进制[小写]
	fmt.Printf("%x\n", 10)
	//转换16进制[大写]
	fmt.Printf("%X\n", 10)

}

//$ go run string_format.go
// 1
// true
// 3.13
// jsser
// 0xc42007e020
// 0xc420086000
// {1 2}
// 0xb
// bool
// int
// main.point
// 1010
// 12
// a
// A
