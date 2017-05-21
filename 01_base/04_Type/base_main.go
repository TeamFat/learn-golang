/**
 *  datetime: 2017-05-21
 *  Go语言-基本数据类型 共有18个基本数据类型
 *  bool,byte,rune,int/uint, int8/uint8,int16/uint16, int32/uint32, int64/uint64, float32, float64, complex64, complex128, string
 *  数值类型支持：8进制，16进制，以及科学计数法
 */

package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	comment := ""
	_ = comment
	fmt.Println(unsafe.Sizeof(comment))  // 占用字节：16
	fmt.Println(reflect.TypeOf(comment)) // 变量类型：string
	/*
	 *  布尔类型
	 *	宽度：1字节
	 *  默认值：false
	 */
	var isActive bool
	fmt.Println(isActive, reflect.TypeOf(isActive))

	/**
	 *  byte
	 *  宽度：1字节
	 *  默认值：0
	 *  字节类型，uint8别名，可以看作为由8位二进制标示的无符号整数类型
	 *  范围：0 ~ 255
	 */
	var firstByte byte
	firstByte = 10                                    // firstByte > 255 or firstByte < 0  is error
	fmt.Println(firstByte, reflect.TypeOf(firstByte)) //0, uint8

	/**
	 *  rune
	 *	宽度：4字节
	 *  默认值：0
	 *  int32别名, 专注于存储unicode编码的单个字符
	 *  范围：-21亿 ~ 21亿
	 */
	var firstRune rune
	firstRune = 'a'
	firstRune = 29999
	fmt.Println(firstRune, unsafe.Sizeof(firstRune)) //占用4个字节

	/**
	 *  int,uint
	 *	整型：根据当前计算机架构自动判断是32位还是64位
	 *  宽度：4 / 8个字节
	 *  默认值：0
	 */
	var i1 uint
	i1 = 10
	fmt.Println(unsafe.Sizeof(i1)) // 32位计算机：4字节，64位计算机：8字节

	/**
	 *  int8,uint8 / int16, uint16 / int32,uint32 / int64, uint64
	 *  整型：表示由8位,16位，32位，64位二进制数的 有符号 / 无符号的整型
	 *  默认值：0
	 */

	// 一般使用int足够了，编译器可以根据操作系统自动初始化32位或者64位的整数
	var i2 int
	i2 = 2000000000000000000
	fmt.Println(reflect.TypeOf(i2), unsafe.Sizeof(i2)) // int 8

	/**
	 *  float32, float64
	 *  浮点型
	 *  默认值： 0
	 *  字节：4字 / 8字节
	 */

	var f1 float32
	fmt.Println(f1)
	var f2 float64
	fmt.Println(f2)

	//可以通过标准库math，查看各个数字类型的取值范围

	// math.MaxFloat32
	// math.MaxFloat64
	// math.MaxInt16
	// math.MaxInt32
	// math.MaxInt8
	// math.MaxUint16
	// math.MaxUint32
	// math.MaxUint64
	// math.MaxUint8

	// 数值类型可用标识：8进制，16进制，科学计数法
	a, b, c, d := 071, 0x1F, 1e9, math.MaxUint16
	_, _, _, _ = a, b, c, d

	/**
	 *   complex64, complex 128
	 *   复数类型
	 *   由float32 / float64 类型的实部和虚部联合表示
	 *   默认值：0+0i
	 */

	var c1 complex64
	c1 = 1 + 1i
	fmt.Println(c1)

	var c2 complex128
	c2 = 1.0 + 3.8i
	fmt.Println(c2)

	/**
	 *	字符串
	 *  string
	 *  默认值：""
	 *  一个字符串类型，表示一个字符串值的集合。字符串实质是字节序列。字符串的值一旦被创建，其内容不可更改
	 *  标准库：strings 中有大量字符串操作方法
	 */
	var s = "hello china a daksjdklasjdklasjdlaskjdlasjldasljd"
	// 占用字节数
	fmt.Println(unsafe.Sizeof(s))
	//查询字符串长度
	fmt.Println(len(s))
	fmt.Println(strings.Contains(s, "hello"))

	//字符串原型格式
	var strFormat = `
		hello
			world
	`
	fmt.Println(strFormat)

}
