/**
 *  datetime: 2017-05-21
 *  Go语言-枚举类型
 *  枚举类型关键字：iota
 *  定义在常量组中，从0开始安行计算的自增枚举值。如果枚举值被非 _ 打算，那么需要手动显示恢复。
 */
package main

import "fmt"

// 星期
const (
	//星期天 初始化为0，后面为以后每一行为自增量
	Sunday = iota
	//星期一
	Monday
	//星期二
	Tuesday
	//星期三
	Wednesday
	//星期四
	Thurday
	//星期五
	Friday
	//星期六
	Saturday
)

const (
	c1, c2, c3 = iota, 3, iota // 0,3,0				   iota: 0
	c4         = 1             //iota被打断，必须显示恢复	iota: 1
	c5         = iota          //					   iota: 2
	c6                         //					   iota: 3
	c7         = 20            //					   iota: 4
	c8         = iota          // 显示恢复				iota: 5
	c9                         // 					   iota: 6
)

const (
	c01, c02 = iota, iota // 0, 0
	c03      = 10         //1
	c04      = iota       //2
)

// 使用 _ 符号  跳过 iota值
const (
	s1 = iota // iota: 0
	s2        // iota: 1
	s3        // iota: 2
	s4        // iota: 3
	_         // iota: 4
	_         // iota: 5
	s5        // iota: 6   不需要iota显示恢复
	s6        // iota: 7
)

// Effective go example

//ByteSize 字节换算
type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println("星期：", Sunday, Monday, Tuesday, Wednesday, Thurday, Friday, Saturday)
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9)
	fmt.Println(c01, c02, c03, c04)
	fmt.Println(s1, s2, s3, s4, s5, s6)
	fmt.Println(KB, MB, TB, PB, EB, ZB, YB)

}

// go run iota_main.go

/**
星期： 0 1 2 3 4 5 6
0 3 0 1 2 3 20 5 6
0 0 10 2
0 1 2 3 6 7
1024 1.048576e+06 1.099511627776e+12 1.125899906842624e+15 1.152921504606847e+18 1.1805916207174113e+21 1.2089258196146292e+24
*/
