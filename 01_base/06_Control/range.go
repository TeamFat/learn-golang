/**
 *  datetime: 2017-05-26
 *  Go语言 - range
 */

package main

import (
	"fmt"
)

func main() {

	//初始化数组
	intArr1 := [10]int{1, 2, 3, 4, 5, 6, 7}
	//for range - 迭代器，返回 数组 每行的索引和值
	for index, item := range intArr1 {
		fmt.Printf("index = %d, item = %d\n", index, item)
	}

	// 遍历字符数组
	alp := "abcdefg"
	for i := range alp {
		fmt.Println(i) // 0,1,2,3,4,5,6  忽略值
	}

	//忽略索引
	for _, v := range alp { //忽略索引则必须用 _
		fmt.Printf("%c\n", v) // a,b,c,d,e,f,g
	}

	//忽略值和索引

	for range alp {

	}

	//初始化map
	m1 := map[string]int{}
	m1["k1"] = 1
	m1["k2"] = 2
	m1["k3"] = 3

	for key, value := range m1 {
		fmt.Printf("key = %s, value = %d\n", key, value)
	}

	//初始化slice
	s1 := []int{10, 20, 30, 40, 50}

	for i, v := range s1 {
		if i == 0 {
			s1 = s1[:2] //这里对slice 修改不会影响到range迭代，也就是说 range引用类型，其底层数据不会被复制
			s1[1] = 100
		}
		fmt.Println(i, v)
	}
	// s1
	fmt.Println("s1 ", s1)         // [10 100]
	fmt.Println("s1 cap", cap(s1)) //5
	fmt.Println("s1 len", len(s1)) // 2
}

//$ go run range.go
