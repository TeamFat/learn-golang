/**
 * datetime: 2017-06-06
 * strings标准库 - 提供了许多有用的字符串相关函数
 *
 *
 */
package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	//查找字符串是否包含子字符串
	p("Contains: ", s.Contains("jsser", "sser"))
	//查找字符串包含子字符的个数
	p("Count: ", s.Count("jsser", "sser"))
	//判断字符串开头字符
	p("HasPrefix: ", s.HasPrefix("jsser", "js"))
	//判断字符串结尾字符
	p("HasSuffix: ", s.HasSuffix("jsser", "er"))
	//查找子字符串在字符串中的索引位置
	p("Index: ", s.Index("jsser", "e"))
	//将切片join成字符串
	p("Join: ", s.Join([]string{"j", "s", "s", "e", "r"}, "-"))
	//填充字符串 - str_pad
	p("Repeat: ", s.Repeat("j", 5))
	//字符串替换
	p("Replace: ", s.Replace("jsser", "s", "S", 1))
	p("Replace: ", s.Replace("jsser", "s", "S", -1))
	//split将字符串切割成数组
	p("Split: ", s.Split("jsser", ""))
	//将字符串转换成小写
	p("ToLower: ", s.ToLower("JSSer"))
	//将字符串转换成大写
	p("ToUpper", s.ToUpper("JSSER"))
	//获取字符串字符长度
	p("LEN: ", len("JSSER"))
	//以字符数组的形式访问，得到的结果为char类型  ASCII值
	p("Char: ", "jsser"[1])
}

//$ go run strings.go

// Contains:  true
// Count:  1
// HasPrefix:  true
// HasSuffix:  true
// Index:  3
// Join:  j-s-s-e-r
// Repeat:  jjjjj
// Replace:  jSser
// Replace:  jSSer
// Split:  [j s s e r]
// ToLower:  jsser
// ToUpper JSSER
// LEN:  5
// Char:  115
